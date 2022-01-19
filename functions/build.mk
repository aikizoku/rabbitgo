define get-project
$(shell yq e '.$1.PROJECT_ID' ../../env.yaml)
endef

define deploy-pubsub-staging
@gcloud functions deploy $(_NAME) \
--project ${call get-project,staging} \
--env-vars-file ./env_staging.yaml \
--source ./ \
--runtime $(_RUNTIME) \
--region $(_REGION) \
--entry-point $(_ENTRY) \
--memory $(_STG_MEMORY) \
--timeout $(_TIMEOUT) \
--trigger-topic $(_TOPIC)
endef

define deploy-pubsub-production
@gcloud functions deploy $(_NAME) \
--project ${call get-project,production} \
--env-vars-file ./env_production.yaml \
--source ./ \
--runtime $(_RUNTIME) \
--region $(_REGION) \
--entry-point $(_ENTRY) \
--memory $(_PRD_MEMORY) \
--timeout $(_TIMEOUT) \
--trigger-topic $(_TOPIC)
endef

define deploy-http-staging
@gcloud functions deploy $(_NAME) \
--project ${call get-project,staging} \
--env-vars-file ./env_staging.yaml \
--source ./ \
--runtime $(_RUNTIME) \
--region $(_REGION) \
--entry-point $(_ENTRY) \
--memory $(_STG_MEMORY) \
--timeout $(_TIMEOUT) \
--trigger-http
endef

define deploy-http-production
@gcloud functions deploy $(_NAME) \
--project ${call get-project,production} \
--env-vars-file ./env_production.yaml \
--source ./ \
--runtime $(_RUNTIME) \
--region $(_REGION) \
--entry-point $(_ENTRY) \
--memory $(_PRD_MEMORY) \
--timeout $(_TIMEOUT) \
--trigger-http
endef
