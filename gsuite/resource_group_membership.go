package gsuite

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupMembershipCreate,
		Read:   resourceGroupMembershipRead,
		Update: resourceGroupMembershipUpdate,
		Delete: resourceGroupMembershipDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGroupMembershipCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceGroupMembershipRead(d, m)
}

func resourceGroupMembershipRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGroupMembershipUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGroupMembershipRead(d, m)
}

func resourceGroupMembershipDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
