package environment

// Project ... プロジェクト
type Project struct {
	Local      string `json:"local"`
	Staging    string `json:"staging"`
	Production string `json:"production"`
}

// Variable ... 値
type Variable struct {
	Local      map[string]string `json:"local"`
	Staging    map[string]string `json:"staging"`
	Production map[string]string `json:"production"`
}
