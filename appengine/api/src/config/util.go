package config

import (
	"fmt"

	"github.com/rabee-inc/go-pkg/deploy"
)

func GetFilePath(path string) string {
	if deploy.IsLocal() {
		return fmt.Sprintf("./%s", path)
	} else {
		return fmt.Sprintf("./appengine/api/%s", path)
	}
}
