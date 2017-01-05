# ipersistence-protos
This repository manages a ipersistence's protocol buffer and a script which generates gRPC codes.

## Usage

### 1.Edit proto buffer.  

```
$ vi ipersistence.proto
```

### 2. Install each language plugins.  

For go:
```
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get -u google.golang.org/grpc
```

For python:
```
$ pip install grpcio
$ pip install grpcio-tools
```

### 3. Generate codes.

```
$ sh proto.sh generate -l [LANG] -o [OUTPUT_DIR]
```

help message.
```
$ sh proto.sh --help
```
