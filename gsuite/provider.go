package gsuite

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/admin/directory/v1"
)

var defaultOAuthScopes = []string{
	admin.AdminDirectoryUserScope,
	admin.AdminDirectoryGroupScope,
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credentials": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"GOOGLE_CREDENTIALS",
					"GOOGLE_CLOUD_KEYFILE_JSON",
					"GCLOUD_KEYFILE_JSON",
				}, ""),
				ValidateFunc: validateCredentials,
			},
			"user_email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"oauth_scopes": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"gsuite_user":             resourceUser(),
			"gsuite_group":            resourceGroup(),
			"gsuite_group_membership": resourceGroupMembership(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func validateCredentials(v interface{}, _ string) ([]string, []error) {
	creds := v.(string)

	// Check for valid JSON
	var tmp interface{}

	if err := json.Unmarshal([]byte(creds), &tmp); err != nil {
		return nil, []error{fmt.Errorf("invalid json: %s", err)}
	}

	return nil, nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	credentials := d.Get("credentials").(string)
	userEmail := d.Get("user_email").(string)

	oauthScopesSet := d.Get("oauth_scopes").(*schema.Set)
	oauthScopes := convertSetToStrings(oauthScopesSet)
	if len(oauthScopes) == 0 {
		oauthScopes = defaultOAuthScopes
	}

	httpClient, err := createAuthenticatedHTTPClient([]byte(credentials), userEmail, oauthScopes)
	if err != nil {
		return nil, err
	}

	svcWrapper := newServiceWrapper(httpClient)

	return svcWrapper, nil
}

type ServiceWrapper struct {
	UsersService  UsersService
	GroupsService GroupsService
}

func newServiceWrapper(httpClient *http.Client) *ServiceWrapper {
	svcWrapper := &ServiceWrapper{}

	adminSvc, _ := admin.New(httpClient)

	svcWrapper.UsersService = WrapUsersService(adminSvc.Users)
	svcWrapper.GroupsService = WrapGroupsService(adminSvc.Groups)

	return svcWrapper
}

func createAuthenticatedHTTPClient(jsonCredentials []byte, userEmail string, oauthScopes []string) (*http.Client, error) {
	jwtConfig, err := google.JWTConfigFromJSON(jsonCredentials, oauthScopes...)
	if err != nil {
		return nil, err
	}

	jwtConfig.Subject = userEmail

	httpClient := jwtConfig.Client(context.Background())

	return httpClient, nil
}

func convertSetToStrings(st *schema.Set) []string {
	ss := []string{}
	for _, s := range st.List() {
		ss = append(ss, s.(string))
	}
	return ss
}
