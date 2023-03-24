package okta

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccOktaIdpOidc_crud
// Classic and OIE orgs.
// Org needs "Core", "Single Sign-On", "Universal Directory" SKUs in Workforce Identity
func TestAccOktaIdpOidc_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idpOidc)
	config := mgr.GetFixtures("generic_oidc.tf", ri, t)
	updatedConfig := mgr.GetFixtures("generic_oidc_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idpOidc)

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idpOidc, createDoesIdpExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceName(ri)),
					resource.TestCheckResourceAttr(resourceName, "authorization_url", "https://idp.example.com/authorize"),
					resource.TestCheckResourceAttr(resourceName, "authorization_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "token_url", "https://idp.example.com/token"),
					resource.TestCheckResourceAttr(resourceName, "token_binding", "HTTP-POST"),
					resource.TestCheckResourceAttr(resourceName, "user_info_url", "https://idp.example.com/userinfo"),
					resource.TestCheckResourceAttr(resourceName, "user_info_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "jwks_url", "https://idp.example.com/keys"),
					resource.TestCheckResourceAttr(resourceName, "jwks_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "efg456"),
					resource.TestCheckResourceAttr(resourceName, "client_secret", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
					resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://id.example.com"),
					resource.TestCheckResourceAttr(resourceName, "username_template", "idpuser.email"),
					resource.TestCheckResourceAttr(resourceName, "request_signature_algorithm", "HS256"),
					resource.TestCheckResourceAttr(resourceName, "request_signature_scope", "REQUEST"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceName(ri)),
					resource.TestCheckResourceAttr(resourceName, "authorization_url", "https://idp.example.com/authorize2"),
					resource.TestCheckResourceAttr(resourceName, "authorization_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "token_url", "https://idp.example.com/token2"),
					resource.TestCheckResourceAttr(resourceName, "token_binding", "HTTP-POST"),
					resource.TestCheckResourceAttr(resourceName, "user_info_url", "https://idp.example.com/userinfo2"),
					resource.TestCheckResourceAttr(resourceName, "user_info_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "jwks_url", "https://idp.example.com/keys2"),
					resource.TestCheckResourceAttr(resourceName, "jwks_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "efg456"),
					resource.TestCheckResourceAttr(resourceName, "client_secret", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
					resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://id.example.com"),
					resource.TestCheckResourceAttr(resourceName, "username_template", "idpuser.email"),
					resource.TestCheckResourceAttr(resourceName, "request_signature_algorithm", "HS256"),
					resource.TestCheckResourceAttr(resourceName, "request_signature_scope", "REQUEST"),
				),
			},
		},
	})
}

// TestAccOktaIdpOidc_algorithm
// Classic and OIE orgs.
// Org needs "Core", "Single Sign-On", "Universal Directory" SKUs in Workforce Identity
func TestAccOktaIdpOidc_algorithm(t *testing.T) {
	config := `
resource "okta_idp_oidc" "test" {
  name                  = "testAcc_replace_with_uuid"
  authorization_url     = "https://idp.example.com/authorize"
  authorization_binding = "HTTP-REDIRECT"
  token_url             = "https://idp.example.com/token"
  token_binding         = "HTTP-POST"
  user_info_url         = "https://idp.example.com/userinfo"
  user_info_binding     = "HTTP-REDIRECT"
  jwks_url              = "https://idp.example.com/keys"
  jwks_binding          = "HTTP-REDIRECT"
  scopes                = ["openid"]
  client_id             = "efg456"
  client_secret         = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  issuer_url            = "https://id.example.com"
  username_template     = "idpuser.email"
  request_signature_algorithm = "SHA-256"
  request_signature_scope = "REQUEST"
}`

	ri := acctest.RandInt()
	mgr := newFixtureManager(idpOidc)
	resourceName := fmt.Sprintf("%s.test", idpOidc)

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idpOidc, createDoesIdpExist()),
		Steps: []resource.TestStep{
			{
				Config: mgr.ConfigReplace(config, ri),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceName(ri)),
					resource.TestCheckResourceAttr(resourceName, "authorization_url", "https://idp.example.com/authorize"),
					resource.TestCheckResourceAttr(resourceName, "authorization_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "token_url", "https://idp.example.com/token"),
					resource.TestCheckResourceAttr(resourceName, "token_binding", "HTTP-POST"),
					resource.TestCheckResourceAttr(resourceName, "user_info_url", "https://idp.example.com/userinfo"),
					resource.TestCheckResourceAttr(resourceName, "user_info_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "jwks_url", "https://idp.example.com/keys"),
					resource.TestCheckResourceAttr(resourceName, "jwks_binding", "HTTP-REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "efg456"),
					resource.TestCheckResourceAttr(resourceName, "client_secret", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
					resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://id.example.com"),
					resource.TestCheckResourceAttr(resourceName, "username_template", "idpuser.email"),
					resource.TestCheckResourceAttr(resourceName, "request_signature_algorithm", "SHA-256"),
					resource.TestCheckResourceAttr(resourceName, "request_signature_scope", "REQUEST"),
				),
			},
		},
	})
}
