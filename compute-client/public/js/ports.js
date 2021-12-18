function initElmPorts(app) {
  // Compute streaming
  var sources = {};

    function sendEventToElm(event) {
        // console.log(`Event: ${JSON.stringify(event.type)} : ${JSON.stringify(event.data)}`)
        app.ports.receiveEvent.send({
            data: event.data, // Can't be null according to spec
            eventType: event.type || null,
            id: event.id || null
        });
    }

    function newEventSource(address) {
        sources[address] = new EventSource(address);
        return sources[address];
    }

    function deleteAllEventSources() {
        for (const [key,] of Object.entries(sources)) {
            deleteEventSource(key);
        }
    }

    function deleteEventSource(address) {
        sources[address].close();
        delete sources[address];
    }

    app.ports.openStream.subscribe(function (args) {
        deleteAllEventSources(); // Pre-emptively close any open streams if we receive a request to open a new stream before seeing 'done'.
        console.log(`JS Port openStream. Args: ${JSON.stringify(args[0])}`)
        var address = args[0]; // We could listen on a specific event and get args[1] from the Elm app. No need for this right now.

        var eventSource = newEventSource(address);
        eventSource.onerror = function(err) {
            console.log(`EventSource failed: ${JSON.stringify(err)}`);
        }
        eventSource.addEventListener('results', sendEventToElm);
        eventSource.addEventListener('alert', sendEventToElm);
        eventSource.addEventListener('error', sendEventToElm);
        eventSource.addEventListener('done', function (event) {
            console.log("Done");
            deleteEventSource(address);
            // Note: 'done:true' is sent in progress too. But we want a 'done' for the entire stream in case we don't see it.
            sendEventToElm({ type: 'done', data: '' })
        });
    });
}