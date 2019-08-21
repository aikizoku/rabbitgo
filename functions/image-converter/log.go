package p

import "log"

// InitLog ... ログを初期化する
func InitLog() {
	log.SetFlags(log.Lshortfile)
}

// LogDebugf ... デバッグログを出力する
func LogDebugf(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args)
}

// LogInfof ... インフォログを出力する
func LogInfof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args)
}

// LogErrorf ... エラーログを出力する
func LogErrorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args)
}
