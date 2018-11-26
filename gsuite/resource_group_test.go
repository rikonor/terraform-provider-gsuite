package gsuite

import (
	"fmt"
	"reflect"
	"testing"

	admin "google.golang.org/api/admin/directory/v1"
)

func TestResourceGroupCreate(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		groupsService      GroupsService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"name":        "name",
				"email":       "email",
				"description": "description",
			},
			groupsService: &StubGroupsService{
				InsertFunc: func(grp *admin.Group) (*admin.Group, error) {
					expectedGroup := &admin.Group{
						Name:        "name",
						Email:       "email",
						Description: "description",
					}

					if !reflect.DeepEqual(grp, expectedGroup) {
						return nil, fmt.Errorf("wrong group %v, expected %v", grp, expectedGroup)
					}

					grp.Id = "id"

					return grp, nil
				},
				GetFunc: func(grpID string) (*admin.Group, error) {
					expectedGroupID := "id"

					if grpID != expectedGroupID {
						return nil, fmt.Errorf("wrong group id: %s, expected %s", grpID, expectedGroupID)
					}

					return &admin.Group{
						Id:          "id",
						Name:        "name",
						Email:       "email",
						Description: "description",
					}, nil
				},
			},
			outputDataResource: map[string]interface{}{
				"id":          "id",
				"name":        "name",
				"email":       "email",
				"description": "description",
			},
		},
	}

	for _, tc := range testCases {
		svcWrapper := &ServiceWrapper{
			GroupsService: tc.groupsService,
		}

		d := resourceGroup().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceGroupCreate(d, svcWrapper); err != nil {
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

func TestResourceGroupRead(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		groupsService      GroupsService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"id": "id",
			},
			groupsService: &StubGroupsService{
				GetFunc: func(grpID string) (*admin.Group, error) {
					expectedGroupID := "id"

					if grpID != expectedGroupID {
						return nil, fmt.Errorf("wrong group id: %s, expected %s", grpID, expectedGroupID)
					}

					return &admin.Group{
						Id:          "id",
						Name:        "name",
						Email:       "email",
						Description: "description",
					}, nil
				},
			},
			outputDataResource: map[string]interface{}{
				"id":          "id",
				"name":        "name",
				"email":       "email",
				"description": "description",
			},
		},
	}

	for _, tc := range testCases {
		svcWrapper := &ServiceWrapper{
			GroupsService: tc.groupsService,
		}

		d := resourceGroup().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceGroupRead(d, svcWrapper); err != nil {
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

func TestResourceGroupDelete(t *testing.T) {
	type testCase struct {
		inputDataResource  map[string]interface{}
		groupsService      GroupsService
		outputDataResource map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			inputDataResource: map[string]interface{}{
				"id": "id",
			},
			groupsService: &StubGroupsService{
				DeleteFunc: func(grpID string) error {
					expectedGroupID := "id"

					if grpID != expectedGroupID {
						return fmt.Errorf("wrong group id: %s, expected %s", grpID, expectedGroupID)
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
			GroupsService: tc.groupsService,
		}

		d := resourceGroup().Data(nil)
		for k, v := range tc.inputDataResource {
			if k == "id" {
				d.SetId(v.(string))
				continue
			}
			d.Set(k, v)
		}

		if err := resourceGroupDelete(d, svcWrapper); err != nil {
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
