GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello, run, deploy, api

hello:
	@echo Hello go project ${GOPHER}

# 実行
run:
	dev_appserver.py ${s}/app.yaml

# デプロイ
deploy:
	@gcloud app deploy ${s}/app.yaml

# APIテスト
api:
	curl -s -X POST --data '{"jsonrpc":"2.0","id":"1","method":"${m}","params":"${p}"}' -H 'Content-Type: application/json' -H 'Authorization: Bearer ${t}' http://localhost:8080/api/v1/rpc | jq .
