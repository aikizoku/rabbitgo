package ogp

import (
	"context"
	"sync"

	"github.com/otiai10/opengraph"
	"golang.org/x/sync/errgroup"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
)

// Get ... OG情報を取得する
func Get(ctx context.Context, url string) (*OpenGraph, error) {
	var dst *OpenGraph

	// OGP取得
	og, err := opengraph.Fetch(url)
	if err != nil {
		log.Debugm(ctx, "opengraph.Fetch", err)
		return nil, err
	}

	// 必要な情報のみを取得
	var imageURL string
	if len(og.Image) > 0 {
		imageURL = og.Image[0].URL
	} else {
		imageURL = ""
	}
	dst = &OpenGraph{
		URL:         og.URL.Source,
		Title:       og.Title,
		Description: og.Description,
		ImageURL:    imageURL,
		FaviconURL:  og.Favicon,
	}
	return dst, nil
}

// GetMulti ... OG情報を複数取得する
func GetMulti(ctx context.Context, urls []string) ([]*OpenGraph, error) {
	ogs := []*OpenGraph{}
	eg := errgroup.Group{}
	mutex := &sync.Mutex{}
	for _, url := range urls {
		url := url
		eg.Go(func() error {
			og, err := Get(ctx, url)
			if err != nil {
				log.Debugm(ctx, "Get", err)
				return err
			}
			mutex.Lock()
			ogs = append(ogs, og)
			mutex.Unlock()
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Debugm(ctx, "eg.Wait", err)
		return ogs, err
	}
	return ogs, nil
}
