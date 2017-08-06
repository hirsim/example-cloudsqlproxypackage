# Cloud SQL Proxy package サンプル
[Cloud SQL Proxy](https://github.com/GoogleCloudPlatform/cloudsql-proxy) を Golang のパッケージとして利用したサンプルです。

## 使用パッケージ

```
$ go get github.com/GoogleCloudPlatform/cloudsql-proxy
$ go get golang.org/x/oauth2

# MySQL
$ go get github.com/go-sql-driver/mysql

# PostgreSQL
$ go get github.com/lib/pq
```

## ビルド

```
# MySQL
$ cd mysql
$ go build -o example-cloudsqlproxypackage main.go

# PostgreSQL
$ cd postgresql
$ go build -o example-cloudsqlproxypackage main.go
```

## 実行

```
# GOOGLE_APPLICATION_CREDENTIALS 環境変数使用
$ GOOGLE_APPLICATION_CREDENTIALS=PATH_TO_CREDENTIAL_FILE ./example-cloudsqlproxypackage -db_addr 'インスタンス接続名' -db_name データベース名 -db_user ユーザ名 -db_pass ユーザパスワード（無しの場合は不要）

# 引数で秘密鍵へのパス指定
$ ./example-cloudsqlproxypackage -db_addr 'インスタンス接続名' -db_name データベース名 -db_user ユーザ名 -db_pass ユーザパスワード（無しの場合は不要） -credential_file 秘密鍵へのパス
```