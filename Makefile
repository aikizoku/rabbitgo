GOPHER = 'ʕ◔ϖ◔ʔ'
STAGING_PROJECT_ID = 'beego-staging-thehero-jp'
PRODUCTION_PROJECT_ID = 'beego-thehero-jp'

hello:
	@echo Hello go project ${GOPHER}

# 実行
run:
	dev_appserver.py deploy/staging/${app}/app.yaml

run-production:
	dev_appserver.py deploy/production/${app}/app.yaml

# デプロイ
deploy-app:
	@gcloud app deploy -q deploy/staging/${app}/app.yaml

deploy-app-production:
	@gcloud app deploy -q deploy/production/${app}/app.yaml

# ディスパッチ設定をデプロイ
deploy-dispatch:
	@gcloud app deploy -q appengine/config/dispatch_staging.yaml --project ${STAGING_PROJECT_ID}

deploy-dispatch-production:
	@gcloud app deploy -q appengine/config/dispatch_production.yaml --project ${PRODUCTION_PROJECT_ID}

# Cron設定をデプロイ
deploy-cron:
	@gcloud app deploy -q appengine/config/cron.yaml --project ${STAGING_PROJECT_ID}

deploy-cron-production:
	@gcloud app deploy -q appengine/config/cron.yaml --project ${PRODUCTION_PROJECT_ID}

# Queue設定をデプロイ
deploy-queue:
	@gcloud app deploy -q appengine/config/queue.yaml --project ${STAGING_PROJECT_ID}

deploy-queue-production:
	@gcloud app deploy -q appengine/config/queue.yaml --project ${PRODUCTION_PROJECT_ID}

# Datastoreの複合インデックス定義をデプロイ
deploy-index:
	@gcloud app deploy -q appengine/config/index.yaml --project ${STAGING_PROJECT_ID}

deploy-index-production:
	@gcloud app deploy -q appengine/config/index.yaml --project ${PRODUCTION_PROJECT_ID}
