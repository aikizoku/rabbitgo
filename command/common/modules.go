package common

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

// LoadEnvFile ... 環境変数ファイルを読み込む
func LoadEnvFile() Env {
	raw, err := ioutil.ReadFile("./env.json")
	if err != nil {
		panic(err)
	}
	var env Env
	err = json.Unmarshal(raw, &env)
	if err != nil {
		panic(err)
	}
	return env
}

// GetProjectIDs ... 環境変数データからProjectIDを取得する
func GetProjectIDs(env Env) ProjectIDs {
	return ProjectIDs{
		Local:      env.Credentials.Local["project_id"],
		Staging:    env.Credentials.Staging["project_id"],
		Production: env.Credentials.Production["project_id"],
	}
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
