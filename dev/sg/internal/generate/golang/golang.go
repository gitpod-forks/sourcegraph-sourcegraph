package golang

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/sourcegraph/run"
	"golang.org/x/sync/semaphore"

	"github.com/sourcegraph/sourcegraph/dev/sg/internal/generate"
	"github.com/sourcegraph/sourcegraph/dev/sg/internal/std"
	"github.com/sourcegraph/sourcegraph/dev/sg/root"
	"github.com/sourcegraph/sourcegraph/lib/errors"
	"github.com/sourcegraph/sourcegraph/lib/output"
)

type OutputVerbosityType int

const (
	VerboseOutput OutputVerbosityType = iota
	NormalOutput
	QuietOutput
)

func Generate(ctx context.Context, args []string, progressBar bool, verbosity OutputVerbosityType) *generate.Report {
	// Save working directory
	cwd, err := os.Getwd()
	if err != nil {
		return &generate.Report{Err: err}
	}
	defer func() {
		os.Chdir(cwd)
	}()

	var (
		start     = time.Now()
		sb        strings.Builder
		reportOut = std.NewOutput(&sb, false)
	)

	// Run go generate [./...]
	if err := runGoGenerate(ctx, args, progressBar, verbosity, reportOut, &sb); err != nil {
		return &generate.Report{Output: sb.String(), Err: err}
	}

	// Run goimports -w
	if err := runGoImports(ctx, verbosity, reportOut, &sb); err != nil {
		return &generate.Report{Output: sb.String(), Err: err}
	}

	// Run go mod tidy
	if err := runGoModTidy(ctx, verbosity, reportOut, &sb); err != nil {
		return &generate.Report{Output: sb.String(), Err: err}
	}

	return &generate.Report{
		Output:   sb.String(),
		Duration: time.Since(start),
	}
}

var goGeneratePattern = regexp.MustCompile(`^//go:generate (.+)$`)

func findFilepathsWithGenerate(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	pathMap := map[string]struct{}{}
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			paths, err := findFilepathsWithGenerate(path)
			if err != nil {
				return nil, err
			}

			for _, path := range paths {
				pathMap[path] = struct{}{}
			}
		} else if filepath.Ext(entry.Name()) == ".go" {
			contents, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}

			for _, line := range bytes.Split(contents, []byte{'\n'}) {
				if goGeneratePattern.Match(line) {
					pathMap[path] = struct{}{}
					break
				}
			}
		}
	}

	paths := make([]string, 0, len(pathMap))
	for path := range pathMap {
		paths = append(paths, path)
	}

	return paths, nil
}

func runGoGenerate(ctx context.Context, args []string, progressBar bool, verbosity OutputVerbosityType, reportOut *std.Output, w io.Writer) (err error) {
	// Use the given packages.
	if len(args) != 0 {
		if verbosity != QuietOutput {
			reportOut.WriteLine(output.Linef(output.EmojiInfo, output.StyleBold, "go generate %s", strings.Join(args, " ")))
		}
		if err := runGoGenerateOnPaths(ctx, args, progressBar, verbosity, reportOut, w); err != nil {
			return errors.Wrap(err, "go generate")
		}

		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// If no packages are given, go for everything except doc/cli/references.
	// We cut down on the number of files we have to generate by looking for a
	// go:generate directive by hand first.
	paths, err := findFilepathsWithGenerate(wd)
	if err != nil {
		return err
	}
	filtered := make([]string, 0, len(paths))
	for _, pkgPath := range paths {
		pkgPath = pkgPath[len(wd)+1:]

		if !strings.HasPrefix(pkgPath, "doc/cli/references") {
			filtered = append(filtered, pkgPath)
		}
	}

	if verbosity != QuietOutput {
		reportOut.WriteLine(output.Linef(output.EmojiInfo, output.StyleBold, "go generate ./... (excluding doc/cli/references)"))
	}
	if err := runGoGenerateOnPaths(ctx, filtered, progressBar, verbosity, reportOut, w); err != nil {
		return errors.Wrap(err, "go generate")
	}

	return nil
}

// For debugging
const showTimings = false

func runGoGenerateOnPaths(ctx context.Context, pkgPaths []string, progressBar bool, verbosity OutputVerbosityType, reportOut *std.Output, w io.Writer) (err error) {
	var (
		done     = 0.0
		total    = float64(len(pkgPaths))
		progress output.Progress
		timings  = map[string]time.Duration{}
	)

	defer func() {
		if showTimings && verbosity == VerboseOutput {
			names := make([]string, 0, len(timings))
			for k := range timings {
				names = append(names, k)
			}

			sort.Slice(names, func(i, j int) bool {
				return timings[names[j]] < timings[names[i]]
			})

			progress.Write("\nDuration\tPackage")
			for _, name := range names {
				progress.Writef("%6dms\t%s", int(timings[name]/time.Millisecond), name)
			}
		}
	}()

	if progressBar {
		progress = std.Out.Progress([]output.ProgressBar{
			{Label: fmt.Sprintf("go generating %d packages", len(pkgPaths)), Max: total},
		}, nil)

		defer func() {
			if err != nil {
				progress.Destroy()
			} else {
				progress.Complete()
			}
		}()
	}

	var (
		m   sync.Mutex
		g   = errors.Group{}
		sem = semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
	)

	for _, pkgPath := range pkgPaths {
		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}

		// Do not capture loop variable in goroutine below
		pkgPath := pkgPath

		g.Go(func() error {
			defer sem.Release(1)

			if verbosity == VerboseOutput {
				progress.Writef("Generating %s...", pkgPath)
			}

			start := time.Now()
			if err := root.Run(run.Cmd(ctx, "go", "generate", pkgPath)).Wait(); err != nil {
				return err
			}
			duration := time.Since(start)

			m.Lock()
			defer m.Unlock()

			if progress != nil {
				done += 1.0
				progress.SetValue(0, done)
				progress.SetLabelAndRecalc(0, fmt.Sprintf("%d/%d packages generated", int(done), int(total)))
			}

			timings[pkgPath] = duration
			return nil
		})
	}

	return g.Wait()
}

func runGoImports(ctx context.Context, verbosity OutputVerbosityType, reportOut *std.Output, w io.Writer) error {
	if verbosity != QuietOutput {
		reportOut.WriteLine(output.Linef(output.EmojiInfo, output.StyleBold, "goimports -w"))
	}

	rootDir, err := root.RepositoryRoot()
	if err != nil {
		return err
	}

	// Determine which goimports we can use
	var goimportsBinary string
	if _, err := exec.LookPath("goimports"); err != nil {
		// Install goimports if not present
		err = run.Cmd(ctx, "go", "install", "golang.org/x/tools/cmd/goimports").
			Environ(os.Environ()).
			Env(map[string]string{
				// Install to local bin
				"GOBIN": filepath.Join(rootDir, ".bin"),
			}).
			Run().
			Stream(w)
		if err != nil {
			return errors.Wrap(err, "go install golang.org/x/tools/cmd/goimports returned an error")
		}

		goimportsBinary = "./.bin/goimports"
	} else {
		goimportsBinary = "goimports"
	}

	if err := root.Run(run.Cmd(ctx, goimportsBinary, "-w")).Stream(w); err != nil {
		return errors.Wrap(err, "goimports -w")
	}

	return nil
}

func runGoModTidy(ctx context.Context, verbosity OutputVerbosityType, reportOut *std.Output, w io.Writer) error {
	if verbosity != QuietOutput {
		reportOut.WriteLine(output.Linef(output.EmojiInfo, output.StyleBold, "go mod tidy"))
	}

	if err := root.Run(run.Cmd(ctx, "go", "mod", "tidy")).Stream(w); err != nil {
		return errors.Wrap(err, "go mod tidy")
	}

	return nil
}
