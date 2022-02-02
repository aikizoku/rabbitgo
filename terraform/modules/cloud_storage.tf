resource "google_storage_bucket" "content" {
  project = var.project_id
  name = var.content_bucket_name
  location = var.region
}

resource "google_storage_bucket_iam_binding" "content" {
  bucket = var.content_bucket_name
    role = "roles/storage.legacyObjectReader"
    members = [
      "allUsers",
    ]
}
