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
	ProjectID     string
	CFirestore    *firestore.Client
	CPubsub       *cloudpubsub.Client
	CImages       *images.Client
	CAlgolia      *search.Client
	CFirebaseAuth *auth.Client
}

// NewDependency ... 依存性を設定
func NewDependency(env string) *Dependency {
	projectID := GetProjectID(env)
	cFirestore := cloudfirestore.NewClient(projectID)
	cPubsub := cloudpubsub.NewClient(projectID, []string{
		images.ConverterTopicID,
	})
	cImages := images.NewClient(cPubsub)
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
	cAlgolia := search.NewClientWithConfig(search.Configuration{
		AppID:  apiKey,
		APIKey: apiSecret,
	})
	cFirebaseAuth := NewAuthClient(env)

	return &Dependency{
		ProjectID:     projectID,
		CFirestore:    cFirestore,
		CPubsub:       cPubsub,
		CImages:       cImages,
		CAlgolia:      cAlgolia,
		CFirebaseAuth: cFirebaseAuth,
	}
}
