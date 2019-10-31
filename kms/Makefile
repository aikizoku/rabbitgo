.PHONY: deploy-topic deploy-topic-production deploy-sub deploy-sub-production

help:
	@echo ほげほげ
	@echo ふがふが


encript:
	${call deploy-topic,$(name),${call get_env,credentials.staging.project_id}}

decript:
	${call deploy-topic,$(name),${call get_env,credentials.production.project_id}}

# マクロ
define get_env
	$(shell node -p "require('../env.json').$1")
endef

define deploy-topic
	@gcloud pubsub topics create $1 \
	--project $2
endef

define deploy-sub
	@gcloud pubsub subscriptions create $1 --topic $2 \
	--project $3
endef
