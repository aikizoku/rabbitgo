.PHONY: init deploy deploy-production

init:
	${call init}

deploy:
	${call init}
	${call deploy,$(name),staging,${call get_env,credentials.staging.project_id},${call get_env,functions.staging['$(name)'].runtime},${call get_env,functions.staging['$(name)'].entry},${call get_env,functions.staging['$(name)'].memory},${call get_env,functions.staging['$(name)'].timeout},${call get_env,functions.staging['$(name)'].trigger}}

deploy-production:
	${call init}
	${call deploy,$(name),production,${call get_env,credentials.production.project_id},${call get_env,functions.production['$(name)'].runtime},${call get_env,functions.production['$(name)'].entry},${call get_env,functions.production['$(name)'].memory},${call get_env,functions.production['$(name)'].timeout},${call get_env,functions.production['$(name)'].trigger}}

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
--runtime $4 \
--region asia-northeast1 \
--entry-point $5 \
--env-vars-file ./env.yaml \
--project $3 \
--memory $6 \
--timeout $7 \
$8
endef
