package images

import (
	"context"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/cloudfirestore"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/cloudpubsub"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
)

// Client ... クライアント
type Client struct {
	psCli     *cloudpubsub.Client
	topicName string
}

// SendConvertRequest ... 画像変換リクエストを送信する
func (c *Client) SendConvertRequest(
	ctx context.Context,
	sourceID string,
	sourceURL string,
	outPath string,
	docRefs []*cloudfirestore.DocRef,
	fieldName string) error {
	src := &ConvRequest{
		SourceID:  sourceID,
		SourceURL: sourceURL,
		OutPath:   outPath,
		DocRefs:   docRefs,
		FieldName: fieldName,
	}
	err := c.psCli.Publish(ctx, c.topicName, src)
	if err != nil {
		log.Errorm(ctx, "c.psCli.Publish", err)
		return err
	}
	return nil
}

// NewClient ... クライアントを作成する
func NewClient(psCli *cloudpubsub.Client, topicName string) *Client {
	return &Client{
		psCli:     psCli,
		topicName: topicName,
	}
}
