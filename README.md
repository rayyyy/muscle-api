# muscle-api
筋トレメンターのマッチングアプリ

## 使い方
環境の起動用
```
docker-compose up -d
docker-compose exec app bash
```

サーバー起動
```
go run server.go
```

buildしてバイナリにしても起動できる
```
go build server.go
./server
```

## ディレクトリ説明
```
# サーバー起動用
server.go
# DBの定義や関連関数が入る予定
models
# コントローラー関数 それぞれのAPIの処理を入れる
controllers
```

## migrateについて

```
# マイグレーションファイル作成
goose -dir migration create create_user sql
# マイグレーション状況確認
goose mysql "root:pass@tcp(mysql:3306)/muscle_development?parseTime=true" status
# マイグレーション実行
goose -dir migration mysql "root:pass@tcp(mysql:3306)/muscle_development?parseTime=true" up
# マイグレーション ロールバック
goose -dir migration mysql "root:pass@tcp(mysql:3306)/muscle_development?parseTime=true" down
```

migrateしたあとは下のコマンドをつかってmodels配下にstructを作る
```
db2struct --host mysql -d muscle_development -t users --package models --struct userTable -p --user root --guregu --gorm
```