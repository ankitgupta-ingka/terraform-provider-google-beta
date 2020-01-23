package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSecretmanagerSecret_import(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecretmanagerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretmanagerSecret_secretConfig_basic(context),
			},
			{
				ResourceName:      "google_secretmanager_secret.secret-basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecretmanagerSecret_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}
`, context)
}
