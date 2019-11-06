.PHONY: h run deploy deploy-prod

# 準備
h:
	@echo 'run         アプリを実行する: ローカル'
	@echo 'deploy      アプリをデプロイする: ステージング'
	@echo 'deploy-prod アプリをデプロイする: 本番'

run:
	${call run}

deploy:
	${call deploy,staging,${call get-project,staging}}

deploy-prod:
	${call deploy,production,${call get-project,production}}

define get-project
$(shell node -p "require('../../project.json').$1")
endef

define run
$(GOPATH)/bin/realize start --server
endef

define deploy
@gcloud app deploy -q app_$1.yaml --project $2
endef
