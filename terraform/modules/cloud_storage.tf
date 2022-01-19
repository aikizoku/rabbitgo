resource "google_storage_bucket" "content" {
  project = var.project_id
  name = var.content_bucket_name
  location = var.region
}

resource "google_storage_bucket_access_control" "content" {
  bucket = var.content_bucket_name
  role   = "READER"
  entity = "allUsers"
}
