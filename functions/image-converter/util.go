package p

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
)

func handleError(err error) {
	LogErrorf(err.Error())
}

func validate(param Param) error {
	if param.SourceID == "" {
		return errors.New("source_id is empty")
	}
	if param.SourceURL == "" {
		return errors.New("source_url is empty")
	}
	if param.OutPath == "" {
		return errors.New("out_path is empty")
	}
	for i, docRef := range param.DocRefs {
		if docRef.CollectionName == "" {
			return fmt.Errorf("docRef[%d].collection_name is empty", i)
		}
		if docRef.DocID == "" {
			return fmt.Errorf("docRef[%d].doc_id is empty", i)
		}
	}
	if param.FieldName == "" {
		return errors.New("field_name is empty")
	}
	return nil
}

func getDominantColor(img image.Image) color.RGBA {
	var r, g, b, count float64
	rect := img.Bounds()
	for i := 0; i < rect.Max.Y; i++ {
		for j := 0; j < rect.Max.X; j++ {
			c := color.RGBAModel.Convert(img.At(j, i))
			r += float64(c.(color.RGBA).R)
			g += float64(c.(color.RGBA).G)
			b += float64(c.(color.RGBA).B)
			count++
		}
	}
	return color.RGBA{
		R: uint8(r / count),
		G: uint8(g / count),
		B: uint8(b / count),
	}
}

func upload(
	ctx context.Context,
	gcsCli *storage.Client,
	img []byte,
	contentType string,
	backetName string,
	filePath string) (string, error) {
	gcsWri := gcsCli.
		Bucket(backetName).
		Object(filePath).
		NewWriter(ctx)
	gcsWri.ContentType = contentType
	if _, err := gcsWri.Write(img); err != nil {
		return "", err
	}
	if err := gcsWri.Close(); err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", backetName, filePath)
	return url, nil
}

func generateDocumentRef(fCli *firestore.Client, docRefs []*DocRef) *firestore.DocumentRef {
	var dst *firestore.DocumentRef
	for i, docRef := range docRefs {
		if i == 0 {
			dst = fCli.Collection(docRef.CollectionName).Doc(docRef.DocID)
		} else {
			dst = dst.Collection(docRef.CollectionName).Doc(docRef.DocID)
		}
	}
	return dst
}
