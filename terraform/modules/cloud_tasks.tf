resource "google_cloud_tasks_queue" "default" {
  project = var.project_id
  name = "default"
  location = var.region

  rate_limits {
    max_dispatches_per_second = 100
  }

  retry_config {
    max_attempts = 1
  }
}
