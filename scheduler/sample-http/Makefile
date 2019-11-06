.PHONY: h deploy deploy-prod

include env.mk

_NAME := $(notdir $(CURDIR))

h:
	@echo 'deploy      ステージング環境にデプロイする'
	@echo 'deploy-prod 本番環境にデプロイする'

deploy:
	${call deploy,${call get-project,staging},${_STG_SCHEDULE},${_STG_DESCRIPTION},${_STG_URI},${_STG_HTTP_METHOD},${_STG_HTTP_HEADERS},${_STG_HTTP_BODY},${_STG_TIME_ZONE},${_STG_TIMEOUT}}

deploy-prod:
	${call deploy,${call get-project,production},${_PRD_SCHEDULE},${_PRD_DESCRIPTION},${_PRD_URI},${_PRD_HTTP_METHOD},${_PRD_HTTP_HEADERS},${_PRD_HTTP_BODY},${_PRD_TIME_ZONE},${_PRD_TIMEOUT}}

# マクロ
define get-project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@echo gcloud scheduler jobs create app-engine $(_NAME) \
--project $1
--schedule "$2" \
--description "$3" \
--uri "$4" \
--http-method $5 \
--headers "$6" \
--message-body "$7" \
--time-zone "$8" \
--attempt-deadline "$9"
endef
