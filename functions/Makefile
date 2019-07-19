.PHONY: init deploy deploy-production

init:
	${call init}

deploy:
	${call init}
	${call deploy,$(name),staging,${call get_env,credentials.staging.project_id},${call get_env,functions.staging.$(name).memory},${call get_env,functions.staging.$(name).timeout},${call get_env,functions.staging.$(name).trigger}}

deploy-production:
	${call init}
	${call deploy,$(name),production,${PRODUCTION_PROJECT_ID},128MB,60s,--trigger-http}

# マクロ
define get_env
	$(shell node -p "require('../env.json').$1")
endef

define init
	@go run ../command/init_functions/main.go 
endef

define deploy
	@cd ./deploy/$2/$1; gcloud functions deploy $1 \
	--source ./ \
	--runtime go111 \
	--region asia-northeast1 \
	--entry-point Main \
	--env-vars-file ./env.yaml \
	--project $3 \
	--memory $4 \
	--timeout $5 \
	$6
endef
