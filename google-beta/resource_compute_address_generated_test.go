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
)

func TestAccComputeAddress_addressBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressBasicExample(context),
			},
			{
				ResourceName:      "google_compute_address.ip_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeAddress_addressBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "ip_address" {
  name = "my-address-%{random_suffix}"
}
`, context)
}

func TestAccComputeAddress_addressWithSubnetworkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressWithSubnetworkExample(context),
			},
			{
				ResourceName:      "google_compute_address.internal_with_subnet_and_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeAddress_addressWithSubnetworkExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  name = "my-network-%{random_suffix}"
}

resource "google_compute_subnetwork" "default" {
  name          = "my-subnet-%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = "${google_compute_network.default.self_link}"
}

resource "google_compute_address" "internal_with_subnet_and_address" {
  name         = "my-internal-address-%{random_suffix}"
  subnetwork   = "${google_compute_subnetwork.default.self_link}"
  address_type = "INTERNAL"
  address      = "10.0.42.42"
  region       = "us-central1"
}
`, context)
}

func TestAccComputeAddress_instanceWithIpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_instanceWithIpExample(context),
			},
			{
				ResourceName:      "google_compute_address.static",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeAddress_instanceWithIpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "static" {
  name = "ipv4-address-%{random_suffix}"
}

data "google_compute_image" "debian_image" {
	family  = "debian-9"
	project = "debian-cloud"
}

resource "google_compute_instance" "instance_with_ip" {
	name         = "vm-instance-%{random_suffix}"
	machine_type = "f1-micro"
	zone         = "us-central1-a"

	boot_disk {
		initialize_params{
			image = "${data.google_compute_image.debian_image.self_link}"
		}
	}

	network_interface {
		network = "default"
		access_config {
			nat_ip = "${google_compute_address.static.address}"
		}
	}
}
`, context)
}

func testAccCheckComputeAddressDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_address" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/addresses/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeAddress still exists at %s", url)
		}
	}

	return nil
}
