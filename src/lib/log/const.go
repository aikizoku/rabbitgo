package log

import "fmt"

// Level ... ログレベル
type Level int

const (
	// LevelDebug ... ログレベルDebug
	LevelDebug Level = 100
	// LevelInfo ... ログレベルInfo
	LevelInfo Level = 200
	// LevelWarning ... ログレベルWarning
	LevelWarning Level = 400
	// LevelError ... ログレベルError
	LevelError Level = 500
	// LevelCritical ... ログレベルCritical
	LevelCritical Level = 600
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelCritical:
		return "CRITICAL"
	}
	panic(fmt.Sprintf("Illegal level: %d", l))
}

// NewLevel ... Levelを作成する
func NewLevel(s string) Level {
	switch s {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARNING":
		return LevelWarning
	case "ERROR":
		return LevelError
	case "CRITICAL":
		return LevelCritical
	}
	panic(fmt.Sprintf("Illegal string: %s", s))
}
