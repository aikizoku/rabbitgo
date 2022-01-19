.PHONY: h get get-prod deploy deploy-prod

_NAME := $(notdir $(CURDIR))

h:
	@echo 'get         Firestoreのインデックスを取得する: ステージング'
	@echo 'get-prod    Firestoreのインデックスを取得する: 本番'
	@echo 'deploy      Firestoreのインデックスをデプロイする: ステージング'
	@echo 'deploy-prod Firestoreのインデックスをデプロイする: 本番'

get:
	${call get,${call get-project,staging}}

get-prod:
	${call get,${call get-project,production}}

deploy:
	${call deploy,${call get-project,staging}}

deploy-prod:
	${call deploy,${call get-project,production}}

define get-project
$(shell yq e '.$1.PROJECT_ID' ../../env.yaml)
endef

define get
@firebase firestore:indexes --project $1 > firestore.indexes.json
endef

define deploy
@firebase deploy --only firestore:indexes --project $1
endef
