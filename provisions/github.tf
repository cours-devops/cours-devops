provider "github" {
  token        = "${var.github_token}"
  organization = "${var.github_organization}"
}

resource "github_repository" "example" {
  name        = "cours-devops"  
  private = false
}


variable "github_token" {
  type = "string"
}

variable "github_organization" {
  type = "string"
}

