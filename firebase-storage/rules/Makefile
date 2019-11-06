.PHONY: deploy deploy-prod

deploy:
	${call deploy,${call get-project,staging}}

deploy-prod:
	${call deploy,${call get-project,production}}

# マクロ
define get-project
$(shell node -p "require('../../project.json').$1")
endef

define deploy
@firebase deploy --only storage:rules --project $1
endef
