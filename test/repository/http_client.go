package repository

// HTTPClient ... HTTP通信を行うリポジトリ
type HTTPClient interface {
	GetForm(u string, params map[string]interface{}, headers map[string]string) (int, string)
	PostJSON(url string, json []byte, headers map[string]string) (int, string)
	PutJSON(url string, json []byte, headers map[string]string) (int, string)
	DeleteJSON(url string, json []byte, headers map[string]string) (int, string)
}
