# math-grpc example
A microservices example using GoKit with gRPC

## Install protoc
1. Download protoc at: https://github.com/protocolbuffers/protobuf/releases/download/v21.1/protoc-21.1-win64.zip
2. Extract, then add extracted folder bin directory to path
3. Install add-on:
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Create pb folder and add math.proto file, paste bellow content:
```
syntax = "proto3";

option go_package = "grpc-math/pb";

service MathService {
  rpc Add(MathRequest) returns (MathResponse) {}
  rpc Subtract(MathRequest) returns (MathResponse) {}
  rpc Multiply(MathRequest) returns (MathResponse) {}
  rpc Divide(MathRequest) returns (MathResponse) {}
}

message MathRequest {
  float numA = 1;
  float numB = 2;
}

message MathResponse {
  float result = 1;
}
```

## Regenerate gRPC code
```
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\pb\math.proto
```

## Ref
```
https://github.com/junereycasuga/gokit-grpc-demo
```