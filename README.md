# これはなに？
goをGAEで色々やるためのテンプレート

APIとWebとTaskが作れるよ

# Goのセットアップ
goenv(Goのバージョン管理)のインストール
```bash
brew install goenv
```

インストール可能なバージョンを確認
```bash
goenv install -l
```

バージョンを指定してインストール
```bash
goenv install 1.8.7
```

バージョン切り替え
```bash
goenv global 1.8.7
```

バージョン確認
```bash
go version
```

GOPATHを通す
```bash
vi .bash_profile

export GOPATH=$HOME/go

source .bash_profile
```

# ghq(リポジトリ管理)のセットアップ
インストール
```bash
brew install ghq
```

設定
```bash
git config --global ghq.root $GOPATH/src
```

Goプロジェクトを取得
```bash
ghq get xxxxxxxxx
```

# dep(パッケージ管理)のセットアップ
インストール
```bash
brew install dep
```


# 動かす






