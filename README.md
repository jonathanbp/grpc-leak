# Heavy memory usage demo for grpc-go

This repo demonstrates how a client with max-message set to a low limit will cause heavy memory usage (and possible leak) in a grpc server.

## Run the server

    > go run server/main.go

### Monitor mem usage

Requires [jplot](https://github.com/rs/jplot).

    > jplot --url http://localhost:8002/debug/vars memstats.HeapSys+memstats.HeapAlloc+memstats.HeapIdle+marker,counter:memstats.NumGC counter:memstats.TotalAlloc memstats.HeapObjects memstats.StackSys+memstats.StackInuse

## Run the client

Note that the client is rather aggressive and will use cause lots of CPU and RAM to be used.

    > go run client/main.go


### Run the client with a larger max memory size

Change ln 22 in `client/main.go` to a higher amount - e.g. `10*1024*1024` this will work much better.

