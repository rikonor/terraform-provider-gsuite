package gsuite

import (
	"fmt"
	"reflect"
	"testing"

	admin "google.golang.org/api/admin/directory/v1"
)

func TestResourceUserCreate(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		usersService       UsersService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"name": []interface{}{
					map[string]interface{}{
						"given_name":  "given_name",
						"family_name": "family_name",
					},
				},
				"primary_email":              "primary_email",
				"password":                   "password",
				"change_password_next_login": true,
			},
			usersService: &StubUsersService{
				InsertFunc: func(u *admin.User) (*admin.User, error) {
					expectedUser := &admin.User{
						Name: &admin.UserName{
							GivenName:  "given_name",
							FamilyName: "family_name",
						},
						PrimaryEmail:              "primary_email",
						Password:                  "password",
						ChangePasswordAtNextLogin: true,
					}

					if !reflect.DeepEqual(u, expectedUser) {
						return nil, fmt.Errorf("wrong user %v, expected %v", u, expectedUser)
					}

					u.Id = "id"

					return u, nil
				},
				GetFunc: func(uID string) (*admin.User, error) {
					expectedUserID := "id"

					if uID != expectedUserID {
						return nil, fmt.Errorf("wrong user id: %s, expected %s", uID, expectedUserID)
					}

					return &admin.User{
						Id: "id",
						Name: &admin.UserName{
							GivenName:  "given_name",
							FamilyName: "family_name",
						},
						PrimaryEmail: "primary_email",
					}, nil
				},
			},
			outputDataResource: map[string]interface{}{
				"id": "id",
				"name": []interface{}{
					map[string]interface{}{
						"given_name":  "given_name",
						"family_name": "family_name",
					},
				},
				"primary_email": "primary_email",
			},
		},
	}

	for _, tc := range testCases {
		svcWrapper := &ServiceWrapper{
			UsersService: tc.usersService,
		}

		d := resourceUser().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceUserCreate(d, svcWrapper); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		for k, v := range tc.outputDataResource {
			if k == "id" {
				vv := d.Id()
				if vv != v {
					t.Fatalf("wrong id %q, expected %q", vv, v)
				}
				continue
			}

			vv, ok := d.GetOk(k)
			if !ok {
				t.Fatalf("expected value to be set: %s", k)
			}

			if !reflect.DeepEqual(vv, v) {
				t.Fatalf("wrong value %q, expected %q", vv, v)
			}
		}
	}
}

func TestResourceUserRead(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		usersService       UsersService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"id": "id",
			},
			usersService: &StubUsersService{
				GetFunc: func(uID string) (*admin.User, error) {
					expectedUserID := "id"

					if uID != expectedUserID {
						return nil, fmt.Errorf("wrong user id: %s, expected %s", uID, expectedUserID)
					}

					return &admin.User{
						Id: "id",
						Name: &admin.UserName{
							GivenName:  "given_name",
							FamilyName: "family_name",
						},
						PrimaryEmail: "primary_email",
					}, nil
				},
			},
			outputDataResource: map[string]interface{}{
				"id": "id",
				"name": []interface{}{
					map[string]interface{}{
						"given_name":  "given_name",
						"family_name": "family_name",
					},
				},
				"primary_email": "primary_email",
			},
		},
	}

	for _, tc := range testCases {
		svcWrapper := &ServiceWrapper{
			UsersService: tc.usersService,
		}

		d := resourceUser().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceUserRead(d, svcWrapper); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		for k, v := range tc.outputDataResource {
			if k == "id" {
				vv := d.Id()
				if vv != v {
					t.Fatalf("wrong id %q, expected %q", vv, v)
				}
				continue
			}

			vv, ok := d.GetOk(k)
			if !ok {
				t.Fatalf("expected value to be set: %s", k)
			}

			if !reflect.DeepEqual(vv, v) {
				t.Fatalf("wrong value %q, expected %q", vv, v)
			}
		}
	}
}

func TestResourceUserDelete(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		usersService       UsersService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"id": "id",
			},
			usersService: &StubUsersService{
				DeleteFunc: func(uID string) error {
					expectedUserID := "id"

					if uID != expectedUserID {
						return fmt.Errorf("wrong user id: %s, expected %s", uID, expectedUserID)
					}

					return nil
				},
			},
			outputDataResource: map[string]interface{}{
				"id": "",
			},
		},
	}

	for _, tc := range testCases {
		svcWrapper := &ServiceWrapper{
			UsersService: tc.usersService,
		}

		d := resourceUser().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceUserDelete(d, svcWrapper); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		for k, v := range tc.outputDataResource {
			if k == "id" {
				vv := d.Id()
				if vv != v {
					t.Fatalf("wrong id %q, expected %q", vv, v)
				}
				continue
			}

			vv, ok := d.GetOk(k)
			if !ok {
				t.Fatalf("expected value to be set: %s", k)
			}

			if !reflect.DeepEqual(vv, v) {
				t.Fatalf("wrong value %q, expected %q", vv, v)
			}
		}
	}
}
