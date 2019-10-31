.PHONY: h deploy deploy-prod

# 共通
_NAME := aaaa
_SCHEDULE := * * * * *
_DESCRIPTION := aaaaa
_HTTP_METHOD := aaa
_HTTP_HEADERS := aaa
_HTTP_BODY := aaa

# HTTP
_URI := aaaaa

# AppEngine
_SERVICE := aaa
_RELATIVE_URL := aaa

h:
	@echo 'deploy:      ステージング環境にデプロイする'
	@echo 'deploy-prod: 本番環境にデプロイする'

deploy:
	${call deploy-appengine,${call get_project,staging}}

deploy-prod:
	${call deploy-appengine,${call get_project,production}}

# マクロ
define get_project
$(shell node -p "require('../../project.json').$1")
endef

define deploy-http
gcloud scheduler jobs create http $(_NAME) \
--schedule "$(_SCHEDULE)" \
--description "$(_DESCRIPTION)" \
--uri "$(_URI)" \
--http-method $(_HTTP_METHOD) \
--headers "$(_HTTP_HEADERS)" \
--message-body "$(_HTTP_BODY)" \
--time-zone "Asia/Tokyo" \
--project $1
endef

define deploy-appengine
gcloud scheduler jobs create app-engine $(_NAME) \
--schedule "$(_SCHEDULE)" \
--description "$(_DESCRIPTION)" \
--service "$(_SERVICE)" \
--relative-url "$(_RELATIVE_URL)" \
--http-method $(_HTTP_METHOD) \
--headers "$(_HTTP_HEADERS)" \
--message-body "$(_HTTP_BODY)" \
--time-zone "Asia/Tokyo" \
--project $1
endef
