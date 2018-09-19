package service

import (
	"context"

	"github.com/aikizoku/beego/src/model"
)

// Beego ... サービスのインターフェースを定義
type Beego interface {
	Beego(ctx context.Context) (model.Beego, error)
}
