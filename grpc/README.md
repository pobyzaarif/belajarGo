protoc --proto_path=proto --go_out=. example.proto
protoc --proto_path=proto --go-grpc_out=. example.proto

python -m grpc_tools.protoc --proto_path=proto --python_out=./proto example.proto
python -m grpc_tools.protoc --proto_path=proto --grpc_python_out=./proto example.proto