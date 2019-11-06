.PHONY: h restore restore-prod

include env.mk

_NAME := $(notdir $(CURDIR))

h:
	@echo 'restore      指定のバックアップファイルでステージング環境を復元をする'
	@echo 'restore-prod 指定のバックアップファイルで本番環境を復元をする'

restore:
	${call restore,${call get-project,staging},$(_STG_BUCKET),$(_STG_FILE_PATH)}

restore-prod:
	${call restore-prod,${call get-project,production},$(_PRD_BUCKET),$(_PRD_FILE_PATH)}

# マクロ
define get-project
$(shell node -p "require('../../project.json').$1")
endef

define restore
@gcloud firestore import gs://$2/$3 \
--project $1
endef
