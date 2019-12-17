package images

import (
	"net/http"

	"github.com/andreiavrammsd/validator"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/errcode"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/parameter"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/renderer"
)

// Handler ... スポットのハンドラ
type Handler struct {
	repo Repository
}

// UpdateObjects ... 変換後の画像をアップデートする
func (h *Handler) UpdateObjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Paramを取得
	var param struct {
		Key     string    `json:"key"     validate:"required"`
		Objects []*Object `json:"objects" validate:"required"`
	}
	err := parameter.GetJSON(r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, "parameter.GetJSON", err)
		return
	}

	// Validation
	v := validator.New()
	if err := v.Struct(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, "v.Struct", err)
		return
	}

	err = h.repo.UpdateObjects(ctx, param.Key, param.Objects)
	if err != nil {
		renderer.HandleError(ctx, w, "h.sSvc.UpdateObjects", err)
		return
	}

	// Response
	renderer.Success(ctx, w)
}

// NewHandler ... ハンドラを作成する
func NewHandler(repo Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
