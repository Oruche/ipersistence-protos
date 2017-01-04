# ipersistence-protos
This repository manages ipersistence's protocol buffer and gRPC code.

## Set up for developments.

You need to download protocol buffer [here](https://github.com/google/protobuf).

Recommendation:
We recommend you to download zip(e.x protoc-3.1.0-osx-x86_64.zip) from [here](https://github.com/google/protobuf/releases)

```
export PATH="/path/to/protoc/bin:$PATH"
```

For golang,
```
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get -u google.golang.org/grpc
```

For python
```
$ pip install grpcio
$ pip install grpcio-tools
```

Example edit proto buffer
```
$ vi ipersistence.proto
```

Example: creating python code.
```
$ python -m grpc_tools.protoc -I./ --python_out=python --grpc_python_out=python ipersistence.proto
```

Example: creating go code.
```
$ protoc -I ./ ipersistence.proto --go_out=plugins=grpc:go
```
