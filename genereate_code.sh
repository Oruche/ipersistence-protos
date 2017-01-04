#!/bin/sh

python -m grpc_tools.protoc ipersistence.proto --python_out=python --grpc_python_out=python --proto_path=./
protoc ipersistence.proto --go_out=plugins=grpc:go
