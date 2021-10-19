package lib

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/cloudpubsub"
	"github.com/rabee-inc/go-pkg/images"
)

// Dependency ... 依存性
type Dependency struct {
	ProjectID string
	FCli      *firestore.Client
	PsCli     *cloudpubsub.Client
	ImgCli    *images.Client
	AlCli     *search.Client
	AtCli     *auth.Client
}

// NewDependency ... 依存性を設定
func NewDependency(env string) *Dependency {
	// ProjectID
	projectID := GetProjectID(env)

	// Firestore
	fCli := cloudfirestore.NewClient(projectID)

	// PubSub
	psCli := cloudpubsub.NewClient(projectID, []string{
		images.ConverterTopicID,
	})

	// ImageConverter
	imgCli := images.NewClient(psCli)

	// Algolia
	var apiKey, apiSecret string
	if env == Production {
		// 本番
		apiKey = ""
		apiSecret = ""
	} else {
		// ステージング
		apiKey = ""
		apiSecret = ""
	}
	alCli := search.NewClientWithConfig(search.Configuration{
		AppID:  apiKey,
		APIKey: apiSecret,
	})

	// Auth
	atCli := NewAuthClient(env)

	return &Dependency{
		ProjectID: projectID,
		FCli:      fCli,
		PsCli:     psCli,
		ImgCli:    imgCli,
		AlCli:     alCli,
		AtCli:     atCli,
	}
}
