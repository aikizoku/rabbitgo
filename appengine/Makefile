.PHONY: init run deploy deploy-production

# 準備
init:
	${call init}

run:
	${call init}
	${call run,local}

deploy:
	${call init}
	${call deploy,staging,${call get_env,credentials.staging.project_id}}

deploy-production:
	${call init}
	${call deploy,production,${call get_env,credentials.production.project_id}}

# マクロ
define get_env
$(shell node -p "require('../env.json').$1")
endef

define init
@go run ../command/init_appengine/main.go 
endef

define run
@cd ./deploy/$1; $(GOPATH)/bin/gin -i --all -t ../../ run enviroment.go dependency.go routing.go main.go
endef

define deploy
@cd ./deploy/$1; gcloud app deploy -q app.yaml --project $2
endef
