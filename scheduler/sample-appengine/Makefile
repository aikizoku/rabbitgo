.PHONY: h deploy deploy-prod

include env.mk

_NAME := $(notdir $(CURDIR))

h:
	@echo 'deploy      Schedulerをデプロイする: ステージング'
	@echo 'deploy-prod Schedulerをデプロイする: 本番'

deploy:
	${call deploy,${call get-project,staging},${_STG_SCHEDULE},${_STG_DESCRIPTION},${_STG_SERVICE},${_STG_RELATIVE_URL},${_STG_HTTP_METHOD},${_STG_HTTP_HEADERS},${_STG_HTTP_BODY},${_STG_TIME_ZONE},${_STG_TIMEOUT}}

deploy-prod:
	${call deploy,${call get-project,production},${_PRD_SCHEDULE},${_PRD_DESCRIPTION},${_PRD_SERVICE},${_PRD_RELATIVE_URL},${_PRD_HTTP_METHOD},${_PRD_HTTP_HEADERS},${_PRD_HTTP_BODY},${_PRD_TIME_ZONE},${_PRD_TIMEOUT}}

define get-project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@gcloud scheduler jobs create app-engine $(_NAME) \
--project $1 \
--schedule "$2" \
--description "$3" \
--service "$4" \
--relative-url "$5" \
--http-method $6 \
--headers "$7" \
--message-body "$8" \
--time-zone "$9" \
--attempt-deadline "${10}"
endef
