# mechanism
- main flow
```
main.go -> dispatcher.StartDispatcher(Active Workers Pool) -> go Serverrun (go signal)
```
- collecotr flow
```
client(Request) -> Collector -> send WorkRequest to chan with pre-define doFunc
```

# check cli
curl -v -X GET "localhost:8001/work?delay=1s&name=foo"

# expect output
```
Starting the dispatcher
Starting worker 1
Starting worker 2
Starting worker 3
Starting worker 4
Registering the collector
HTTP server listening on 127.0.0.1:8001
awaiting signal

2019/08/06 11:38:19 Work request queued
Received work requeust
Dispatching work request
2019/08/06 11:38:19 Doing shiit

2019/08/06 11:38:20 Work request queued
Received work requeust
Dispatching work request
2019/08/06 11:38:20 Doing shiit

2019/08/06 11:38:22 Work request queued
Received work requeust
Dispatching work request
2019/08/06 11:38:22 Doing shiit

```
