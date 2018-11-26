provider "gsuite" {
  credentials = "${file("/root/.config/gcloud/knickknacklabs-root-svc-accnt.json")}"
  user_email  = "or@knickknacklabs.com"

  oauth_scopes = [
    "https://www.googleapis.com/auth/admin.directory.group",
  ]
}

# resource "gsuite_group" "group" {
#   name        = "test"
#   description = "Test Group 1"
#   email       = "testgroup@knickknacklabs.com"
# }


# resource "gsuite_user" "user" {
#   address = "beep"
# }


# resource "gsuite_group_membership" "membership" {
#   address = "beep"
# }

