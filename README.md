# これはなに？
GAE/Go環境での爆速で汎用性高い開発を目指したフレームワーク
とても早くて軽いAPIやWorkerをワンソースで作る事が出来る
- インフラをあまり考えなくて良い
- 適当に作っていても循環参照が発生しない
- 緩い命名規則で縛っているので、柔軟かつ迷わない命名が可能
- 新しい機能を追加するときもフレームワークの対応を待たないで素直に実装可能
- 実務で困らない範囲の役務分担と抽象化
- 難しく考えずにサクサク開発できる

# 開発環境構築
## Goのセットアップ
```bash
# goenv(Goのバージョン管理)のインストール
brew install goenv

# インストール可能なバージョンを確認
goenv install -l

# バージョンを指定してインストール(Go1.11.x系の最新選択)
goenv install 1.12.x

# バージョン切り替え
goenv global 1.12.x

# バージョン確認
go version
```

## ghq(リポジトリ管理)のセットアップ
```bash
# インストール
brew install ghq

# 設定
git config --global ghq.root $GOPATH/src

# Goプロジェクトを取得(例:rabbitgoの場合)
ghq get git@github.com:aikizoku/rabbitgo.git
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
```

## 依存パッケージのインストール
```bash
cd appengine/default
GO111MODULE=on go test
```

# 動かす
## 起動
```bash
# GoogleAppEngine
cd appengine
make run name=default
```

## ローカルで確認
```
# 動作確認
http://localhost:8080/ping

# 状況確認
http://localhost:5002/
```

## デプロイ
```bash
# Google App Engine
cd appengine
make deploy name=default            # ステージング環境
make deploy-production name=default # 本番環境

# Cloud Functions
cd functions
make deploy name=function-name            # ステージング環境
make deploy-production name=function-name # 本番環境

# Cloud Scheduler
cd scheduler
make deploy-http name=schedule-name                 # HTTPのURLを実行する ステージング環境
make deploy-http-production name=schedule-name      # HTTPのURLを実行する 本番環境
make deploy-appengine name=schedule-name            # GAEのAPIを実行する ステージング環境
make deploy-appengine-production name=schedule-name # GAEのAPIを実行する 本番環境
make deploy-pubsub name=schedule-name               # PubSubのトピックに送信する ステージング環境
make deploy-pubsub-production name=schedule-name    # PubSubのトピックに送信する 本番環境

# Cloud Tasks
cd tasks
make deploy            # ステージング環境
make deploy-production # 本番環境

# Cloud Pub/Sub
cd pubsub
make deploy            # ステージング環境
make deploy-production # 本番環境
```

# 開発で使う便利なコマンド集
```bash
###### Go ######

# goenv(Goのバージョン管理)のインストール
brew install goenv

# goenv(Goのバージョン管理)のアップデート
brew upgrade goenv

# インストール可能なバージョンを確認
goenv install -l

# バージョンを指定してインストール
goenv install 1.12.7

# バージョン切り替え
goenv global 1.12.7

# バージョン確認
go version

###### ghq ######

# Goプロジェクトを取得(例:rabbitgoの場合)
ghq get git@github.com:aikizoku/rabbitgo.git

###### Google Cloud SDK ######

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

### Go Modules ###
# 初期化
go mod init

# 依存パッケージのインストール
go get

# 依存パッケージの更新
go get -u

# 依存パッケージの追加
go get -u hogehoge

# 依存パッケージの整理
go mod tidy
```

# よく使うコード
```golang
/****** Logging ******/
// エラーログを出力したい時
// 出力されるログ: time [ERROR] foo/bar.go:21 hoge 123
log.Errorf(ctx, "hoge %d", 123)

// エラーログを出力したいが文言考えるの面倒なので定型文を使いたい時
// 出力されるログ: time [ERROR] foo/bar.go:21 h.svc.Sample error: invalid params
log.Errorm(ctx, "h.svc.Sample", err)

// エラーログを出力すると同時にエラーを作成したい
// 出力されるログ: time [ERROR] foo/bar.go:21 hoge 123
err := log.Errore(ctx, "hoge %d", 123)

// エラーに任意のエラーコードを埋め込む
err = errcode.Set(err, 404)

// エラーからエラーコードを取り出す
code, ok := errcode.Get(err)

// エラーログを出力すると同時にエラーコードを含むエラーを作成したい
// 出力されるログ: time [ERROR] foo/bar.go:21 hoge 123
err := log.Errorc(ctx, http.StatusNotFound, "hoge %d", 123)

/****** リクエストの値を取得 ******/

// HTTPHeaderの値を取得
headerParams := httpheader.GetParams(ctx)
log.Debugf(ctx, "HeaderParams: %v", headerParams)

// URLParamの値を取得
urlParam := handler.GetURLParam(r, "sample")
if urlParam == "" {
  h.handleError(ctx, w, http.StatusBadRequest, "invalid url param is empty")
  return
}
log.Debugf(ctx, "URLParam: %s", urlParam)

// フォームの値を取得
formParam := handler.GetFormValue(r, "sample")
if formParam == "" {
  h.handleError(ctx, w, http.StatusBadRequest, "invalid form param is empty")
  return
}
log.Debugf(ctx, "FormParams: %s", formParam)

// FirebaseAuthのユーザーIDを取得
userID := firebaseauth.GetUserID(ctx)
log.Debugf(ctx, "UserID: %s", userID)

// FirebaseAuthのJWTClaimsの値を取得
claims := firebaseauth.GetClaims(ctx)
log.Debugf(ctx, "Claims: %v", claims)

/****** HTTP ******/

func Get(ctx context.Context) error {
	// リクエストを送信
	status, body, err := httpclient.Get(ctx, "https://www.google.co.jp/", nil)
	if err != nil {
		log.Errorm(ctx, "httpclient.Get", err)
		return err
	}
	// HTTPステータスを確認
	if status != http.StatusOK {
		err := fmt.Errorf("http status: %d", status)
		return err
	}
	// Bodyをごにょごにょする
	str := util.BytesToStr(body)
	log.Debugf(ctx, "body length: %d", len(str))
	return nil
}
```
