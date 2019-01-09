package ogp

import (
	"context"
	"net/http"
	"strings"

	"github.com/aikizoku/gocci/src/lib/httpclient"
	"github.com/aikizoku/gocci/src/lib/log"
	"github.com/aikizoku/gocci/src/lib/util"
	"github.com/dyatlov/go-opengraph/opengraph"
	"golang.org/x/sync/errgroup"
)

// Get ... OG情報を取得する
func Get(ctx context.Context, url string) (*OpenGraph, error) {
	// 対象のHTMLを取得
	status, body, err := httpclient.Get(ctx, url, nil)
	if err != nil {
		log.Warningm(ctx, "httpclient.Get", err)
		return nil, err
	}
	if status != http.StatusOK {
		log.Warningf(ctx, "httpclient.Get error: status=%s, url=%s", status, url)
		return nil, err
	}
	html := util.BytesToStr(body)

	// 解析
	obj := opengraph.NewOpenGraph()
	err = obj.ProcessHTML(strings.NewReader(html))
	if err != nil {
		log.Warningm(ctx, "client.ProcessHTML", err)
		return nil, err
	}

	// 必要な情報のみを取得
	ogImgs := []*OpenGraphImage{}
	for _, image := range obj.Images {
		ogImg := &OpenGraphImage{
			URL:       image.URL,
			SecureURL: image.SecureURL,
			Type:      image.Type,
			Width:     image.Width,
			Height:    image.Height,
		}
		ogImgs = append(ogImgs, ogImg)
	}
	og := &OpenGraph{
		Type:        obj.URL,
		URL:         obj.URL,
		Title:       obj.Title,
		Description: obj.Description,
		SiteName:    obj.SiteName,
		Images:      ogImgs,
	}
	return og, nil
}

// GetMulti ... OG情報を複数取得する
func GetMulti(ctx context.Context, urls []string) ([]*OpenGraph, error) {
	var ogs []*OpenGraph
	eg := errgroup.Group{}
	for _, url := range urls {
		url := url
		eg.Go(func() error {
			og, err := Get(ctx, url)
			if err != nil {
				log.Errorm(ctx, "Get", err)
				return err
			}
			ogs = append(ogs, og)
			return nil
		})
	}
	_ = eg.Wait()
	return ogs, nil
}
