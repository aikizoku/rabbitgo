resource "google_project_iam_member" "cloudbuild_serviceaccount" {
  count = "${length(var.cloudbuild_serviceaccount_roles)}"
  role = "${element(var.cloudbuild_serviceaccount_roles, count.index)}"
  member = "serviceAccount:${var.project_num}@cloudbuild.gserviceaccount.com"
  project = var.project_id
}

variable "cloudbuild_serviceaccount_roles" {
  default = [
    "roles/appengine.appAdmin",
    "roles/iam.serviceAccountUser"
  ]
}

resource "google_project_iam_member" "appengine_serviceaccount" {
  count = "${length(var.appengine_serviceaccount_roles)}"
  role = "${element(var.appengine_serviceaccount_roles, count.index)}"
  member = "serviceAccount:${var.project_id}@appspot.gserviceaccount.com"
  project = var.project_id
}

variable "appengine_serviceaccount_roles" {
  default = [
    "roles/iam.serviceAccountTokenCreator",
    "roles/storage.admin"
  ]
} 
