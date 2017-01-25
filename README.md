# protoshell
protoshell is a tool for generating gRPC code from protocol buffer easily.

## Usage

### 1.Create and edit protocol buffer.  

```
$ cp ./protos/sample.proto ./protos/<YOUR_PROTO>.proto
$ vi ./protos/<YOUR_PROTO>.proto
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
$ sh proto.sh generate -p [PROTO_FILE] -l [LANG] -o [OUTPUT_DIR]

# example  
$ sh proto.sh generate -p sample.proto -l go -o ~/proto_dist/
```

help message.
```
$ sh proto.sh --help
```
