## 生成go代码

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./pkg/grpc/ --go-grpc_opt==paths=source_relative pkg/grpc/service.proto
```