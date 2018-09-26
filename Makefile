GOPHER = 'ʕ◔ϖ◔ʔ'
STAGING_PROJECT_ID = 'beego-staging-thehero-jp'
PRODUCTION_PROJECT_ID = 'beego-thehero-jp'

hello:
	@echo Hello go project ${GOPHER}

# 準備
init:
	@rm -rf deploy
	@mkdir -p deploy
	@mkdir -p deploy/appengine
	@mkdir -p deploy/appengine/staging
	@mkdir -p deploy/appengine/production

	# API
	$(call init-appengine,staging,api)
	$(call init-appengine,production,api)

	# Worker
	$(call init-appengine,staging,worker)
	$(call init-appengine,production,worker)

# [GAE] アプリの実行
run-appengine-app:
	${call run-appengine-app,staging,${app}}

run-appengine-app-production:
	${call run-appengine-app,production,${app}}

# [GAE] アプリのデプロイ
deploy-appengine-app:
	${call deploy-appengine-app,staging,${app}}

deploy-appengine-app-production:
	${call deploy-appengine-app,production,${app}}

# [GAE] ディスパッチ設定をデプロイ
deploy-appengine-dispatch:
	${call deploy-appengine-config,dispatch_staging.yaml,${STAGING_PROJECT_ID}}

deploy-appengine-dispatch-production:
	${call deploy-appengine-config,dispatch_production.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Cron設定をデプロイ
deploy-appengine-cron:
	${call deploy-appengine-config,cron.yaml,${STAGING_PROJECT_ID}}

deploy-appengine-cron-production:
	${call deploy-appengine-config,cron.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Queue設定をデプロイ
deploy-appengine-queue:
	${call deploy-appengine-config,queue.yaml,${STAGING_PROJECT_ID}}

deploy-appengine-queue-production:
	${call deploy-appengine-config,queue.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Datastoreの複合インデックス定義をデプロイ
deploy-appengine-index:
	${call deploy-appengine-config,index.yaml,${STAGING_PROJECT_ID}}

deploy-appengine-index-production:
	${call deploy-appengine-config,index.yaml,${PRODUCTION_PROJECT_ID}}

# マクロ
define init-appengine
	@mkdir -p deploy/appengine/$1/$2
	@ln -s ../../../../appengine/app/$2/app_$1.yaml deploy/appengine/$1/$2/app.yaml
	@ln -s ../../../../appengine/app/$2/main.go deploy/appengine/$1/$2/main.go
	@ln -s ../../../../appengine/app/$2/dependency.go deploy/appengine/$1/$2/dependency.go
	@ln -s ../../../../appengine/app/$2/routing.go deploy/appengine/$1/$2/routing.go
	@ln -s ../../../../appengine/config/cron.yaml deploy/appengine/$1/$2/cron.yaml
	@ln -s ../../../../appengine/config/dispatch_$1.yaml deploy/appengine/$1/$2/dispatch.yaml
	@ln -s ../../../../appengine/config/index.yaml deploy/appengine/$1/$2/index.yaml
	@ln -s ../../../../appengine/config/queue.yaml deploy/appengine/$1/$2/queue.yaml
	@ln -s ../../../../appengine/secret/env_variables_$1.yaml deploy/appengine/$1/$2/env_variables.yaml
	@ln -s ../../../../appengine/secret/firebase_credentials_$1.json deploy/appengine/$1/$2/firebase_credentials.json
endef

define run-appengine-app
	dev_appserver.py deploy/appengine/$1/$2/app.yaml
endef

define deploy-appengine-app
	@gcloud app deploy -q deploy/appengine/$1/$2/app.yaml
endef

define deploy-appengine-config
	@gcloud app deploy -q appengine/config/$1 --project $2
endef
