.PHONY: h deploy-cors deploy-cors-prod

include env.mk

h:
	@echo 'deploy-cors:      CORS設定をステージング環境にデプロイする'
	@echo 'deploy-cors-prod: CORS設定を本番環境にデプロイする'

deploy-cors:
	${call deploy-cors,$(_STG_BUCKET)}

deploy-cors-prod:
	${call deploy-cors,$(_PRD_BUCKET)}

# マクロ
define get-project
$(shell node -p "require('../project.json').$1")
endef

define deploy-cors
@gsutil cors set cors.json gs://$1
endef
