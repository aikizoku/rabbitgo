# これはなに？
GAE/goで色々やるためのテンプレート

APIとWebとTaskが作れるよ

# 開発環境構築
## Goのセットアップ
```bash
# goenv(Goのバージョン管理)のインストール
brew install goenv

# インストール可能なバージョンを確認
goenv install -l

# バージョンを指定してインストール
goenv install 1.8.7

# バージョン切り替え
goenv global 1.8.7

# バージョン確認
go version
```

## GOPATHを通す
```bash
vi .bash_profile

export GOPATH=$HOME/go

source .bash_profile
```

## ghq(リポジトリ管理)のセットアップ
```bash
# インストール
brew install ghq

# 設定
git config --global ghq.root $GOPATH/src

# Goプロジェクトを取得
ghq get xxxxxxxxx
```

## Google Cloud SDKのセットアップ
```bash
# 対話型パッケージ
curl https://sdk.cloud.google.com | bash

# シェルを再起動
exec -l $SHELL

# 初期化
gcloud init
```

## 依存パッケージのインストール
```bash
go get ./...
```

## 動かす
```bash
# API
make run svc=api

# Web
make run svc=web

# Task
make run svc=task
```

## デプロイ
```bash
# API
make deploy svc=api

# Web
make deploy svc=web

# Task
make deploy svc=task
```



