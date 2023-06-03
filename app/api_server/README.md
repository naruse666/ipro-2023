# APIサーバ

# 仕様・設計
- grpcを採用する。(protoc-gen-go-grpcを使用する。`protoc <.proto> --go-grpc_out=<path>`)
- e-statのAPIを叩いてレスポンスを返す。
