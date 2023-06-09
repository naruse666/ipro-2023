# アプリケーション
- サービスごとにディレクトリを分ける
- [finch](https://github.com/runfinch/finch)を使って開発をする

# API関連
## 規約の重要なとこ
- アプリケーションIDを使用してAPIが叩けるようになる
- 5条3項: 本機能の負荷状況に応じてアクセス制限をかけることがあります
- 第8条(禁止事項): 短時間における大量のアクセスその他本機能の運用に支障を与える行為
- サービス公開のために,`「このサービスは、政府統計総合窓口(e-Stat)のAPI機能を使用していますが、サービスの内容は国によって保証されたものではありません。」`のクレジット表示が必要. [参考](https://developer.yahoo.co.jp/attribution/)

## 仕様
- 最新バージョンの`3.0`を利用する
- [API3.0の詳細仕様書](https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0)
- [提供データ大分類](https://www.e-stat.go.jp/api/api-info/statsfield)
- [提供データ一覧](https://www.e-stat.go.jp/stat-search/database?page=1)
- [マイページ](https://www.e-stat.go.jp/mypage/view/api)から`URL`を変更して利用するドメインを指定。開発用に`http://test.localhost/`が提供されている

## APIドキュメント
- [e-statの機能概要](https://www.e-stat.go.jp/api/api-info)
- [e-stat swagger](https://api.e-stat.go.jp/swagger-ui/e-statapi3.0.html#/)
- [e-stat mypage (appIdとかが置いてある)](https://www.e-stat.go.jp/mypage/view/api)
- [e-stat database 検索ページ](https://www.e-stat.go.jp/stat-search/database)

# 構成
<img width="852" alt="image" src="https://github.com/naruse666/ipro-2023/assets/62323683/754f59f3-2ae3-4bfb-81e1-3ae70e5f493b">


# ビルド
各フォルダでfinchを実行する。  
```
finch compose up -d
```
