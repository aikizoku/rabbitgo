package common

import (
	"fmt"
)

// Env ... 環境変数ファイルの定義
type Env struct {
	Credentials EnvData `json:"credentials"`
	Appengine   EnvData `json:"appengine"`
	Functions   EnvData `json:"functions"`
}

// EnvData ... 環境変数ファイルの環境毎のデータの定義
type EnvData struct {
	Local      map[string]interface{} `json:"local,omitempty"`
	Staging    map[string]interface{} `json:"staging"`
	Production map[string]interface{} `json:"production"`
}

// ProjectIDs ... 各環境のProjectIDの定義
type ProjectIDs struct {
	Local      string
	Staging    string
	Production string
}

// GetByEnv ... 指定した環境のProjectIDを取得する
func (m *ProjectIDs) GetByEnv(env string) string {
	switch env {
	case Local:
		return m.Local
	case Staging:
		return m.Staging
	case Production:
		return m.Production
	default:
		panic(fmt.Errorf("invalid env: %s", env))
	}
}
