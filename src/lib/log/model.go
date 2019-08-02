package log

// Entry ... 構造ログ定義
// https://cloud.google.com/logging/docs/agent/configuration#special-fields
type Entry struct {
	Severity    string            `json:"severity"`
	HTTPRequest *EntryHTTPRequest `json:"httpRequest,omitempty"`
	Time        Time              `json:"time"`
	Trace       string            `json:"logging.googleapis.com/trace"`
	TraceID     string            `json:"traceId"`
	Childs      []*EntryChild     `json:"childs"`
	Message     string            `json:"message,omitempty"`
}

// EntryHTTPRequest ... HTTPリクエストの構造ログ定義
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#httprequest
type EntryHTTPRequest struct {
	RequestMethod                  string   `json:"requestMethod"`
	RequestURL                     string   `json:"requestUrl"`
	RequestSize                    int64    `json:"requestSize,string,omitempty"`
	Status                         int      `json:"status"`
	ResponseSize                   int64    `json:"responseSize,string,omitempty"`
	UserAgent                      string   `json:"userAgent,omitempty"`
	Referer                        string   `json:"referer,omitempty"`
	Latency                        Duration `json:"latency,omitempty"`
	CacheLookup                    *bool    `json:"cacheLookup,omitempty"`
	CacheHit                       *bool    `json:"cacheHit,omitempty"`
	CacheValidatedWithOriginServer *bool    `json:"cacheValidatedWithOriginServer,omitempty"`
	CacheFillBytes                 *int64   `json:"cacheFillBytes,string,omitempty"`
	Protocol                       string   `json:"protocol"`
}

// EntryChild ... 子ログの構造ログ定義
type EntryChild struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Time     Time   `json:"time"`
}

// EntrySourceLocation ... SourceLocationの構造ログ定義
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logentrysourcelocation
type EntrySourceLocation struct {
	File     string `json:"file,omitempty"`
	Line     int64  `json:"line,string,omitempty"`
	Function string `json:"function,omitempty"`
}
