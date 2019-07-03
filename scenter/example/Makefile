
proto-go:
	protoc -I api/ \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:api \
    api/api.proto

proto-py:
	python3 -m grpc_tools.protoc -I api/ \
    --python_out=./api --grpc_python_out=./api \
    api/api.proto