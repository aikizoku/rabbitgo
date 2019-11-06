.PHONY: h restore restore-prod

include env.mk

_NAME := $(notdir $(CURDIR))

h:
	@echo 'restore      Firestoreを指定のバックアップで復元する: ステージング'
	@echo 'restore-prod Firestoreを指定のバックアップで復元する: 本番'

restore:
	${call restore,${call get-project,staging},$(_STG_BACKUP_BUCKET),$(_STG_BACKUP_FILE_PATH)}

restore-prod:
	${call restore-prod,${call get-project,production},$(_PRD_BACKUP_BUCKET),$(_PRD_BACKUP_FILE_PATH)}

define get-project
$(shell node -p "require('../../project.json').$1")
endef

define restore
@gcloud firestore import gs://$2/$3 \
--project $1
endef
