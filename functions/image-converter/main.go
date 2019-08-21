package p

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/nfnt/resize"
	"google.golang.org/api/option"
)

// PubSubMessage ... PubSubから受け取るメッセージ
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// Handle ... Functionsで実行される関数
func Handle(ctx context.Context, m PubSubMessage) error {
	InitLog()

	// 環境変数を取得
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentials == "" {
		return fmt.Errorf("env GOOGLE_APPLICATION_CREDENTIALS is not found")
	}
	backetName := os.Getenv("BACKET_NAME")
	if backetName == "" {
		return fmt.Errorf("env BACKET_NAME is not found")
	}

	// Presetsを取得
	presetsRaw, err := ioutil.ReadFile("./presets.json")
	if err != nil {
		handleError(err)
		return err
	}
	var presets []*Preset
	err = json.Unmarshal(presetsRaw, &presets)
	if err != nil {
		handleError(err)
		return err
	}

	// パラメータを取得
	var param Param
	if err := json.Unmarshal(m.Data, &param); err != nil {
		handleError(err)
		return err
	}

	// Validation
	if err := validate(param); err != nil {
		handleError(err)
		return err
	}

	// 画像を取得
	resp, err := http.Get(param.SourceURL)
	if err != nil {
		handleError(err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("image http get status: %d, url: %s", resp.StatusCode, param.SourceURL)
		handleError(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handleError(err)
		return err
	}

	// 画像を解析
	img, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		handleError(err)
		return err
	}
	imgConf, ext, err := image.DecodeConfig(bytes.NewReader(body))
	if err != nil {
		handleError(err)
		return err
	}
	var contentType string
	switch ext {
	case "jpeg", "jpg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	case "gif":
		contentType = "image/gif"
	default:
		err = fmt.Errorf("invalid ext: %s", ext)
		handleError(err)
		return err
	}

	// ドミナントカラーを取得
	dmColor := getDominantColor(img)
	dmColorStr := fmt.Sprintf("#%x%x%x", dmColor.R, dmColor.G, dmColor.B)

	// GCS準備
	gcsCli, err := storage.NewClient(ctx, option.WithCredentialsFile(credentials))

	// オリジナル画像をアップロード
	orURL, err := upload(
		ctx,
		gcsCli,
		body,
		contentType,
		backetName,
		fmt.Sprintf("%s/%s", param.OutPath, param.SourceID))
	if err != nil {
		handleError(err)
		return err
	}

	// 画像をプリセット毎にリサイズ
	sizes := map[string]*Size{}
	for _, preset := range presets {
		// 大きいサイズを取得
		var width, height uint
		var nWidth, nHeight float64
		if imgConf.Width > imgConf.Height {
			width = uint(preset.Width)
			height = 0
			nWidth = float64(preset.Width)
			nHeight = float64(imgConf.Height) * (float64(preset.Width) / float64(imgConf.Width))
		} else {
			width = 0
			height = uint(preset.Height)
			nWidth = float64(imgConf.Width) * (float64(preset.Height) / float64(imgConf.Height))
			nHeight = float64(preset.Height)
		}

		// 画像をリサイズ
		rImg := resize.Resize(width, height, img, resize.Bilinear)

		// 画像を作成
		var aImg []byte
		switch ext {
		case "jpeg", "jpg":
			// jpg画像を作成
			aImgBuf := new(bytes.Buffer)
			err = jpeg.Encode(aImgBuf, rImg, &jpeg.Options{
				Quality: 80,
			})
			if err != nil {
				LogErrorf(err.Error())
				return err
			}
			aImg = aImgBuf.Bytes()
		case "png":
			// png画像を作成
			aImgBuf := new(bytes.Buffer)
			err = png.Encode(aImgBuf, rImg)
			if err != nil {
				LogErrorf(err.Error())
				return err
			}
			aImg = aImgBuf.Bytes()
		case "gif":
			// gif画像を作成
			aImgBuf := new(bytes.Buffer)
			err = gif.Encode(aImgBuf, rImg, nil)
			if err != nil {
				LogErrorf(err.Error())
				return err
			}
			aImg = aImgBuf.Bytes()
		default:
			err = fmt.Errorf("invalid ext: %s", ext)
			handleError(err)
			return err
		}

		// リサイズ画像をアップロード
		url, err := upload(
			ctx,
			gcsCli,
			aImg,
			contentType,
			backetName,
			fmt.Sprintf("%s/%s_%s", param.OutPath, param.SourceID, preset.Name))
		if err != nil {
			handleError(err)
			return err
		}
		size := &Size{
			URL:    url,
			Width:  int(math.Ceil(nWidth)),
			Height: int(math.Ceil(nHeight)),
		}
		sizes[preset.Name] = size
	}
	obj := &Object{
		ID:            param.SourceID,
		URL:           orURL,
		DominantColor: dmColorStr,
		Sizes:         sizes,
	}

	// Firestoreに保存
	fApp, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credentials))
	if err != nil {
		handleError(err)
		return err
	}
	fCli, err := fApp.Firestore(ctx)
	if err != nil {
		handleError(err)
		return err
	}
	docRef := generateDocumentRef(fCli, param.DocRefs)
	err = fCli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		dsnp, err := tx.Get(docRef)
		if dsnp != nil && !dsnp.Exists() {
			return nil
		}
		if err != nil {
			handleError(err)
			return err
		}
		err = tx.Update(docRef, []firestore.Update{
			{Path: param.FieldName, Value: obj},
		})
		if err != nil {
			handleError(err)
			return err
		}
		return nil
	})
	if err != nil {
		handleError(err)
		return err
	}
	return nil
}
