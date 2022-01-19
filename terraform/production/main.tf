locals {
    env = yamldecode(file("../../env.yaml")).production
    deploy = "production"
    region = "asia-northeast1"
    time_zone = "Asia/Tokyo"
    default_domain = "hoge.com"
    api_domain = "api.hoge.com"
}

module "infrastructures" {
    source = "../modules"
    project_id = local.env.PROJECT_ID
    project_num = local.env.PROJECT_NUM
    deploy = local.deploy
    internal_auth_token = local.env.INTERNAL_AUTH_TOKEN
    region = local.region
    time_zone = local.time_zone
    default_domain = local.default_domain
    api_domain = local.api_domain
    content_bucket_name = local.env.CONTENT_BUCKET_NAME
    algolia_app_id = local.env.ALGOLIA_APP_ID
    algolia_api_key = local.env.ALGOLIA_API_KEY
}
