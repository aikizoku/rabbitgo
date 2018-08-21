GOPHER = 'ʕ◔ϖ◔ʔ'
DEV_PROJECT_ID = 'beego-dev-thehero-jp'
PROD_PROJECT_ID = 'beego-prd-thehero-jp'

.PHONY: hello, run, deploy, dispatch, cron, queue, index

hello:
	@echo Hello go project ${GOPHER}

# 実行
run:
	dev_appserver.py gae/${s}/app.yaml

# デプロイ
deploy:
	@gcloud app deploy -q gae/${s}/app.yaml

deploy-prod:
	@gcloud app deploy -q gae/${s}/app_prod.yaml

# ディスパッチ設定をデプロイ
dispatch:
	@gcloud app deploy -q gae/config/dispatch.yaml --project ${DEV_PROJECT_ID}

dispatch-prod:
	@gcloud app deploy -q gae/config/dispatch_prod.yaml --project ${PROD_PROJECT_ID}

# Cron設定をデプロイ
cron:
	@gcloud app deploy -q gae/config/cron.yaml --project ${DEV_PROJECT_ID}

cron-prod:
	@gcloud app deploy -q gae/config/cron.yaml --project ${PROD_PROJECT_ID}

# Queue設定をデプロイ
queue:
	@gcloud app deploy -q gae/config/queue.yaml --project ${DEV_PROJECT_ID}

queue-prod:
	@gcloud app deploy -q gae/config/queue.yaml --project ${PROD_PROJECT_ID}

# Datastoreの複合インデックス定義をデプロイ
index:
	@gcloud app deploy -q gae/config/index.yaml --project ${DEV_PROJECT_ID}

index-prod:
	@gcloud app deploy -q gae/config/index.yaml --project ${PROD_PROJECT_ID}
