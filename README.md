# study-record-app

## 概要

自分自身の勉強時間を記録するアプリです。

## go-api

### ローカル環境

docker 起動

```shell
docker compose -f ./docker/docker-compose.yml up -d
```

※注意

テストデータを変更したい場合は以下実行後、再度 docker を起動

```shell
docker compose -f ./docker/docker-compose.yml down -v
```

API 起動

```shell
go run cmd/main.go
```
