package main

import (
	"flag"
	"fmt"

	"github.com/aikizoku/skgo/test/config"
	"github.com/aikizoku/skgo/test/repository"
	"github.com/aikizoku/skgo/test/scenario"
	"github.com/aikizoku/skgo/test/service"
)

func main() {
	// Args
	snro := flag.String("scenario", "normal", "test scenario name")
	url := flag.String("url", "http://localhost:8080", "api endpoint url")
	auth := flag.String("auth", "", "authorization header value")
	flag.Parse()

	// Dependency
	d := &Dependency{}
	d.Inject(*snro, *url, *auth)

	// Execute
	d.Scenario.Run()
}

// Dependency ... 依存性
type Dependency struct {
	Scenario scenario.Interfaces
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject(snro string, apiURL string, authToken string) {
	// Repository
	fRepo := repository.NewFile(config.DocumentDirPath)
	hRepo := repository.NewHTTPClient()
	tRepo := repository.NewTemplateClient()

	// Service
	dSvc := service.NewDocument(fRepo, tRepo)
	rSvc := service.NewRest(
		hRepo,
		apiURL,
		map[string]string{
			"Authorization": fmt.Sprintf("%s%s", config.AuthorizationPrefix, authToken),
		},
		config.StagingURL,
		config.ProductionURL)
	jSvc := service.NewJSONRPC2()

	// Scenario
	switch snro {
	case "normal":
		d.Scenario = scenario.NewNormal(dSvc, rSvc, jSvc)
	case "abnormal":
		d.Scenario = scenario.NewAbnormal(dSvc, rSvc, jSvc)
	default:
		panic(fmt.Errorf("invalid scenario: %s", snro))
	}
}
