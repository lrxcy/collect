# intro
> The dispatcher is an object (usually a mixin to other objects) that can register callback functions for particular events. 
Then when a dispatch method is called with an event, the dispatcher calls each callback function in order of their registration and passes them a copy of the event.

# quick start
1. go run main.go
2. curl -v -X POST "localhost:8000/work?delay=1s&name=foo"

# Feature Check
```
1.
When dispatching the event, a single error terminates all event handling.
It might be better to create a specific error type that terminates event handling (e.g. do not propagate)
and then collect all other errors into a slice and return them from the dispatcher.

2.
The event can technically be modified by callback functions since it’s a pointer.
It might be better to pass by value to guarantee that all callbacks see the original event.

3.
Callback handling is in order of registration, which gets to point number one about canceling event propagation.
An alternative is to do all the callbacks concurrently using Go routines; which is something I want to investigate further.
```

# refer
- https://bbengfort.github.io/snippets/2017/07/21/event-dispatcher.html
- https://github.com/mefellows/golang-worker-example