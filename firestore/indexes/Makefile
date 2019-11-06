.PHONY: h get get-prod deploy deploy-prod

_NAME := $(notdir $(CURDIR))

h:
	@echo 'get         インデックスをステージング環境から取得する'
	@echo 'get-prod    インデックスを本番環境から取得する'
	@echo 'deploy      インデックスをステージング環境にデプロイする'
	@echo 'deploy-prod インデックスを本番環境にデプロイする'

get:
	${call get,${call get-project,staging}}

get-prod:
	${call get,${call get-project,production}}

deploy:
	${call deploy,${call get-project,staging}}

deploy-prod:
	${call deploy,${call get-project,production}}

# マクロ
define get-project
$(shell node -p "require('../../project.json').$1")
endef

define get
@firebase firestore:indexes --project $1 > firestore.indexes.json
endef

define deploy
@firebase deploy --only firestore:indexes --project $1
endef
