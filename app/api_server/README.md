# APIサーバ

# 仕様・設計
- grpcを採用する。
- protoc / protoc-go-inject-tagを使用する。
```
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. <.proto>
protoc-go-inject-tag -input=<path>
```
- e-statのAPIを叩いてレスポンスを返す。
