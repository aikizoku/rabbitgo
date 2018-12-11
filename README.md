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

# Goプロジェクトを取得(例:beegoの場合)
ghq get git@github.com:aikizoku/beego.git
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

# プロジェクトの切り替え
gcloud config set project <your-project-id>
```

## 依存パッケージのインストール
```bash
# バージョン管理ツールのインストール
brew install dep

# 依存パッケージのインストール
dep ensure
```

# 初期化
```bash
make init
```

# 動かす
## 起動
```bash
# API
make run app=api
make run-production app=api

# Worker
make run app=worker
make run-production app=worker
```

## ローカルで確認
```
// アプリを確認
http://localhost:8080/ping

// 各種データの確認
http://localhost:8000/instances
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

# Dispatch
make deploy-dispatch
make deploy-dispatch-production

# Index
make deploy-index
make deploy-index-production

# Queue
make deploy-queue
make deploy-queue-production
```

# 開発で使う便利なコマンド集
```bash
### Go ###
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

### ghq ###
# Goプロジェクトを取得(例:beegoの場合)
ghq get git@github.com:aikizoku/beego.git

### Google Cloud SDK ###
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

### dep ###
# 依存パッケージのインストール
dep ensure

# 依存パッケージのアップデート
dep ensure update

# 依存パッケージの追加
dep ensure -add <package-name>
```

# よく使うコード
```golang
/* REST Handler テンプレ */

// XXXXHandler ... XXXXのハンドラ
type XXXXHandler struct {
}

// TestHTTP ... HTTPテスト
func (h *XXXXHandler) Hoge(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	handler.RenderSuccess(w)
}

func (h *XXXXHandler) handleError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Errorf(ctx, msg)
	handler.RenderError(w, status, msg)
}

// NewXXXXHandler ... XXXXHandlerを作成する
func NewXXXXHandler() *XXXXHandler {
	return &XXXXHandler{}
}

/* JSONRPC2 Handler テンプレ */

// dependency.go
d.JSONRPC2 = jsonrpc2.NewMiddleware()
d.JSONRPC2.Register("sample", api.NewSampleJSONRPC2Handler(svc))

// XXXXHandler ... XXXXのハンドラ
type XXXXHandler struct {
}

type xxxxParams struct {
	Hoge string `json:"hoge"`
}

type xxxxResponse struct {
  Fuga string `json:"fuga"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *XXXXHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params xxxxParams
	err := json.Unmarshal(*msg, &params)
	return params, err
}

// Exec ... 処理をする
func (h *XXXXHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	// パラメータを取得
	hoge := params.(xxxxParams).Hoge
	log.Debugf(ctx, hoge)

	return xxxxResponse{
		Fuga: "",
	}, nil
}

// NewXXXXHandler ... XXXXHandlerを作成する
func NewXXXXHandler() *XXXXHandler {
	return &XXXXHandler{}
}

/* Service テンプレ */

//// interfaces
// XXXX ... XXXXのサービス
type XXXX interface {
}

//// implementation
type xxxx struct {
}

// NewXXXX ... XXXXを取得する
func NewXXXX() XXXX {
	return &xxxx{}
}

/* Repository テンプレ */

//// interfaces
// XXXX ... XXXXのリポジトリ
type XXXX interface {
}

//// implementation
type xxxx struct {
}

// NewXXXX ... XXXXを取得する
func NewXXXX() XXXX {
	return &xxxx{}
}

/* Middleware テンプレ */

// XXXX ... XXXXのミドルウェア
func XXXX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

/* リクエストの値を取得 */

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
```
