GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello, run, deploy, browse, api

hello:
	@echo Hello go project ${GOPHER}

# 実行
run:
	dev_appserver.py ${svc}/app.yaml

# デプロイ
deploy-dev:
	@gcloud app deploy ${svc}/app.yaml

deploy-prod:
	@gcloud app deploy ${svc}/app.yaml

# 閲覧
browse:
	@gcloud app browse

# APIテスト
api:
	curl -s -X POST --data '{"jsonrpc":"2.0","id":1,"method":"${method}","params":"${params}"}' -H 'Authorization: Bearer ${token}' http://localhost:8080/api/v1/rpc | jq .
