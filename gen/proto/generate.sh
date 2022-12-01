#!/bin/bash

GO_OUT_RELATIVE_PATH="../../service-go/internal/api"

protoc --go_out=$GO_OUT_RELATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$GO_OUT_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/interview.proto