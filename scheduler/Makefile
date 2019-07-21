.PHONY: deploy-http deploy-http-production deploy-appengine deploy-appengine-production deploy-pubsub deploy-pubsub-production

deploy-http:
	${call deploy-http,$(name),${call get_env,credentials.staging.project_id},${call get_env,scheduler.staging['$(name)'].schedule},${call get_env,scheduler.staging['$(name)'].description},${call get_env,scheduler.staging['$(name)'].url},${call get_env,scheduler.staging['$(name)'].method},${call get_env,scheduler.staging['$(name)'].headers},${call get_env,scheduler.staging['$(name)'].body}}

deploy-http-production:
	${call deploy-http,$(name),${call get_env,credentials.production.project_id},${call get_env,scheduler.production['$(name)'].schedule},${call get_env,scheduler.production['$(name)'].description},${call get_env,scheduler.production['$(name)'].url},${call get_env,scheduler.production['$(name)'].method},${call get_env,scheduler.production['$(name)'].headers},${call get_env,scheduler.production['$(name)'].body}}

deploy-appengine:
	${call deploy-appengine,$(name),${call get_env,credentials.staging.project_id},${call get_env,scheduler.staging['$(name)'].schedule},${call get_env,scheduler.staging['$(name)'].description},${call get_env,scheduler.staging['$(name)'].service},${call get_env,scheduler.staging['$(name)'].relative_url},${call get_env,scheduler.staging['$(name)'].method},${call get_env,scheduler.staging['$(name)'].headers},${call get_env,scheduler.staging['$(name)'].body}}

deploy-appengine-production:
	${call deploy-appengine,$(name),${call get_env,credentials.production.project_id},${call get_env,scheduler.production['$(name)'].schedule},${call get_env,scheduler.production['$(name)'].description},${call get_env,scheduler.production['$(name)'].service},${call get_env,scheduler.production['$(name)'].relative_url},${call get_env,scheduler.production['$(name)'].method},${call get_env,scheduler.production['$(name)'].headers},${call get_env,scheduler.production['$(name)'].body}}

deploy-pubsub:
	${call deploy-pubsub,$(name),${call get_env,credentials.staging.project_id},${call get_env,scheduler.staging['$(name)'].schedule},${call get_env,scheduler.staging['$(name)'].description},${call get_env,scheduler.staging['$(name)'].topic},${call get_env,scheduler.staging['$(name)'].body}}

deploy-pubsub-production:
	${call deploy-pubsub,$(name),${call get_env,credentials.production.project_id},${call get_env,scheduler.production['$(name)'].schedule},${call get_env,scheduler.production['$(name)'].description},${call get_env,scheduler.production['$(name)'].topic},${call get_env,scheduler.production['$(name)'].body}}

# マクロ
define get_env
$(shell node -p "require('../env.json').$1")
endef

define deploy-http
@gcloud scheduler jobs create http $1 \
--schedule "$3" \
--description "$4" \
--uri "$5" \
--http-method $6 \
--headers "$7" \
--message-body "$8" \
--time-zone "Asia/Tokyo" \
--project $2
endef

define deploy-appengine
@gcloud scheduler jobs create app-engine $1 \
--schedule "$3" \
--description "$4" \
--service "$5" \
--relative-url "$6" \
--http-method $7 \
--headers "$8" \
--message-body "$9" \
--time-zone "Asia/Tokyo" \
--project $2
endef

define deploy-pubsub
@gcloud scheduler jobs create pubsub $1 \
--schedule "$3" \
--description "$4" \
--topic "$5" \
--message-body "$6" \
--time-zone "Asia/Tokyo" \
--project $2
endef
