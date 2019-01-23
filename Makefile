GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello init run deploy

hello:
	@echo Hello go project ${GOPHER}

# 準備
init:
	@rm -rf deploy
	@mkdir -p deploy
	@mkdir -p deploy/appengine
	@mkdir -p deploy/appengine/local
	@mkdir -p deploy/appengine/staging
	@mkdir -p deploy/appengine/production
	${call apps}

# [GAE] アプリの実行
run:
	${call run,local,${app}}

# [GAE] アプリのデプロイ
deploy:
	${call deploy,staging,${app},${STAGING_PROJECT_ID}}

deploy-production:
	${call deploy,production,${app},${PRODUCTION_PROJECT_ID}}

# [GAE] ディスパッチ設定をデプロイ
deploy-dispatch:
	${call deploy-config,staging,dispatch.yaml,${STAGING_PROJECT_ID}}

deploy-dispatch-production:
	${call deploy-config,production,dispatch.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Cron設定をデプロイ
deploy-cron:
	${call deploy-config,staging,cron.yaml,${STAGING_PROJECT_ID}}

deploy-cron-production:
	${call deploy-config,production,cron.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Queue設定をデプロイ
deploy-queue:
	${call deploy-config,staging,queue.yaml,${STAGING_PROJECT_ID}}

deploy-queue-production:
	${call deploy-config,production,queue.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Datastoreの複合インデックス定義をデプロイ
deploy-index:
	${call deploy-config,staging,index.yaml,${STAGING_PROJECT_ID}}

deploy-index-production:
	${call deploy-config,production,index.yaml,${PRODUCTION_PROJECT_ID}}

# [Firestore] 全データ削除
firestore-delete:
	${call firestore-delete,${LOCAL_PROJECT_ID}}

firestore-delete-staging:
	${call firestore-delete,${STAGING_PROJECT_ID}}

# マクロ
define init
	@mkdir -p deploy/appengine/$1/$2
	@ln -s ../../../../appengine/app/$2/app_$1.yaml deploy/appengine/$1/$2/app.yaml
	@ln -s ../../../../appengine/app/$2/main.go deploy/appengine/$1/$2/main.go
	@ln -s ../../../../appengine/app/$2/dependency.go deploy/appengine/$1/$2/dependency.go
	@ln -s ../../../../appengine/app/$2/routing.go deploy/appengine/$1/$2/routing.go
	@ln -s ../../../../appengine/config/cron.yaml deploy/appengine/$1/$2/cron.yaml
	@ln -s ../../../../appengine/config/dispatch_$1.yaml deploy/appengine/$1/$2/dispatch.yaml
	@ln -s ../../../../appengine/config/index.yaml deploy/appengine/$1/$2/index.yaml
	@ln -s ../../../../appengine/config/queue.yaml deploy/appengine/$1/$2/queue.yaml
	@ln -s ../../../../appengine/env/values_$1.yaml deploy/appengine/$1/$2/values.yaml
	@ln -s ../../../../appengine/env/credentials_$1.json deploy/appengine/$1/$2/credentials.json
endef

define run
	dev_appserver.py deploy/appengine/$1/$2/app.yaml
endef

define deploy
	@gcloud app deploy -q deploy/appengine/$1/$2/app.yaml --project=$3
endef

define deploy-config
	@gcloud app deploy -q deploy/appengine/$1/api/$2 --project $3
endef

define firestore-delete
	firebase firestore:delete --all-collections --project $1
endef

include env.mk
