package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
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
