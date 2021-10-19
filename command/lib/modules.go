package lib

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// GetProjectID ... プロジェクトIDを取得する
func GetProjectID(deploy string) string {
	// プロジェクト
	file, err := ioutil.ReadFile("../../project.json")
	if err != nil {
		panic(err)
	}
	prj := map[string]string{}
	err = json.Unmarshal(file, &prj)
	if err != nil {
		panic(err)
	}
	return prj[deploy]
}

// NewAuthClient ... Authクライアントを作成
func NewAuthClient(env string) *auth.Client {
	var path string
	if env == Production {
		path = "../lib/credentials_production.json"
	} else {
		path = "../lib/credentials_staging.json"
	}
	ctx := context.Background()
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	cli, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}
	return cli
}

// GetEnv ... 実行環境を取得する
func GetEnv() string {
	var text string
	for {
		fmt.Print(fmt.Sprintf("%s(default) or %s? :", Staging, Production))
		scanner := bufio.NewScanner(os.Stdin)

		// 入力待ち
		scanner.Scan()

		// 入力値判定
		text = scanner.Text()
		switch text {
		case "", Staging:
			return Staging
		case Production:
			return Production
		default:
			fmt.Println(fmt.Sprintf("Please input \"%s\" or \"%s\"", Staging, Production))
		}
	}
}
