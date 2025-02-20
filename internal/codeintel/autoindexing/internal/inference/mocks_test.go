// Code generated by go-mockgen 1.3.2; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package inference

import (
	"context"
	"io"
	"sync"

	regexp "github.com/grafana/regexp"
	api "github.com/sourcegraph/sourcegraph/internal/api"
	gitserver "github.com/sourcegraph/sourcegraph/internal/gitserver"
	luasandbox "github.com/sourcegraph/sourcegraph/internal/luasandbox"
)

// MockGitService is a mock implementation of the GitService interface (from
// the package
// github.com/sourcegraph/sourcegraph/internal/codeintel/autoindexing/internal/inference)
// used for unit testing.
type MockGitService struct {
	// ArchiveFunc is an instance of a mock function object controlling the
	// behavior of the method Archive.
	ArchiveFunc *GitServiceArchiveFunc
	// ListFilesFunc is an instance of a mock function object controlling
	// the behavior of the method ListFiles.
	ListFilesFunc *GitServiceListFilesFunc
}

// NewMockGitService creates a new mock of the GitService interface. All
// methods return zero values for all results, unless overwritten.
func NewMockGitService() *MockGitService {
	return &MockGitService{
		ArchiveFunc: &GitServiceArchiveFunc{
			defaultHook: func(context.Context, api.RepoName, gitserver.ArchiveOptions) (r0 io.ReadCloser, r1 error) {
				return
			},
		},
		ListFilesFunc: &GitServiceListFilesFunc{
			defaultHook: func(context.Context, api.RepoName, string, *regexp.Regexp) (r0 []string, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockGitService creates a new mock of the GitService interface.
// All methods panic on invocation, unless overwritten.
func NewStrictMockGitService() *MockGitService {
	return &MockGitService{
		ArchiveFunc: &GitServiceArchiveFunc{
			defaultHook: func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error) {
				panic("unexpected invocation of MockGitService.Archive")
			},
		},
		ListFilesFunc: &GitServiceListFilesFunc{
			defaultHook: func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error) {
				panic("unexpected invocation of MockGitService.ListFiles")
			},
		},
	}
}

// NewMockGitServiceFrom creates a new mock of the MockGitService interface.
// All methods delegate to the given implementation, unless overwritten.
func NewMockGitServiceFrom(i GitService) *MockGitService {
	return &MockGitService{
		ArchiveFunc: &GitServiceArchiveFunc{
			defaultHook: i.Archive,
		},
		ListFilesFunc: &GitServiceListFilesFunc{
			defaultHook: i.ListFiles,
		},
	}
}

// GitServiceArchiveFunc describes the behavior when the Archive method of
// the parent MockGitService instance is invoked.
type GitServiceArchiveFunc struct {
	defaultHook func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error)
	hooks       []func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error)
	history     []GitServiceArchiveFuncCall
	mutex       sync.Mutex
}

// Archive delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitService) Archive(v0 context.Context, v1 api.RepoName, v2 gitserver.ArchiveOptions) (io.ReadCloser, error) {
	r0, r1 := m.ArchiveFunc.nextHook()(v0, v1, v2)
	m.ArchiveFunc.appendCall(GitServiceArchiveFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Archive method of
// the parent MockGitService instance is invoked and the hook queue is
// empty.
func (f *GitServiceArchiveFunc) SetDefaultHook(hook func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Archive method of the parent MockGitService instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *GitServiceArchiveFunc) PushHook(hook func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitServiceArchiveFunc) SetDefaultReturn(r0 io.ReadCloser, r1 error) {
	f.SetDefaultHook(func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitServiceArchiveFunc) PushReturn(r0 io.ReadCloser, r1 error) {
	f.PushHook(func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error) {
		return r0, r1
	})
}

func (f *GitServiceArchiveFunc) nextHook() func(context.Context, api.RepoName, gitserver.ArchiveOptions) (io.ReadCloser, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitServiceArchiveFunc) appendCall(r0 GitServiceArchiveFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitServiceArchiveFuncCall objects
// describing the invocations of this function.
func (f *GitServiceArchiveFunc) History() []GitServiceArchiveFuncCall {
	f.mutex.Lock()
	history := make([]GitServiceArchiveFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitServiceArchiveFuncCall is an object that describes an invocation of
// method Archive on an instance of MockGitService.
type GitServiceArchiveFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 api.RepoName
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 gitserver.ArchiveOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 io.ReadCloser
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitServiceArchiveFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitServiceArchiveFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitServiceListFilesFunc describes the behavior when the ListFiles method
// of the parent MockGitService instance is invoked.
type GitServiceListFilesFunc struct {
	defaultHook func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error)
	hooks       []func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error)
	history     []GitServiceListFilesFuncCall
	mutex       sync.Mutex
}

// ListFiles delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitService) ListFiles(v0 context.Context, v1 api.RepoName, v2 string, v3 *regexp.Regexp) ([]string, error) {
	r0, r1 := m.ListFilesFunc.nextHook()(v0, v1, v2, v3)
	m.ListFilesFunc.appendCall(GitServiceListFilesFuncCall{v0, v1, v2, v3, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the ListFiles method of
// the parent MockGitService instance is invoked and the hook queue is
// empty.
func (f *GitServiceListFilesFunc) SetDefaultHook(hook func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ListFiles method of the parent MockGitService instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *GitServiceListFilesFunc) PushHook(hook func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitServiceListFilesFunc) SetDefaultReturn(r0 []string, r1 error) {
	f.SetDefaultHook(func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitServiceListFilesFunc) PushReturn(r0 []string, r1 error) {
	f.PushHook(func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error) {
		return r0, r1
	})
}

func (f *GitServiceListFilesFunc) nextHook() func(context.Context, api.RepoName, string, *regexp.Regexp) ([]string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitServiceListFilesFunc) appendCall(r0 GitServiceListFilesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitServiceListFilesFuncCall objects
// describing the invocations of this function.
func (f *GitServiceListFilesFunc) History() []GitServiceListFilesFuncCall {
	f.mutex.Lock()
	history := make([]GitServiceListFilesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitServiceListFilesFuncCall is an object that describes an invocation of
// method ListFiles on an instance of MockGitService.
type GitServiceListFilesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 api.RepoName
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 *regexp.Regexp
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitServiceListFilesFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitServiceListFilesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// MockSandboxService is a mock implementation of the SandboxService
// interface (from the package
// github.com/sourcegraph/sourcegraph/internal/codeintel/autoindexing/internal/inference)
// used for unit testing.
type MockSandboxService struct {
	// CreateSandboxFunc is an instance of a mock function object
	// controlling the behavior of the method CreateSandbox.
	CreateSandboxFunc *SandboxServiceCreateSandboxFunc
}

// NewMockSandboxService creates a new mock of the SandboxService interface.
// All methods return zero values for all results, unless overwritten.
func NewMockSandboxService() *MockSandboxService {
	return &MockSandboxService{
		CreateSandboxFunc: &SandboxServiceCreateSandboxFunc{
			defaultHook: func(context.Context, luasandbox.CreateOptions) (r0 *luasandbox.Sandbox, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockSandboxService creates a new mock of the SandboxService
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockSandboxService() *MockSandboxService {
	return &MockSandboxService{
		CreateSandboxFunc: &SandboxServiceCreateSandboxFunc{
			defaultHook: func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error) {
				panic("unexpected invocation of MockSandboxService.CreateSandbox")
			},
		},
	}
}

// NewMockSandboxServiceFrom creates a new mock of the MockSandboxService
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockSandboxServiceFrom(i SandboxService) *MockSandboxService {
	return &MockSandboxService{
		CreateSandboxFunc: &SandboxServiceCreateSandboxFunc{
			defaultHook: i.CreateSandbox,
		},
	}
}

// SandboxServiceCreateSandboxFunc describes the behavior when the
// CreateSandbox method of the parent MockSandboxService instance is
// invoked.
type SandboxServiceCreateSandboxFunc struct {
	defaultHook func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error)
	hooks       []func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error)
	history     []SandboxServiceCreateSandboxFuncCall
	mutex       sync.Mutex
}

// CreateSandbox delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockSandboxService) CreateSandbox(v0 context.Context, v1 luasandbox.CreateOptions) (*luasandbox.Sandbox, error) {
	r0, r1 := m.CreateSandboxFunc.nextHook()(v0, v1)
	m.CreateSandboxFunc.appendCall(SandboxServiceCreateSandboxFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the CreateSandbox method
// of the parent MockSandboxService instance is invoked and the hook queue
// is empty.
func (f *SandboxServiceCreateSandboxFunc) SetDefaultHook(hook func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// CreateSandbox method of the parent MockSandboxService instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *SandboxServiceCreateSandboxFunc) PushHook(hook func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *SandboxServiceCreateSandboxFunc) SetDefaultReturn(r0 *luasandbox.Sandbox, r1 error) {
	f.SetDefaultHook(func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *SandboxServiceCreateSandboxFunc) PushReturn(r0 *luasandbox.Sandbox, r1 error) {
	f.PushHook(func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error) {
		return r0, r1
	})
}

func (f *SandboxServiceCreateSandboxFunc) nextHook() func(context.Context, luasandbox.CreateOptions) (*luasandbox.Sandbox, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *SandboxServiceCreateSandboxFunc) appendCall(r0 SandboxServiceCreateSandboxFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of SandboxServiceCreateSandboxFuncCall objects
// describing the invocations of this function.
func (f *SandboxServiceCreateSandboxFunc) History() []SandboxServiceCreateSandboxFuncCall {
	f.mutex.Lock()
	history := make([]SandboxServiceCreateSandboxFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// SandboxServiceCreateSandboxFuncCall is an object that describes an
// invocation of method CreateSandbox on an instance of MockSandboxService.
type SandboxServiceCreateSandboxFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 luasandbox.CreateOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *luasandbox.Sandbox
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c SandboxServiceCreateSandboxFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c SandboxServiceCreateSandboxFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
