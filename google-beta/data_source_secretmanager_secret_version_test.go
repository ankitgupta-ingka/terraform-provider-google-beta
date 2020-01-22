package google

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatasourceSecretmanagerSecretVersion_basic(t *testing.T) {
	t.Parallel()

	randomString := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckSecretmanagerSecretVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceSecretmanagerSecretVersion_basic(randomString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatasourceSecretmanagerSecretVersion("data.google_secretmanager_secret_version.basic", "1"),
				),
			},
		},
	})
}

func TestAccDatasourceSecretmanagerSecretVersion_latest(t *testing.T) {
	t.Parallel()

	randomString := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckSecretmanagerSecretVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceSecretmanagerSecretVersion_latest(randomString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatasourceSecretmanagerSecretVersion("data.google_secretmanager_secret_version.latest", "2"),
				),
			},
		},
	})
}

func testAccCheckDatasourceSecretmanagerSecretVersion(n, expected string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find Secret Version data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("data source ID not set.")
		}

		version, ok := rs.Primary.Attributes["version"]
		if !ok {
			return errors.New("can't find 'version' attribute")
		}

		if version != expected {
			return fmt.Errorf("expected %s, got %s, version not found", expected, version)
		}
		return nil
	}
}

func testAccDatasourceSecretmanagerSecretVersion_latest(randomString string) string {
	return fmt.Sprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta
  secret_id = "tf-test-secret-version-%s"
  labels = {
    label = "my-label"
  }
  replication {
    automatic = true
  }
}

resource "google_secretmanager_secret_version" "secret-version-basic-1" {
  provider = google-beta
  secret = google_secretmanager_secret.secret-basic.name
  secret_data = "my-tf-test-secret-first"
}

resource "google_secretmanager_secret_version" "secret-version-basic-2" {
  provider = google-beta
  secret = google_secretmanager_secret.secret-basic.name
  secret_data = "my-tf-test-secret-second"

  depends_on = [google_secretmanager_secret_version.secret-version-basic-1]
}

data "google_secretmanager_secret_version" "latest" {
	provider = google-beta
	secret = google_secretmanager_secret_version.secret-version-basic-2.secret
}
`, randomString)
}

func testAccDatasourceSecretmanagerSecretVersion_basic(randomString string) string {
	return fmt.Sprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta
  secret_id = "tf-test-secret-version-%s"
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
  secret_data = "my-tf-test-secret-%s"
}

data "google_secretmanager_secret_version" "basic" {
	provider = google-beta
	secret = google_secretmanager_secret_version.secret-version-basic.secret
	version = 1
}
`, randomString, randomString)
}
