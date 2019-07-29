.PHONY: deploy deploy-production

deploy-cors:
	${call deploy-cors,$(name)}

deploy-cors-production:
	${call deploy-cors,$(name)}

# マクロ
define get_env
	$(shell node -p "require('../env.json').$1")
endef

define deploy-cors
	@gsutil cors set cors.json gs://$1
endef
