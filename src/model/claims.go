package model

// Claims ... JWT認証のClaims
type Claims struct {
	// 個人情報等、リクエスト内で見せたくない情報を扱う場合はここに入れる
	Sample string
}

// SetMap ... mapから取得する
func (m *Claims) SetMap(cmap map[string]interface{}) {
}

// ToMap ... mapで出力する
func (m *Claims) ToMap() map[string]interface{} {
	cmap := map[string]interface{}{}
	return cmap
}
