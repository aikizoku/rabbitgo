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
	AcLog     *search.Client
	AtCli     *auth.Client
}

// NewDependency ... 依存性を設定
func NewDependency(env string) *Dependency {
	// ProjectID
	projectID := GetProjectID(env)

	// Firestore
	cFirestore := cloudfirestore.NewClient(projectID)

	// PubSub
	cPubsub := cloudpubsub.NewClient(projectID, []string{
		images.ConverterTopicID,
	})

	// ImageConverter
	cImages := images.NewClient(cPubsub)

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
	acLog := search.NewClientWithConfig(search.Configuration{
		AppID:  apiKey,
		APIKey: apiSecret,
	})

	// Auth
	atCli := NewAuthClient(env)

	return &Dependency{
		ProjectID: projectID,
		FCli:      cFirestore,
		PsCli:     cPubsub,
		ImgCli:    cImages,
		AcLog:     acLog,
		AtCli:     atCli,
	}
}
