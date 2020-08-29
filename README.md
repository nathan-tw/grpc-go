
# hello-grpc-go

this is the first try for me to build a client-server connection via gRPC in golang


## Contrast

|  feature  |              gRPC              |    http APIs   |
|:---------:|:------------------------------:|:--------------:|
|  contract |            required            |    optional    |
|  protocol |             http/2             |      http      |
|  payload  |            protobuf            |  json(or more) |
| Streaming | client, server, bi-directional | client, server |

