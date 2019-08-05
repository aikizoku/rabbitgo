.PHONY: init run deploy deploy-production

# 準備
init:
	${call init}

run:
	${call init}
	${call run,local,$(name)}

deploy:
	${call init}
	${call deploy,staging,$(name),${call get_env,credentials.staging.project_id}}

deploy-production:
	${call init}
	${call deploy,production,$(name),${call get_env,credentials.production.project_id}}

seed:
	${call seed}

# マクロ
define get_env
$(shell node -p "require('../env.json').$1")
endef

define init
@go run ../command/init_appengine/main.go 
endef

define run
@cd ./deploy/$1/$2; $(GOPATH)/bin/realize start --server
endef

define deploy
@cd ./deploy/$1/$2; gcloud app deploy -q app.yaml --project $3
endef

define seed
@go run ../command/seed/main.go 
endef
