package gsuite

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	admin "google.golang.org/api/admin/directory/v1"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"given_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"family_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				MinItems: 1,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
			},

			"primary_email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"change_password_next_login": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	u, err := svcWrapper.UsersService.Insert(&admin.User{
		Name: &admin.UserName{
			GivenName:  d.Get("name.0.given_name").(string),
			FamilyName: d.Get("name.0.family_name").(string),
		},
		PrimaryEmail:              d.Get("primary_email").(string),
		Password:                  d.Get("password").(string),
		ChangePasswordAtNextLogin: d.Get("change_password_next_login").(bool),
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}

	d.SetId(u.Id)

	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	uID := d.Id()
	u, err := svcWrapper.UsersService.Get(uID)
	if err != nil {
		return fmt.Errorf("failed to get user %s: %s", uID, err)
	}

	d.SetId(u.Id)

	d.Set("name", []interface{}{
		map[string]interface{}{
			"given_name":  u.Name.GivenName,
			"family_name": u.Name.FamilyName,
		},
	})

	d.Set("primary_email", u.PrimaryEmail)

	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	uID := d.Id()
	if err := svcWrapper.UsersService.Delete(uID); err != nil {
		return fmt.Errorf("failed to delete user %s: %s", uID, err)
	}

	d.SetId("")
	return nil
}
