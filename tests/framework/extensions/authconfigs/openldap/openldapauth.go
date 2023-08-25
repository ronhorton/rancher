package openldapauth

import (
	"github.com/rancher/rancher/tests/framework/clients/rancher"
	management "github.com/rancher/rancher/tests/framework/clients/rancher/generated/management/v3"
)

const openLdapConfigNameBase = "openLdapAuthConfig"

type OpenLDAPCredentialConfig struct {
	Servers                         []string `json:"servers"`
	ServiceAccountDistinguishedName string   `json:"serviceAccountDistinguishedName"`
	ServiceAccountPassword          string   `json:"serviceAccountPassword"`
	UserSearchBase                  string   `json:"userSearchBase"`
	Port                            int      `json:"port"`
	TLS                             bool     `json:"tls"`
	// Certificate would typically be loaded from a file or another source. Commenting out for this example.
	// Certificate                  string `json:"certificate"`
	TestUsername string `json:"testUsername"`
	TestPassword string `json:"testPassword"`
}

// CreateOpenLDAPAuthConfig is a helper function that creates
// an openLDAP auth config, enables it, and returns the AuthConfig response
func CreateOpenLDAPAuthConfig(rancherClient *rancher.Client) (*management.AuthConfig, error) {
	// Hardcoding the values for this config
	openLdapCredentialConfig := OpenLDAPCredentialConfig{
		Servers:                         []string{"openldapqa.qa.rancher.space"},
		ServiceAccountDistinguishedName: "cn=admin,dc=qa,dc=rancher,dc=space",
		ServiceAccountPassword:          "cattle@123",
		UserSearchBase:                  "dc=qa,dc=rancher,dc=space",
		Port:                            389,
		TLS:                             false,
		TestUsername:                    "testuser1",
		TestPassword:                    "Tacos86!",
	}

	authConfig := management.AuthConfig{
		Name:                     openLdapConfigNameBase,
		OpenLDAPCredentialConfig: &openLdapCredentialConfig,
		Enabled:                  true,
	}

	resp := &management.AuthConfig{}
	err := rancherClient.Management.APIBaseClient.Ops.DoCreate(management.AuthConfigType, authConfig, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
