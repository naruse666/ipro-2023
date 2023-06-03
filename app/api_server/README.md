# APIサーバ

# 仕様・設計
- grpcを採用する。(protocを使用する。`protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. <.proto> `)
- e-statのAPIを叩いてレスポンスを返す。
