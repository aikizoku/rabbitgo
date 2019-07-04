# これはなに？
GAE/Go環境で動作するサーバー開発のテンプレート
とても早くて軽いAPI/Workerをワンソースで作る事が出来ます。
- インフラをあまり考えなくて良い
- 適当に作っていても循環参照が発生しない
- 緩い命名規則で縛っているので、柔軟かつ迷わない命名が可能
- 新しい機能を追加するときも本テンプレートの対応を待たないで素直に実装可能
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
goenv install 1.11.x

# バージョン切り替え
goenv global 1.11.x

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

# Goプロジェクトを取得(例:merlinの場合)
ghq get git@github.com:aikizoku/merlin.git
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
cd appengine/app/api
GO111MODULE=on go test
```

# 動かす
## 起動
```bash
# API
make run app=api
make run-staging app=api
make run-production app=api

# Worker
make run app=worker
make run-staging app=worker
make run-production app=worker
```

## ローカルで確認
```
// アプリを確認
http://localhost:8080/ping
```

## デプロイ
```bash
# API
make deploy app=api
make deploy-production app=api

# Worker
make deploy app=worker
make deploy-production app=worker

# Cron
make deploy-cron
make deploy-cron-production

# Queue
make deploy-queue
make deploy-queue-production
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
goenv install 1.11.1

# バージョン切り替え
goenv global 1.11.1

# バージョン確認
go version

###### ghq ######

# Goプロジェクトを取得(例:merlinの場合)
ghq get git@github.com:aikizoku/merlin.git

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

/****** Middleware ******/

// XXXX ... XXXXのミドルウェア
func XXXX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

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
