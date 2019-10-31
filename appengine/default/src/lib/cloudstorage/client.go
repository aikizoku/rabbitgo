package cloudstorage

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/vincent-petithory/dataurl"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/errcode"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
)

// Client ... GCSのクライアント
type Client struct {
	cli    *storage.Client
	bucket string
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
		Bucket(c.bucket).
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
	url := GenerateFileURL(c.bucket, path, name)
	return url, nil
}

// GetReader ... 指定ファイルのReaderを取得する
func (c *Client) GetReader(
	ctx context.Context,
	path string) (*storage.Reader, error) {
	reader, err := c.cli.
		Bucket(c.bucket).
		Object(path).
		NewReader(ctx)
	if err != nil {
		log.Errorm(ctx, "c.cli.NewReader", err)
		return nil, err
	}

	return reader, nil
}

// GetBucket ... バケット名
func (c *Client) GetBucket() string {
	return c.bucket
}

// NewClient ... クライアントを作成する
func NewClient(bucket string) *Client {
	ctx := context.Background()
	gOpt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                30 * time.Millisecond,
		Timeout:             20 * time.Millisecond,
		PermitWithoutStream: true,
	}))
	cli, err := storage.NewClient(ctx, gOpt)
	if err != nil {
		panic(err)
	}
	return &Client{
		cli:    cli,
		bucket: bucket,
	}
}
