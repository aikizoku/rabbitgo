_STG_RUNTIME := nodejs10
_STG_REGION  := asia-northeast1
_STG_ENTRY   := subscribeSlack
_STG_MEMORY  := 128MB
_STG_TIMEOUT := 60s
_STG_TRIGGER := --trigger-topic cloud-builds
_STG_ENV     := SLACK_WEBHOOK_URL=

_PRD_RUNTIME := nodejs10
_PRD_REGION  := asia-northeast1
_PRD_ENTRY   := subscribeSlack
_PRD_MEMORY  := 128MB
_PRD_TIMEOUT := 60s
_PRD_TRIGGER := --trigger-topic cloud-builds
_PRD_ENV     := SLACK_WEBHOOK_URL=
