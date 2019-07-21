.PHONY: deploy deploy-production

deploy:
	${call deploy,${call get_env,credentials.staging.project_id}}

deploy-production:
	${call deploy,${call get_env,credentials.production.project_id}}

# マクロ
define get_env
$(shell node -p "require('../env.json').$1")
endef

define deploy
@gcloud app deploy -q queue.yaml \
--project $1
endef
