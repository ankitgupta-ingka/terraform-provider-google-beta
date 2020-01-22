// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSecretmanagerSecretVersion_secretVersionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckSecretmanagerSecretVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretmanagerSecretVersion_secretVersionBasicExample(context),
			},
		},
	})
}

func testAccSecretmanagerSecretVersion_secretVersionBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta

  secret_id = "tf-test-secret-version-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    automatic = true
  }
}


resource "google_secretmanager_secret_version" "secret-version-basic" {
  provider = google-beta

  secret = google_secretmanager_secret.secret-basic.name

  secret_data = "my-tf-test-secret%{random_suffix}"
}
`, context)
}

func testAccCheckSecretmanagerSecretVersionDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_secretmanager_secret_version" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{SecretmanagerBasePath}}{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("SecretmanagerSecretVersion still exists at %s", url)
		}
	}

	return nil
}
