package gsuite

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
