#!/bin/bash

SERVICE_RELATIVE_PATH="../../service-go/internal/api"
SERVICE_CLIENT_PATH="../../client-go/internal/api"

protoc --go_out=$SERVICE_RELATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$SERVICE_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/interview.proto

protoc --go_out=$SERVICE_CLIENT_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$SERVICE_CLIENT_PATH --go-grpc_opt=paths=source_relative \
    interview/interview.proto