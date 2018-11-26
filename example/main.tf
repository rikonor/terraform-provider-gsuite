provider "gsuite" {
  credentials = "${file("/root/.config/gcloud/knickknacklabs-root-svc-accnt.json")}"
  user_email  = "or@knickknacklabs.com"

  oauth_scopes = [
    "https://www.googleapis.com/auth/admin.directory.group",
    "https://www.googleapis.com/auth/admin.directory.user",
  ]
}

resource "gsuite_group" "group" {
  name        = "test"
  description = "Test Group 1"
  email       = "testgroup@knickknacklabs.com"
}

resource "gsuite_user" "user" {
  name {
    given_name  = "Test"
    family_name = "User"
  }

  primary_email              = "testuser@knickknacklabs.com"
  password                   = "weioneoi2f2"
  change_password_next_login = true
}

# resource "gsuite_group_membership" "membership" {
#   address = "beep"
# }

