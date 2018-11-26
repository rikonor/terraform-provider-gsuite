package gsuite

import (
	"fmt"
	"reflect"
	"testing"

	admin "google.golang.org/api/admin/directory/v1"
)

func TestResourceGroupMembershipCreate(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		membersService     MembersService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"group":  "group",
				"member": "member",
				"role":   "role",
			},
			membersService: &StubMembersService{
				InsertFunc: func(groupKey string, member *admin.Member) (*admin.Member, error) {
					expectedGroupKey := "group"

					if groupKey != expectedGroupKey {
						return nil, fmt.Errorf("wrong group key %s, expected %s", groupKey, expectedGroupKey)
					}

					expectedMember := &admin.Member{
						Email: "member",
						Role:  "role",
					}

					if !reflect.DeepEqual(member, expectedMember) {
						return nil, fmt.Errorf("wrong member %v, expected %v", member, expectedMember)
					}

					member.Id = "id"

					return member, nil
				},
				GetFunc: func(groupKey string, memberKey string) (*admin.Member, error) {
					expectedGroupKey := "group"

					if groupKey != expectedGroupKey {
						return nil, fmt.Errorf("wrong group key: %s, expected %s", groupKey, expectedGroupKey)
					}

					expectedMemberKey := "member"

					if memberKey != expectedMemberKey {
						return nil, fmt.Errorf("wrong member key: %s, expected %s", memberKey, expectedMemberKey)
					}

					return &admin.Member{
						Id:    "id",
						Email: "member",
						Role:  "role",
					}, nil
				},
			},
			outputDataResource: map[string]interface{}{
				"id":     "id",
				"group":  "group",
				"member": "member",
				"role":   "role",
			},
		},
	}

	for _, tc := range testCases {
		svcWrapper := &ServiceWrapper{
			MembersService: tc.membersService,
		}

		d := resourceGroupMembership().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceGroupMembershipCreate(d, svcWrapper); err != nil {
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

func TestResourceGroupMembershipRead(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		membersService     MembersService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"id":     "id",
				"group":  "group",
				"member": "member",
			},
			membersService: &StubMembersService{
				GetFunc: func(groupKey string, memberKey string) (*admin.Member, error) {
					expectedGroupKey := "group"

					if groupKey != expectedGroupKey {
						return nil, fmt.Errorf("wrong group key: %s, expected %s", groupKey, expectedGroupKey)
					}

					expectedMemberKey := "member"

					if memberKey != expectedMemberKey {
						return nil, fmt.Errorf("wrong member key: %s, expected %s", memberKey, expectedMemberKey)
					}

					return &admin.Member{
						Id:    "id",
						Email: "member",
						Role:  "role",
					}, nil
				},
			},
			outputDataResource: map[string]interface{}{
				"id":     "id",
				"group":  "group",
				"member": "member",
				"role":   "role",
			},
		},
	}

	for _, tc := range testCases {
		svcWrapper := &ServiceWrapper{
			MembersService: tc.membersService,
		}

		d := resourceGroupMembership().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceGroupMembershipRead(d, svcWrapper); err != nil {
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

func TestResourceGroupMembershipDelete(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		membersService     MembersService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"id":     "id",
				"group":  "group",
				"member": "member",
			},
			membersService: &StubMembersService{
				DeleteFunc: func(groupKey string, memberKey string) error {
					expectedGroupKey := "group"

					if groupKey != expectedGroupKey {
						return fmt.Errorf("wrong group key: %s, expected %s", groupKey, expectedGroupKey)
					}

					expectedMemberKey := "member"

					if memberKey != expectedMemberKey {
						return fmt.Errorf("wrong member key: %s, expected %s", memberKey, expectedMemberKey)
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
			MembersService: tc.membersService,
		}

		d := resourceGroupMembership().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceGroupMembershipDelete(d, svcWrapper); err != nil {
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
