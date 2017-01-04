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

## Usage 

Edit proto buffer.  
```
$ vi ipersistence.proto
```

Generate codes.  
```
$ sh generate_code.sh
```

