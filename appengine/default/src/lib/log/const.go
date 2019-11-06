package log

import (
	"encoding/json"
	"fmt"
	"time"
)

// Severity ... ログレベル
// spec: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity
type Severity int

const (
	// SeverityDefault ... ログレベル: Default
	SeverityDefault Severity = 0
	// SeverityDebug ... ログレベル: Debug
	SeverityDebug Severity = 100
	// SeverityInfo ... ログレベル: Info
	SeverityInfo Severity = 200
	// SeverityWarning ... ログレベル: Warning
	SeverityWarning Severity = 400
	// SeverityError ... ログレベル: Error
	SeverityError Severity = 500
	// SeverityCritical ... ログレベル: Critical
	SeverityCritical Severity = 600
)

func (c Severity) String() string {
	switch c {
	case SeverityDefault:
		return "DEFAULT"
	case SeverityDebug:
		return "DEBUG"
	case SeverityInfo:
		return "INFO"
	case SeverityWarning:
		return "WARNING"
	case SeverityError:
		return "ERROR"
	case SeverityCritical:
		return "CRITICAL"
	default:
		panic(fmt.Sprintf("invalid log Severity: %d", c))
	}
}

// NewSeverity ... Severityを作成する
func NewSeverity(s string) Severity {
	switch s {
	case "DEFAULT":
		return SeverityDefault
	case "DEBUG":
		return SeverityDebug
	case "INFO":
		return SeverityInfo
	case "WARNING":
		return SeverityWarning
	case "ERROR":
		return SeverityError
	case "CRITICAL":
		return SeverityCritical
	default:
		panic(fmt.Sprintf("invalid log string: %s", s))
	}
}

// Time ... ProtocolBufferのフォーマット
type Time time.Time

// MarshalJSON ... JSONに変換
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339Nano))
}

var _ json.Marshaler = Duration(0)

// Duration ... ProtocolBufferのフォーマット
type Duration time.Duration

// MarshalJSON ... JSONに変換
func (d Duration) MarshalJSON() ([]byte, error) {
	nanos := time.Duration(d).Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	v := make(map[string]interface{})
	v["seconds"] = int64(secs)
	v["nanos"] = int32(nanos)
	return json.Marshal(v)
}
