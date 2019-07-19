.PHONY: create-topic create-topic-production

create-topic:
	${call create-topic,$(name),${call get_env,credentials.staging.project_id}}

create-topic-production:
	${call create-topic,$(name),${call get_env,credentials.production.project_id}}

# マクロ
define get_env
	$(shell node -p "require('../env.json').$1")
endef

define create-topic
	@gcloud pubsub topics create $1 \
	--project $2
endef
