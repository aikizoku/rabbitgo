resource "google_app_engine_application_url_dispatch_rules" "default" {
  project = var.project_id
  
  dispatch_rules {
    domain  = var.default_domain
    path = "/*"
    service = "default"
  }

  dispatch_rules {
    domain  = var.api_domain
    path = "/*"
    service = "api"
  }
}
