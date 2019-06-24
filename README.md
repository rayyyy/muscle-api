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