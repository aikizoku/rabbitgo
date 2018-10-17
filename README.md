# これはなに？
GAE/Go環境で動作するサーバー開発のテンプレート
とても早くて軽いAPI/Workerをワンソースで作る事が出来ます。
- インフラをあまり考えなくて良い
- 適当に作っていても循環参照が発生しない
- 緩い命名規則で縛っているので、柔軟かつ迷わない命名が可能
- 新しい機能を追加するときも本テンプレートの対応を待たないで素直に実装可能

# 開発環境構築
## Goのセットアップ
```bash
# goenv(Goのバージョン管理)のインストール
brew install goenv

# goenv(Goのバージョン管理)のアップデート
brew upgrade goenv

# インストール可能なバージョンを確認
goenv install -l

# バージョンを指定してインストール
goenv install 1.11.1

# バージョン切り替え
goenv global 1.11.1

# バージョン確認
go version
```

## GOPATHを通す
```bash
vi .bash_profile

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

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

# 新しいアカウントでログイン
gcloud auth login

# アカウントリスト
$ gcloud auth list

# アカウントの切り替え
$ gcloud config set account <your-account>

# 自分のプロジェクトリスト
gcloud projects list

# プロジェクトの切り替え
gcloud config set project <your-project-id>
```

## 依存パッケージのインストール
```bash
# インストール
brew install dep

# 依存パッケージのインストール
dep ensure

# 依存パッケージのアップデート
dep ensure update
```

# 動かす
## 起動
```bash
# API
make run-appengine-app app=api
make run-appengine-app-production app=api

# Worker
make run-appengine-app app=worker
make run-appengine-app-production app=worker
```

## 各種データを確認
```
http://localhost:8000/instances
```

## デプロイ
```bash
# API
make deploy-appengine-app app=api
make deploy-appengine-app-production app=api

# Worker
make deploy-appengine-app app=worker
make deploy-appengine-app-production app=worker

# Cron
make deploy-appengine-cron
make deploy-appengine-cron-production

# Dispatch
make deploy-appengine-dispatch
make deploy-appengine-dispatch-production

# Index
make deploy-appengine-index
make deploy-appengine-index-production

# Queue
make deploy-appengine-queue
make deploy-appengine-queue-production
```
