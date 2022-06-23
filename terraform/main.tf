terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.3.0"
    }
  }
}


provider "google" {
  project = "vibrant-reach-354209"
  region  = "europe-west2"
  zone    = "europe-west2-a"
}

resource "google_container_cluster" "primary" {
  name               = "vibrant-reach-354209-gke"
  location           = "europe-west2"
  enable_autopilot   = true
  initial_node_count = 1
}