.PHONY: deploy deploy-production

deploy:
	${call deploy,$(name),${call get_env,credentials.staging.project_id}}

deploy-production:
	${call deploy,$(name),${call get_env,credentials.production.project_id}}

# マクロ
define get_env
	$(shell node -p "require('../env.json').$1")
endef

define deploy
	@gcloud pubsub topics create $1 \
	--project $2
endef
