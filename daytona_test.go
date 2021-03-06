/*
   Copyright 2019 GM Cruise LLC

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
)

const testToken = "s.iyNUhq8Ov4hIAx6snw5mB2nL"
const testTokenLookupPayload = `
{
  "data": {
	"accessor": "8609694a-cdbc-db9b-d345-e782dbb562ed",
	"creation_time": 1523979354,
	"creation_ttl": 2764800,
	"display_name": "ldap2-hortensia",
	"entity_id": "7d2e3179-f69b-450c-7179-ac8ee8bd8ca9",
	"expire_time": "2018-05-19T11:35:54.466476215-04:00",
	"explicit_max_ttl": 0,
	"id": "cf64a70f-3a12-3f6c-791d-6cef6d390eed",
	"identity_policies": [
	  "dev-group-policy"
	],
	"issue_time": "2018-04-17T11:35:54.466476078-04:00",
	"meta": {
	  "username": "hortensia"
	},
	"num_uses": 0,
	"orphan": true,
	"path": "auth/ldap2/login/hortensia",
	"policies": [
	  "default",
	  "testgroup2-policy"
	],
	"renewable": true,
	"ttl": 2764790
  }
}
`

const testAuthResponse = `
{
	"auth": {
	  "client_token": "b.AAAAAQL_tyer_gNuQqvQYPVQgsNxjap_YW1NB2m4CDHHadQo7rF2XLFGdw-NJplAZNKbfloOvifrbpRCGdgG1taTqmC7D-a_qftN64zeL10SmNwEoDTiPzC_1aS1KExbtVftU3Sx16cBVqaynwsYRDfVnfTAffE",
	  "accessor": "0e9e354a-520f-df04-6867-ee81cae3d42d",
	  "policies": [
		"default",
		"dev",
		"prod"
	  ],
	  "metadata": {
		"project_id": "my-project",
		"role": "my-role",
		"service_account_email": "dev1@project-123456.iam.gserviceaccount.com",
		"service_account_id": "111111111111111111111"
	  },
	  "lease_duration": 2764800,
	  "renewable": true
	}
  }
`

func getTestClient(url string) (*api.Client, error) {
	vaultConfig := api.DefaultConfig()
	vaultConfig.Address = url
	vaultConfig.ConfigureTLS(&api.TLSConfig{Insecure: true})
	client, err := api.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func generateStandardTestConfig() {
	config.k8sAuth = true
	config.maximumAuthRetry = 1
}

func TestInvalidConfig(t *testing.T) {
	config.k8sAuth = true
	config.iamAuth = true
	config.gcpIAMAuth = false
	err := validateConfig()
	if err == nil {
		t.Fatal("expected an error from invalid config")
	}

	config.k8sAuth = true
	config.iamAuth = false
	config.gcpIAMAuth = false
	err = validateConfig()
	assert.Equal(t, "You must supply a role name via VAULT_AUTH_ROLE or -vault-auth-role", err.Error())
}
