GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello, run, deploy, api

hello:
	@echo Hello go project ${GOPHER}

# 実行
run:
	dev_appserver.py gae/${s}/app.yaml

# デプロイ
deploy:
	@gcloud app deploy gae/${s}/app.yaml

# ドメイン設定をデプロイ
deploy-domain:
	@gcloud app deploy gae/dispatch.yaml --project pj-trial-id

# Datastoreの複合インデックス定義をデプロイ
deploy-index:
	@gcloud app deploy gae/index.yaml

# APIテスト
api:
	curl -s -X POST --data '{"jsonrpc":"2.0","id":"1","method":"${m}","params":"${p}"}' -H 'Content-Type: application/json' -H 'Authorization: Bearer ${t}' http://localhost:8080/api/v1/rpc | jq .
