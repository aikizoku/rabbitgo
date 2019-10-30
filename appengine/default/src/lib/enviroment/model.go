package enviroment

// Object ... 環境変数
type Object struct {
	Credentials struct {
		Local      *Credentials `json:"local"`
		Staging    *Credentials `json:"staging"`
		Production *Credentials `json:"production"`
	} `json:"credentials"`
	Appengine map[string]*struct {
		Local      map[string]string `json:"local"`
		Staging    map[string]string `json:"staging"`
		Production map[string]string `json:"production"`
	} `json:"appengine"`
	Functions map[string]*struct {
		Local      map[string]interface{} `json:"local"`
		Staging    map[string]interface{} `json:"staging"`
		Production map[string]interface{} `json:"production"`
	} `json:"functions"`
	Scheduler map[string]*struct {
		Local      map[string]string `json:"local"`
		Staging    map[string]string `json:"staging"`
		Production map[string]string `json:"production"`
	} `json:"scheduler"`
}

// Credentials ... GCPの認証情報
type Credentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}
