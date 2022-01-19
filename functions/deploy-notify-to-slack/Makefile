.PHONY: deploy deploy-prod

include ../build.mk
include env.mk

_NAME := $(notdir $(CURDIR))

deploy:
	${call deploy-pubsub-staging}

deploy-prod:
	${call deploy-pubsub-production}
