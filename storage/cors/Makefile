.PHONY: h deploy-cors deploy-prod-cors

_BUCKET_STAGING := aaaaaa
_BUCKET_PRODUCTION := aaaaaa

h:
	@echo 'deploy-cors:      CORS設定をステージング環境にデプロイする'
	@echo 'deploy-cors-prod: CORS設定を本番環境にデプロイする'

deploy:
	${call deploy,$(_BUCKET_STAGING)}

deploy-production:
	${call deploy,$(_BUCKET_PRODUCTION)}

# マクロ
define get_project
$(shell node -p "require('../project.json').$1")
endef

define deploy
	@gsutil cors set cors.json gs://$1
endef
