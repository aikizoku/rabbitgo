GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello, run, run-prod, deploy, deploy-prod, domain, domain-prod, index, index-prod, api

hello:
	@echo Hello go project ${GOPHER}

# 実行
run:
	dev_appserver.py gae/${s}/app_dev.yaml
run-prod:
	dev_appserver.py gae/${s}/app_prod.yaml	

# デプロイ
deploy:
	@gcloud app deploy gae/${s}/app_dev.yaml
deploy-prod:
	@gcloud app deploy gae/${s}/app_prod.yaml

# ドメイン設定をデプロイ
domain:
	@gcloud app deploy gae/dispatch_dev.yaml --project pj-trial-id
domain-prod:
	@gcloud app deploy gae/dispatch_prod.yaml --project pj-trial-id

# Datastoreの複合インデックス定義をデプロイ
index:
	@gcloud app deploy gae/index.yaml
index-prod:
	@gcloud app deploy gae/index.yaml

# APIテスト
api:
	curl -s -X POST --data '{"jsonrpc":"2.0","id":"1","method":"${m}","params":"${p}"}' -H 'Content-Type: application/json' -H 'Authorization: Bearer ${t}' http://localhost:8080/api/v1/rpc | jq .
