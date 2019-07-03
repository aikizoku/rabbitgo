GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello init run deploy

hello:
	@echo Hello go project ${GOPHER}

# 準備
init:
	${call init}

# [GAE] アプリの実行
run:
	${call init}
	${call run,local,${app}}

run-staging:
	${call init}
	${call run,staging,${app}}

run-production:
	${call init}
	${call run,production,${app}}

# [GAE] アプリのデプロイ
deploy:
	${call init}
	${call deploy,staging,${app}}

deploy-production:
	${call init}
	${call deploy,production,${app}}

# [GAE] Cron設定をデプロイ
deploy-cron:
	${call init}
	${call deploy-config,staging,cron}

deploy-cron-production:
	${call init}
	${call deploy-config,production,cron}

# [GAE] Queue設定をデプロイ
deploy-queue:
	${call init}
	${call deploy-config,staging,queue}

deploy-queue-production:
	${call init}
	${call deploy-config,production,queue}

# マクロ
define init
	@GO111MODULE=off go run ./command/init/main.go 
endef

define run
	@GO111MODULE=off go run ./command/run/main.go -env $1 -app $2
endef

define deploy
	@GO111MODULE=off go run ./command/deploy/main.go -env $1 -app $2
endef

define deploy-config
	@GO111MODULE=off go run ./command/deploy_config/main.go -env $1 -cfg $2
endef
