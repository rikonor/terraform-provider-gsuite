provider "gsuite" {
  credentials = "${var.gsuite_credentials}"
  user_email  = "${var.user_email}"

  oauth_scopes = [
    "https://www.googleapis.com/auth/admin.directory.group",
    "https://www.googleapis.com/auth/admin.directory.user",
  ]
}

resource "gsuite_group" "group" {
  name        = "group"
  description = "Group"
  email       = "group@example.com"
}

resource "gsuite_user" "user" {
  name {
    given_name  = "User"
    family_name = "Manual"
  }

  primary_email              = "user@example.com"
  password                   = "3ej8de29XC"
  change_password_next_login = true
}

resource "gsuite_group_membership" "membership" {
  group  = "${gsuite_group.group.email}"
  member = "${gsuite_user.user.primary_email}"
}
