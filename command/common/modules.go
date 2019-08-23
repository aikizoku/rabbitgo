package common

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"google.golang.org/api/option"
)

// LoadEnvFile ... 環境変数ファイルを読み込む
func LoadEnvFile(deploy string) Env {
	raw, err := ioutil.ReadFile("../env.json")
	if err != nil {
		panic(err)
	}
	src := map[string]interface{}{}
	err = json.Unmarshal(raw, &src)
	if err != nil {
		panic(err)
	}
	apps := src["appengine"].(map[string]interface{})["apps"].([]interface{})
	appSts := []string{}
	for _, app := range apps {
		appSts = append(appSts, app.(string))
	}
	var dst Env
	dst.Apps = appSts
	dst.Credentials = src["credentials"].(map[string]interface{})[deploy].(map[string]interface{})
	if env, ok := src["appengine"].(map[string]interface{})[deploy]; ok {
		dst.Appengine = env.(map[string]interface{})
	}
	if env, ok := src["functions"].(map[string]interface{})[deploy]; ok {
		dst.Functions = env.(map[string]interface{})
	}
	return dst
}

// CreateFile ... 任意の場所に任意のファイルを作成してデータを書き込む
func CreateFile(path string, text string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	fmt.Fprintln(file, text)
}

// WriteFile ... 任意のファイルを開いてデータを書き込む
func WriteFile(path string, text string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	fmt.Fprintln(file, text)
}

// ReplaceFile ... 任意のファイルを開いてデータを置換する
func ReplaceFile(path string, old string, new string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		file.Close()
		panic(err.Error())
	}
	file.Close()

	rData := strings.Replace(string(data), old, new, -1)

	file, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	fmt.Fprintf(file, rData)
}

// ExecCommand ... 任意のコマンドを実行して結果を出力する
func ExecCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err.Error())
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err.Error())
	}

	cmd.Start()

	go PrintOutput(stdout)
	go PrintOutput(stderr)

	cmd.Wait()
}

// PrintOutput ... 任意の出力をフックする
func PrintOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// NewFirestoreClient ... Firestoreのクライアントを取得する
func NewFirestoreClient(env Env) *firestore.Client {
	b, err := json.Marshal(env.Credentials)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	opt := option.WithCredentialsJSON(b)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	return client
}

// NewAlgoliaClient ... Algoliaのクライアントを取得する
func NewAlgoliaClient(env Env) *search.Client {
	client := search.NewClientWithConfig(search.Configuration{
		AppID:  env.Appengine["ALGOLIA_APP_ID"].(string),
		APIKey: env.Appengine["ALGOLIA_API_KEY"].(string),
	})
	return client
}
