package worker

import (
	"net/http"

	"github.com/aikizoku/merlin/src/lib/log"
	"github.com/aikizoku/merlin/src/lib/renderer"
)

// AdminHandler ... 管理用のハンドラ
type AdminHandler struct {
}

// MigrateMasterData ... マスターデータを作成する
func (h *AdminHandler) MigrateMasterData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ここでマスターデータを作成する処理を入れる
	log.Debugf(ctx, "call migrate master data handler")

	renderer.Success(w)
}

// MigrateTestData ... テストデータを作成する
func (h *AdminHandler) MigrateTestData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ここでテストデータを作成する処理を入れる
	log.Debugf(ctx, "call migrate test data handler")

	renderer.Success(w)
}

// NewAdminHandler ... AdminHandlerを作成する
func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}
