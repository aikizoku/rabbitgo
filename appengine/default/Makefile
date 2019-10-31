.PHONY: run deploy deploy-prod

# 準備
help:
	@echo 'run         : ローカル環境で実行する'
	@echo 'deploy      : ステージング環境にデプロイする'
	@echo 'deploy-prod : 本番環境にデプロイする'

run:
	${call run}

deploy:
	${call deploy,staging,${call get_project,staging}}

deploy-prod:
	${call deploy,production,${call get_project,production}}

# マクロ
define get_project
$(shell node -p "require('../../project.json').$1")
endef

define run
$(GOPATH)/bin/realize start --server
endef

define deploy
@gcloud app deploy -q app_$1.yaml --project $2
endef
