GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello, run, deploy, api

hello:
	@echo Hello go project ${GOPHER}

# 実行
run:
	dev_appserver.py ${svc}/app.yaml

# デプロイ
deploy:
	@gcloud app deploy ${svc}/app.yaml

# APIテスト
api:
	curl -s -X POST --data '{"jsonrpc":"2.0","id":1,"method":"${method}","params":"${params}"}' -H 'Authorization: Bearer ${token}' http://localhost:8080/api/v1/rpc | jq .
