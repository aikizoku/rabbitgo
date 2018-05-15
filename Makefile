GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello, run, deploy, browse, api

hello:
	@echo Hello go project ${GOPHER}

# セットアップ
setup:
	dep 

# 実行
run-:
	@go run 

# デプロイ
deploy-dev:
	gcloud app deploy

deploy-prod:
	gcloud app deploy

# 閲覧
browse:
	gcloud app browse

# APIテスト
init

api:
	@echo curl -s -X POST --data '{"jsonrpc":"2.0","id":1,"method":"${method}","params":"${params}"}' -H 'Authorization: Bearer ${token}' http://localhost:8080/v1/rpc | jq .
