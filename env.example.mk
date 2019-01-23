LOCAL_PROJECT_ID = 'develop-xxxxx-rabee-jp'
STAGING_PROJECT_ID = 'staging-skgo-rabee-jp'
PRODUCTION_PROJECT_ID = 'skgo-rabee-jp'

define apps
	$(call init,local,api)
	$(call init,staging,api)
	$(call init,production,api)

	$(call init,local,worker)
	$(call init,staging,worker)
	$(call init,production,worker)
endef
