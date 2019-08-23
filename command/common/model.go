package common

// Env ... 環境変数ファイルの定義
type Env struct {
	Apps        []string
	Credentials map[string]interface{}
	Appengine   map[string]interface{}
	Functions   map[string]interface{}
}

// GetProjectID ... 環境変数データからProjectIDを取得する
func (m *Env) GetProjectID() string {
	return m.Credentials["project_id"].(string)
}
