.PHONY: h deploy-cors deploy-cors-prod

include env.mk

h:
	@echo 'deploy-cors:      GCSにCORS設定をデプロイする: ステージング'
	@echo 'deploy-cors-prod: GCSにCORS設定をデプロイする: 本番'

deploy-cors:
	${call deploy-cors,$(_STG_BUCKET)}

deploy-cors-prod:
	${call deploy-cors,$(_PRD_BUCKET)}

define deploy-cors
@gsutil cors set cors.json gs://$1
endef
