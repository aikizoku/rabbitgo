package worker

import (
	"net/http"

	"github.com/aikizoku/beego/src/handler"
	"google.golang.org/appengine"
	"github.com/aikizoku/beego/src/lib/log"
)

// AdminHandler ... 管理用のハンドラ
type AdminHandler struct {
}

// MigrateMasterData ... マスターデータを作成する
func (h *AdminHandler) MigrateMasterData(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// ここでマスターデータを作成する処理を入れる
	log.Debugf(ctx, "call migrate master data handler")

	handler.RenderSuccess(w)
}

// MigrateTestData ... テストデータを作成する
func (h *AdminHandler) MigrateTestData(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// ここでテストデータを作成する処理を入れる
	log.Debugf(ctx, "call migrate test data handler")

	handler.RenderSuccess(w)
}

// NewAdminHandler ... AdminHandlerを作成する
func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}
