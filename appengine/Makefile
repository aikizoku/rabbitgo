.PHONY: init run deploy

# 準備
init:
	${call init}

run:
	${call init}
	${call run,local,$(app)}

deploy:
	${call init}
	${call deploy,staging,$(app),${call read_env,credentials.staging.project_id}}

# マクロ
define read_env
	$(shell node -p "require('../env.json').$1")
endef

define init
	@go run ../command/init/main.go 
endef

define run
	@cd ./deploy/$1/$2; go run enviroment.go dependency.go routing.go main.go
endef

define deploy
	@cd ./deploy/$1/$2; gcloud app deploy -q app.yaml --project $3
endef