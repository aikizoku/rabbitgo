package repository

// File ... ファイル操作に関するリポジトリ
type File interface {
	CreateDir(path string)
	WriteFile(path string, body string)
	RemoveAll()
}
