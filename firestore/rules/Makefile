.PHONY: h deploy deploy-prod

_NAME := $(notdir $(CURDIR))

h:
	@echo 'deploy      ルールをステージング環境にデプロイする'
	@echo 'deploy-prod ルールを本番環境にデプロイする'

deploy:
	${call deploy,${call get-project,staging}}

deploy-prod:
	${call deploy,${call get-project,production}}


# マクロ
define get-project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@firebase deploy --only firestore:indexes --project $1
endef
