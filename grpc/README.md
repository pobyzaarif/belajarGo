```
..../grpc $ protoc --proto_path=proto --go_out=. example.proto                                      // create proto/example.pb.go
..../grpc $ protoc --proto_path=proto --go-grpc_out=. example.proto                                 // create proto/example_grpc.pb.go

..../grpc $ python -m grpc_tools.protoc --proto_path=proto --python_out=./proto example.proto       // create proto/example_pb2.py
..../grpc $ python -m grpc_tools.protoc --proto_path=proto --grpc_python_out=./proto example.proto  // create proto/example_pb2_grpc.py
```