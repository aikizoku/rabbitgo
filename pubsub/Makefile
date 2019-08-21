.PHONY: deploy deploy-production

deploy-topic:
	${call deploy-topic,$(name),${call get_env,credentials.staging.project_id}}

deploy-topic-production:
	${call deploy-topic,$(name),${call get_env,credentials.production.project_id}}

deploy-sub:
	${call deploy-sub,$(name),${topic},${call get_env,credentials.staging.project_id}}

deploy-sub-production:
	${call deploy-sub,$(name),${topic},${call get_env,credentials.production.project_id}}

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
