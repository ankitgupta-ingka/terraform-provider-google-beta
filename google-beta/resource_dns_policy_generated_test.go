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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDnsPolicy_dnsPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDnsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsPolicy_dnsPolicyBasicExample(context),
			},
		},
	})
}

func testAccDnsPolicy_dnsPolicyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_policy" "example-policy" {
  provider = "google-beta"

  name = "example-policy-%{random_suffix}"
  enable_inbound_forwarding = true

  enable_logging = true

  alternative_name_server_config {
    target_name_servers {
      ipv4_address = "172.16.1.10"
    }
    target_name_servers {
      ipv4_address = "172.16.1.20"
    }
  }

  networks {
    network_url =  "${google_compute_network.network-1.self_link}"
  }
  networks {
    network_url =  "${google_compute_network.network-2.self_link}"
  }
}

resource "google_compute_network" "network-1" {
  provider = "google-beta"

  name = "network-1-%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  provider = "google-beta"

  name = "network-2-%{random_suffix}"
  auto_create_subnetworks = false
}

provider "google-beta"{
  region = "us-central1"
  zone   = "us-central1-a"
}
`, context)
}

func testAccCheckDnsPolicyDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_dns_policy" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/dns/v1beta2/projects/{{project}}/policies/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("DnsPolicy still exists at %s", url)
		}
	}

	return nil
}
