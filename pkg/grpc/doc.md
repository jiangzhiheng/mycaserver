## 生成go代码

```shell
cd pkg/grpc
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt==paths=source_relative service.proto
```