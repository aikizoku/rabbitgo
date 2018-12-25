package repository

// TemplateClient ... テンプレートファイルに関するリポジトリ
type TemplateClient interface {
	GetMarged(path string, src interface{}) string
}
