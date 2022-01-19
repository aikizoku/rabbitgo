resource "google_cloudbuild_trigger" "backend_api_deploy" {
  project = var.project_id
  name = "backend-api-deploy"
  substitutions = {
    _SERVICE = "api"
  }

  github {
    name  = "xxxx"
    owner = "xxxx"

    push {
      branch = "${var.deploy == "production" ? "master" : "develop"}"
    }
  }

  build {
    step {
      name = "gcr.io/cloud-builders/gcloud"
      args = ["app", "deploy", "-q", "appengine/api/app_${var.deploy}.yaml", "--version", var.deploy]
    }
  }
}

resource "google_project_iam_member" "appengine_account" {
  project = var.project_id
  role    = "roles/appengine.appAdmin"
  member  = "serviceAccount:${var.project_num}@cloudbuild.gserviceaccount.com"
}

resource "google_project_iam_member" "service_account" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${var.project_num}@cloudbuild.gserviceaccount.com"
}
