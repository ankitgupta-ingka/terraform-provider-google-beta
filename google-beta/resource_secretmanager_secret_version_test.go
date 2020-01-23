package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSecretmanagerSecretVersion_import(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecretmanagerSecretVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretmanagerSecretVersion_secretVersion_basic(context),
			},
			{
				ResourceName:      "google_secretmanager_secret_version.secret-version-basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecretmanagerSecretVersion_secretVersion_update(context),
			},
			{
				ResourceName:      "google_secretmanager_secret_version.secret-version-basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecretmanagerSecretVersion_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  secret_id = "tf-test-secret-version-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    automatic = true
  }
}


resource "google_secretmanager_secret_version" "secret-version-basic" {
  secret = google_secretmanager_secret.secret-basic.name

  secret_data = "my-tf-test-secret%{random_suffix}"
}
`, context)
}

func testAccSecretmanagerSecretVersion_secretVersion_update(context map[string]interface{}) string {
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

  secret_data = "my-tf-test-secret-2-%{random_suffix}"
}
`, context)
}
