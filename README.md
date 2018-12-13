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

# Goプロジェクトを取得(例:beegoの場合)
ghq get git@github.com:aikizoku/beego.git

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
/****** REST Handler ******/

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

/****** JSONRPC2 Handler ******/

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

/****** Service ******/

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

/****** Datastore ******/

import (
	_ "go.mercari.io/datastore/aedatastore" // mercari/datastoreの初期化
)

func Get(ctx context.Context, id int64) (*model.Xxxx, error) {
	dst := &model.Xxxx{
		ID: id,
	}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return dst, err
	}
	if err := b.Get(dst); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, err
		}
		log.Errorm(ctx, "b.Get", err)
		return nil, err
	}
	return dst, nil
}

func GetMulti(ctx context.Context, ids []int64) ([]*model.Xxxx, error) {
	dsts := []*model.Xxxx{}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return dsts, err
	}
	bt := b.Batch()
	for _, id := range ids {
		dst := &model.Xxxx{
			ID: id,
		}
		bt.Get(dst, func(err error) error {
			if err != nil {
				if err == datastore.ErrNoSuchEntity {
					return nil
				}
				log.Errorm(ctx, "bt.Get", err)
				return err
			}
			ret = append(ret, dst)
			return nil
		})
	}
	err = bt.Exec()
	if err != nil {
		log.Errorm(ctx, "bt.Exec", err)
		return dsts, err
	}
	return dsts, nil
}

func GetByQuery(ctx context.Context, userID string) ([]*model.Xxxx, error) {
	dsts := []*model.Xxxx{}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return dsts, err
	}
	q := b.NewQuery("Xxxx").
		Filter("UserID =", userID).
		Filter("Enabled =", true).
		Order("-CreatedAt").
		KeysOnly()
	keys, err := b.GetAll(q, nil)
	if err != nil {
		log.Errorm(ctx, "b.GetAll", err)
		return dsts, err
	}
	ids := []int64{}
	for _, key := range keys {
		ids = append(ids, key.ID())
	}
	dsts, err = GetMulti(ctx, ids)
	if err != nil {
		log.Errorm(ctx, "GetMulti", err)
		return dsts, err
	}
	return dsts, nil
}

func Upsert(ctx context.Context, src *model.Xxxx) (int64, error) {
	var id int64
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return id, err
	}
	key, err := b.Put(src)
	if err != nil {
		log.Errorm(ctx, "bt.Put", err)
		return id, err
	}
	id = key.ID()
	return id, nil
}

func UpsertMulti(ctx context.Context, srcs []*model.Xxxx) ([]int64, error) {
	ids := []int64{}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return ids, err
	}
	bt := b.Batch()
	for _, src := range srcs {
		bt.Put(src, func(key datastore.Key, err error) error {
			if err != nil {
				log.Errorm(ctx, "bt.Put", err)
				return err
			}
			ids = append(ids, key.ID())
			return nil
		})
	}
	err = bt.Exec()
	if err != nil {
		log.Errorm(ctx, "bt.Exec", err)
		return ids, err
	}
	return ids, nil
}

func Delete(ctx context.Context, id int64) error {
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return ret, err
	}
	src := &model.Xxxx{
		ID: id,
	}
	err = b.Delete(src)
	if err != nil {
		log.Errorm(ctx, "bt.Delete", err)
		return err
	}
	return nil
}

func DeleteMulti(ctx context.Context, ids []int64) error {
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return ret, err
	}
	bt := b.Batch()
	for _, id := range ids {
		src := &model.Xxxx{
			ID: id,
		}
		bt.Delete(src, func(err error) error {
			if err != nil {
				if err == datastore.ErrNoSuchEntity {
					return nil
				}
				log.Errorm(ctx, "bt.Delete", err)
				return err
			}
			return nil
		})
	}
	err = bt.Exec()
	if err != nil {
		log.Errorm(ctx, "bt.Exec", err)
		return err
	}
	return nil
}

/****** Firestore ******/



/****** CloudSQL ******/

func Get(ctx context.Context, id int64) (*model.Xxxx, error) {
	var dst *model.Xxxx
	q := sq.Select(
		"id",
		"category",
		"name",
		"enabled",
		"created_at",
		"updated_at").
		From("xxxx").
		Where(sq.Eq{
			"id":      id,
			"enabled": 1,
		})
	cloudsql.DumpSelectQuery(ctx, q)

	row := q.RunWith(r.csql).QueryRowContext(ctx)
	err := row.Scan(
		&ret.ID,
		&ret.Category,
		&ret.Name,
		&ret.Enabled,
		&ret.CreatedAt,
		&ret.UpdatedAt)
	if err != nil {
		log.Errorm(ctx, "q.RunWith.QueryRowContext", err)
		return dst, err
	}

	return dst, nil
}

func GetMulti(ctx context.Context, ids []int64) ([]*model.Xxxx, error) {
	var dsts []*model.Xxxx
	q := sq.Select(
		"id",
		"name",
		"category",
		"enabled",
		"created_at",
		"updated_at").
		From("sample").
		Where(sq.Eq{
			"id":      ids,
			"enabled": 1,
		})
	cloudsql.DumpSelectQuery(ctx, q)

	rows, err := q.RunWith(r.csql).QueryContext(ctx)
	if err != nil {
		log.Errorm(ctx, "q.RunWith.QueryContext", err)
		return dsts, err
	}

	for rows.Next() {
		var dst *model.Xxxx
		err := rows.Scan(
			&ret.ID,
			&ret.Name,
			&ret.Category,
			&ret.Enabled,
			&ret.CreatedAt,
			&ret.UpdatedAt)
		if err != nil {
			log.Errorm(ctx, "rows.Scan", err)
			rows.Close()
			return dsts, err
		}
		dsts = append(dsts, dst)
	}

	return dsts, nil
}

func Insert(ctx context.Context, src *model.Xxxx) error {
	now := util.TimeNowUnix()

	q := sq.Insert("xxxx").
		Columns("id", "category", "name", "enabled", "created_at", "updated_at").
		Values(src.ID, src.Category, src.Name, 1, now, now)
	cloudsql.DumpInsertQuery(ctx, q)

	_, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorm(ctx, "q.RunWith.ExecContext", err)
		return err
	}

	return nil
}

func Update(ctx context.Context, src *model.Xxxx) error {
	now := util.TimeNowUnix()

	q := sq.Update("xxxx").
		Set("name", src.Name).
		Set("category", src.Category).
		Set("enabled", src.Enabled).
		Set("updated_at", now).
		Where(sq.Eq{"id": src.ID})
	cloudsql.DumpUpdateQuery(ctx, q)

	res, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorm(ctx, "q.RunWith.ExecContext", err)
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = fmt.Errorf("no affected id = %d", obj.ID)
		log.Errorf(ctx, err.Error())
		return err
	}

	return nil
}

func Upsert(ctx context.Context, src *model.Xxxx) error {
	now := util.TimeNowUnix()

	q := sq.Insert("xxxx").
		Columns("id", "name", "category", "enabled", "created_at", "updated_at").
		Values(src.ID, src.Category, src.Name, 1, now, now).
		Suffix("ON DUPLICATE KEY UPDATE name = VALUES(name), updated_at = VALUES(updated_at)")
	cloudsql.DumpInsertQuery(ctx, q)

	_, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorm(ctx, "q.RunWith.ExecContext", err)
		return err
	}

	return nil
}

// 指定のレコードを削除する
func Delete(ctx context.Context, id int64) error {
	// クエリを作成
	q := sq.Delete("xxxx").Where(sq.Eq{"id": id})
	
	// デバッグ用にクエリを出力
	cloudsql.DumpDeleteQuery(ctx, q)

	// クエリを実行
	res, err := q.RunWith(r.csql).ExecContext(ctx)
	if err != nil {
		log.Errorm(ctx, "q.RunWith.ExecContext", err)
		return err
	}
	
	// 削除する対象が存在しなかった場合のハンドリング
	if affected, _ := res.RowsAffected(); affected == 0 {
		err = fmt.Errorf("no affected id = %d", id)
		log.Errorf(ctx, err.Error())
		return err
	}

	return nil
}

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
