.PHONY: h deploy deploy-prod

include env.mk

_NAME := $(notdir $(CURDIR))

h:
	@echo 'deploy      Functionsをデプロイする: ステージング'
	@echo 'deploy-prod Functionsをデプロイする: 本番'

deploy:
	${call deploy,${call get-project,staging},$(_STG_RUNTIME),$(_STG_REGION),$(_STG_ENTRY),$(_STG_ENV),$(_STG_MEMORY),$(_STG_TIMEOUT),$(_STG_TRIGGER)}

deploy-prod:
	${call deploy,${call get-project,production},$(_PRD_RUNTIME),$(_PRD_REGION),$(_PRD_ENTRY),$(_PRD_ENV),$(_PRD_MEMORY),$(_PRD_TIMEOUT),$(_PRD_TRIGGER)}

define get-project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@gcloud functions deploy $(_NAME) \
--project $1 \
--source ./ \
--runtime $2 \
--region $3 \
--entry-point $4 \
--set-env-vars $5 \
--memory $6 \
--timeout $7 \
$8
endef
