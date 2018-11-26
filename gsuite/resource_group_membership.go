package gsuite

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	admin "google.golang.org/api/admin/directory/v1"
)

func resourceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupMembershipCreate,
		Read:   resourceGroupMembershipRead,
		Delete: resourceGroupMembershipDelete,

		Schema: map[string]*schema.Schema{
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"member": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"role": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "MEMBER",
				ValidateFunc: validation.StringInSlice([]string{"MEMBER", "MANAGER", "OWNER"}, false),
				ForceNew:     true,
			},
		},
	}
}

func resourceGroupMembershipCreate(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	mmbr, err := svcWrapper.MembersService.Insert(d.Get("group").(string), &admin.Member{
		Email: d.Get("member").(string),
		Role:  d.Get("role").(string),
	})

	if err != nil {
		return fmt.Errorf("failed to create membership: %s", err)
	}

	d.SetId(mmbr.Id)

	return resourceGroupMembershipRead(d, m)
}

func resourceGroupMembershipRead(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	mmbrGroup := d.Get("group").(string)
	mmbrEmail := d.Get("member").(string)

	mmbr, err := svcWrapper.MembersService.Get(mmbrGroup, mmbrEmail)
	if err != nil {
		return fmt.Errorf("failed to get member %s for group %s: %s", mmbrEmail, mmbrGroup, err)
	}

	d.SetId(mmbr.Id)
	d.Set("group", mmbrGroup)
	d.Set("member", mmbr.Email)
	d.Set("role", mmbr.Role)

	return nil
}

func resourceGroupMembershipDelete(d *schema.ResourceData, m interface{}) error {
	svcWrapper := m.(*ServiceWrapper)

	mmbrGroup := d.Get("group").(string)
	mmbrEmail := d.Get("member").(string)

	if err := svcWrapper.MembersService.Delete(mmbrGroup, mmbrEmail); err != nil {
		return fmt.Errorf("failed to remove member %s from group %s: %s", mmbrEmail, mmbrGroup, err)
	}

	d.SetId("")
	return nil
}
