# これはなに？

# 開発環境構築

## Go

```bash
# goのインストール
brew install go

# バージョン確認
go version
```

## ghq

```bash
# インストール
brew install ghq

# 設定
git config --global ghq.root ~/workspace/src

# プロジェクト取得
ghq get git@github.com:aikizoku/rabbitgo.git
```

## yq

```bash
brew install yq
```

## Google Cloud SDK

```bash
# インストール
curl https://sdk.cloud.google.com | bash
exec -l $SHELL

# 初期化
gcloud init

# 新しいアカウントでログイン
gcloud auth login
```

## Rundoc

```bash
cd ~/
GO111MODULE=off go get github.com/aikizoku/rundoc
```

## Air

https://github.com/cosmtrek/air

## Terraform

```bash
brew install terraform
```

# 動かす

## 起動

```bash
cd appengine/default
air
```

## ローカルで確認

http://localhost:8080/ping

# デプロイ

## Cloud Build

develop branch に push で staging
master branch に push で production

## 手動

### AppEngine

```bash
cd appengine/default
gcloud app deploy app_staging.yaml --project xxxxxxx
gcloud app deploy app_production.yaml --project xxxxxxx
```

### Functions

```bash
cd functions/sample-handler
make deploy      # ステージング環境
make deploy-prod # 本番環境
```

# インフラ設定

## Terraform

```bash
cd terraform/staging
terraform plan
terraform apply

cd terraform/production
terraform plan
terraform apply
```

## Firestore の index をステージングから本番に同期

```bash
cd command/firestore_index
make get
make deploy-prod
```

# エラーに関するFAQ

## 起動時に `google: could not find default credentials.` が発生した

```bash
gcloud auth application-default login
```

## VScode で Functions 内のコードがエラー表示になる

https://qiita.com/chanhama/items/a21ca7d5cd43d6f3f90d

