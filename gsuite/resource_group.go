package gsuite

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	admin "google.golang.org/api/admin/directory/v1"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Delete: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	grp, err := svcWrapper.GroupsService.Insert(&admin.Group{
		Name:        d.Get("name").(string),
		Email:       d.Get("email").(string),
		Description: d.Get("description").(string),
	})

	if err != nil {
		return fmt.Errorf("failed to create group: %s", err)
	}

	d.SetId(grp.Id)

	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	grpID := d.Id()
	grp, err := svcWrapper.GroupsService.Get(grpID)
	if err != nil {
		return fmt.Errorf("failed to get group %s: %s", grpID, err)
	}

	d.SetId(grp.Id)
	d.Set("name", grp.Name)
	d.Set("email", grp.Email)
	d.Set("description", grp.Description)

	return nil
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	grpID := d.Id()
	if err := svcWrapper.GroupsService.Delete(grpID); err != nil {
		return fmt.Errorf("failed to delete group %s: %s", grpID, err)
	}

	d.SetId("")
	return nil
}
