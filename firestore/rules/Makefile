.PHONY: h deploy deploy-prod

_NAME := $(notdir $(CURDIR))

h:
	@echo 'deploy      Firestoreのルールをデプロイする: ステージング'
	@echo 'deploy-prod Firestoreのルールをデプロイする: 本番'

deploy:
	${call deploy,${call get-project,staging}}

deploy-prod:
	${call deploy,${call get-project,production}}

define get-project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@firebase deploy --only firestore:indexes --project $1
endef
