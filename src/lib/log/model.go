package log

// Entry ... ログの内容
type Entry struct {
	Method    string
	URL       string
	UserAgent string
	Level     Level
	File      string
	Line      int
	Message   string
}
