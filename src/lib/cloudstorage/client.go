package cloudstorage

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rabee-inc/nearby-backend/src/lib/errcode"
	"github.com/rabee-inc/nearby-backend/src/lib/log"
	"github.com/vincent-petithory/dataurl"
	"google.golang.org/api/option"
)

// Client ... GCSのクライアント
type Client struct {
	cli    *storage.Client
	backet string
}

// UploadForDataURL ... DataURLのファイルをアップロードする
func (c *Client) UploadForDataURL(
	ctx context.Context,
	path string,
	name string,
	cacheMode *CacheMode,
	dataURL string) (string, error) {
	// Base64をデコード
	res, err := dataurl.DecodeString(dataURL)
	if err != nil {
		log.Warningm(ctx, "dataurl.DecodeString", err)
		err = errcode.Set(err, http.StatusBadRequest)
		return "", err
	}

	// アップロード
	return c.Upload(ctx, path, name, res.ContentType(), cacheMode, res.Data)
}

// Upload ... ファイルをアップロードする
func (c *Client) Upload(
	ctx context.Context,
	path string,
	name string,
	contentType string,
	cacheMode *CacheMode,
	data []byte) (string, error) {
	// Writerを作成
	w := c.cli.
		Bucket(c.backet).
		Object(strings.Join([]string{path, name}, "/")).
		NewWriter(ctx)

	// ContentTypeを設定
	w.ContentType = contentType

	// Cache-Controllを設定
	if cacheMode != nil {
		var cc string
		if cacheMode.Disabled {
			cc = "no-cache"
		} else {
			cc = fmt.Sprintf("max-age=%d", cacheMode.Expire/time.Second)
		}
		w.CacheControl = cc
	}
	w.ChunkSize = ChunkSize

	// アップロード
	if _, err := w.Write(data); err != nil {
		log.Errorm(ctx, "w.Write", err)
		return "", err
	}
	if err := w.Close(); err != nil {
		log.Errorm(ctx, "w.Close", err)
		return "", err
	}

	// URLを作成
	url := GenerateFileURL(c.backet, path, name)
	return url, nil
}

// NewClient ... クライアントを作成する
func NewClient(credentialsPath string, backet string) *Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialsPath)
	cli, err := storage.NewClient(ctx, opt)
	if err != nil {
		panic(err)
	}
	return &Client{
		cli:    cli,
		backet: backet,
	}
}
