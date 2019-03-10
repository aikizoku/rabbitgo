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

# [GAE] ディスパッチ設定をデプロイ
deploy-dispatch:
	${call init}
	${call deploy-config,staging,dispatch}

deploy-dispatch-production:
	${call init}
	${call deploy-config,production,dispatch}

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

# [GAE] Datastoreの複合インデックス定義をデプロイ
deploy-index:
	${call init}
	${call deploy-config,staging,index}

deploy-index-production:
	${call init}
	${call deploy-config,production,index}

# [Firestore] 全データ削除
firestore-delete:
	${call init}
	${call firestore-delete,local}

firestore-delete-staging:
	${call init}
	${call firestore-delete,staging}

# マクロ
define init
	@go run ./command/init/main.go 
endef

define run
	@go run ./command/run/main.go -env $1 -app $2
endef

define deploy
	@go run ./command/deploy/main.go -env $1 -app $2
endef

define deploy-config
	@go run ./command/deploy_config/main.go -env $1 -cfg $2
endef

define firestore-delete
	@go run ./command/firestore_delete/main.go -env $1
endef
