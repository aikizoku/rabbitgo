.PHONY: h deploy deploy-prod

h:
	@echo 'deploy      TasksのQueueをデプロイする: ステージング'
	@echo 'deploy-prod TasksのQueueをデプロイする: 本番'

deploy:
	${call deploy,${call get-project,staging}}

deploy-prod:
	${call deploy,${call get-project,production}}

define get-project
$(shell node -p "require('../project.json').$1")
endef

define deploy
@gcloud app deploy -q queue.yaml \
--project $1
endef
