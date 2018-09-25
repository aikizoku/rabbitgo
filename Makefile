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

# デプロイ準備
deploy-init:
	@rm -rf deploy
	@mkdir deploy
	@mkdir deploy/staging
	@mkdir deploy/production

	# API
	@mkdir deploy/staging/api
	@ln -s ../../../appengine/app/api/app_staging.yaml deploy/staging/api/app.yaml
	@ln -s ../../../appengine/app/api/main.go deploy/staging/api/main.go
	@ln -s ../../../appengine/app/api/dependency.go deploy/staging/api/dependency.go
	@ln -s ../../../appengine/app/api/routing.go deploy/staging/api/routing.go
	@ln -s ../../../appengine/config/cron.yaml deploy/staging/api/cron.yaml
	@ln -s ../../../appengine/config/dispatch_staging.yaml deploy/staging/api/dispatch.yaml
	@ln -s ../../../appengine/config/index.yaml deploy/staging/api/index.yaml
	@ln -s ../../../appengine/config/queue.yaml deploy/staging/api/queue.yaml
	@ln -s ../../../appengine/secret/env_variables_staging.yaml deploy/staging/api/env_variables.yaml
	@ln -s ../../../appengine/secret/firebase_credentials_staging.yaml deploy/staging/api/firebase_credentials.yaml

	@mkdir deploy/production/api
	@ln -s ../../../appengine/app/api/app_production.yaml deploy/production/api/app.yaml
	@ln -s ../../../appengine/app/api/main.go deploy/production/api/main.go
	@ln -s ../../../appengine/app/api/dependency.go deploy/production/api/dependency.go
	@ln -s ../../../appengine/app/api/routing.go deploy/production/api/routing.go
	@ln -s ../../../appengine/config/cron.yaml deploy/production/api/cron.yaml
	@ln -s ../../../appengine/config/dispatch_production.yaml deploy/production/api/dispatch.yaml
	@ln -s ../../../appengine/config/index.yaml deploy/production/api/index.yaml
	@ln -s ../../../appengine/config/queue.yaml deploy/production/api/queue.yaml
	@ln -s ../../../appengine/secret/env_variables_production.yaml deploy/production/api/env_variables.yaml
	@ln -s ../../../appengine/secret/firebase_credentials_production.yaml deploy/production/api/firebase_credentials.yaml

	# Worker
	@mkdir deploy/staging/worker
	@ln -s ../../../appengine/app/worker/app_staging.yaml deploy/staging/worker/app.yaml
	@ln -s ../../../appengine/app/worker/main.go deploy/staging/worker/main.go
	@ln -s ../../../appengine/app/worker/dependency.go deploy/staging/worker/dependency.go
	@ln -s ../../../appengine/app/worker/routing.go deploy/staging/worker/routing.go
	@ln -s ../../../appengine/config/cron.yaml deploy/staging/worker/cron.yaml
	@ln -s ../../../appengine/config/dispatch_staging.yaml deploy/staging/worker/dispatch.yaml
	@ln -s ../../../appengine/config/index.yaml deploy/staging/worker/index.yaml
	@ln -s ../../../appengine/config/queue.yaml deploy/staging/worker/queue.yaml
	@ln -s ../../../appengine/secret/env_variables_staging.yaml deploy/staging/worker/env_variables.yaml
	@ln -s ../../../appengine/secret/firebase_credentials_staging.yaml deploy/staging/worker/firebase_credentials.yaml

	@mkdir deploy/production/worker
	@ln -s ../../../appengine/app/worker/app_production.yaml deploy/production/worker/app.yaml
	@ln -s ../../../appengine/app/worker/main.go deploy/production/worker/main.go
	@ln -s ../../../appengine/app/worker/dependency.go deploy/production/worker/dependency.go
	@ln -s ../../../appengine/app/worker/routing.go deploy/production/worker/routing.go
	@ln -s ../../../appengine/config/cron.yaml deploy/production/worker/cron.yaml
	@ln -s ../../../appengine/config/dispatch_production.yaml deploy/production/worker/dispatch.yaml
	@ln -s ../../../appengine/config/index.yaml deploy/production/worker/index.yaml
	@ln -s ../../../appengine/config/queue.yaml deploy/production/worker/queue.yaml
	@ln -s ../../../appengine/secret/env_variables_production.yaml deploy/production/worker/env_variables.yaml
	@ln -s ../../../appengine/secret/firebase_credentials_production.yaml deploy/production/worker/firebase_credentials.yaml

# アプリのデプロイ
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
