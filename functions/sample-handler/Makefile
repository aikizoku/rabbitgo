.PHONY: h deploy deploy-prod

_RUNTIME := go111
_REGION := asia-northeast1
_ENTRY := Handle
_MEMORY := 128MB
_TIMEOUT := 60s
_TRIGGER := --trigger-http

h:
	@echo 'deploy      : ステージング環境にデプロイする'
	@echo 'deploy-prod : 本番環境にデプロイする'

deploy:
	${call deploy,$(notdir $(CURDIR)),staging,${call get_project,staging}}

deploy-prod:
	${call deploy,$(notdir $(CURDIR)),production,${call get_project,production}}

# マクロ
define get_project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@gcloud functions deploy $1 \
--source ./ \
--runtime $(_RUNTIME) \
--region $(_REGION) \
--entry-point $(_ENTRY) \
--env-vars-file ./env_$2.yaml \
--project $3 \
--memory $(_MEMORY) \
--timeout $(_TIMEOUT) \
$(_TRIGGER)
endef
