package basicauth

import (
	"fmt"
	"os"
	"strings"
)

// BasicAuthConfig ... ベーシック認証の設定
type BasicAuthConfig struct {
	Accounts map[string]string
}

// GetBasicAuthConfig ... ベーシック認証の設定を読み込む
func GetBasicAuthConfig(aKeys []string) *BasicAuthConfig {
	if len(aKeys) == 0 {
		panic(fmt.Errorf("no param basic auth account keys"))
	}

	acs := map[string]string{}
	for _, aKey := range aKeys {
		aKey = strings.ToUpper(aKey)

		uKey := fmt.Sprintf("BASIC_AUTH_%s_USER", aKey)
		u := os.Getenv(uKey)
		if u == "" {
			panic(fmt.Errorf("no config basic auth user: account_key=%s", aKey))
		}

		pKey := fmt.Sprintf("BASIC_AUTH_%s_PASSWORD", aKey)
		p := os.Getenv(pKey)
		if p == "" {
			panic(fmt.Errorf("no config basic auth password: account_key=%s", aKey))
		}

		acs[u] = p
	}

	return &BasicAuthConfig{
		Accounts: acs,
	}
}
