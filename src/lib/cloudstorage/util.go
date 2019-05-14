package cloudstorage

import "strings"

// GenerateFileURL ... GCSのファイルURLを作成する
func GenerateFileURL(backet string, path string, name string) string {
	return strings.Join([]string{BaseURL, backet, path, name}, "/")
}
