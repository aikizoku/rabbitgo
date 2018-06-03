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

# ディスパッチ設定をデプロイ
dispatch:
	@gcloud app deploy gae/dispatch.yaml --project pj-trial-id

# Cron設定をデプロイ
cron:
	@gcloud app deploy gae/task/cron.yaml --project pj-trial-id

# Queue設定をデプロイ
queue:
	@gcloud app deploy gae/task/queue.yaml --project pj-trial-id

# Datastoreの複合インデックス定義をデプロイ
index:
	@gcloud app deploy gae/index.yaml --project pj-trial-id

# APIテスト
api:
	curl -s -X POST --data '{"jsonrpc":"2.0","id":"1","method":"${m}","params":"${p}"}' -H 'Content-Type: application/json' -H 'Authorization: Bearer ${t}' http://localhost:8080/api/v1/rpc | jq .
