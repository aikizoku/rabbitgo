.PHONY: h deploy deploy-prod

h:
	@echo 'deploy      FirebaseStorageにルールをデプロイする: ステージング'
	@echo 'deploy-prod FirebaseStorageにルールをデプロイする: 本番'

deploy:
	${call deploy,${call get-project,staging}}

deploy-prod:
	${call deploy,${call get-project,production}}

define get-project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@firebase deploy --only storage:rules --project $1
endef
