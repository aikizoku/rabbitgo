.PHONY: deploy deploy-prod

h:
	@echo 'deploy:      ステージング環境にデプロイする'
	@echo 'deploy-prod: 本番環境にデプロイする'

deploy:
	${call deploy,${call get_project,staging}}

deploy-prod:
	${call deploy,${call get_project,production}}

# マクロ
define get_project
$(shell node -p "require('../project.json').$1")
endef

define deploy
@gcloud app deploy -q queue.yaml \
--project $1
endef
