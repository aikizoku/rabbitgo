terraform {
    required_providers {
        algolia = {
            source = "k-yomo/algolia"
            version = ">= 0.1.0, < 1.0.0"
        }
    }
}

provider "algolia" {
  app_id = var.algolia_app_id
  api_key = var.algolia_api_key
}

locals {
  algolia_ranking = [
    "typo",
    "geo",
    "words",
    "filters",
    "proximity",
    "attribute",
    "exact",
    "custom",
  ]
  algolia_attributes_to_retrieve = [
    "objectID",
  ]
  algolia_typo_tolerance = "false"
  algolia_index_languages = ["ja"]
  algolia_query_languages = ["ja"]
  algolia_pagination_limited_to = 1000
}