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

func TestAccComputeRegionDisk_regionDiskBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionDiskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_regionDiskBasicExample(context),
			},
			{
				ResourceName:      "google_compute_region_disk.regiondisk",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionDisk_regionDiskBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name = "my-region-disk-%{random_suffix}"
  snapshot = "${google_compute_snapshot.snapdisk.self_link}"
  type = "pd-ssd"
  region = "us-central1"
  physical_block_size_bytes = 4096

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_disk" "disk" {
  name = "my-disk-%{random_suffix}"
  image = "debian-cloud/debian-9"
  size = 50
  type = "pd-ssd"
  zone = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name = "my-snapshot-%{random_suffix}"
  source_disk = "${google_compute_disk.disk.name}"
  zone = "us-central1-a"
}
`, context)
}

func testAccCheckComputeRegionDiskDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_region_disk" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/disks/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeRegionDisk still exists at %s", url)
		}
	}

	return nil
}
