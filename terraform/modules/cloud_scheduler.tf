resource "google_cloud_scheduler_job" "xxxx" {
  project = var.project_id
  name = "xxxx"
  schedule = "* * * * *"
  description = "XXXX"
  region = var.region
  time_zone = var.time_zone

  app_engine_http_target {
    http_method = "GET"

    app_engine_routing {
      service  = "api"
    }

    relative_uri = "/worker/xxxx/xxxx"

    headers = {
        Authorization = var.internal_auth_token
    }
  }
}
