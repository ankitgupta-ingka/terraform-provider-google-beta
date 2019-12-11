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
)

func TestAccComputeAutoscaler_autoscalerSingleInstanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeAutoscalerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAutoscaler_autoscalerSingleInstanceExample(context),
			},
		},
	})
}

func testAccComputeAutoscaler_autoscalerSingleInstanceExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_autoscaler" "default" {
  provider = google-beta

  name   = "my-autoscaler%{random_suffix}"
  zone   = "us-central1-f"
  target = google_compute_instance_group_manager.default.self_link

  autoscaling_policy {
    max_replicas    = 5
    min_replicas    = 1
    cooldown_period = 60

    metric {
      name                       = "pubsub.googleapis.com/subscription/num_undelivered_messages"
      filter                     = "resource.type = pubsub_subscription AND resource.label.subscription_id = our-subscription"
      single_instance_assignment = 65535
    }
  }
}

resource "google_compute_instance_template" "default" {
  provider = google-beta

  name           = "my-instance-template%{random_suffix}"
  machine_type   = "n1-standard-1"
  can_ip_forward = false

  tags = ["foo", "bar"]

  disk {
    source_image = data.google_compute_image.debian_9.self_link
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_target_pool" "default" {
  provider = google-beta

  name = "my-target-pool%{random_suffix}"
}

resource "google_compute_instance_group_manager" "default" {
  provider = google-beta

  name = "my-igm%{random_suffix}"
  zone = "us-central1-f"

  version {
    instance_template = google_compute_instance_template.default.self_link
    name              = "primary"
  }

  target_pools       = [google_compute_target_pool.default.self_link]
  base_instance_name = "autoscaler-sample"
}

data "google_compute_image" "debian_9" {
  provider = google-beta

  family  = "debian-9"
  project = "debian-cloud"
}

provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}
`, context)
}

func TestAccComputeAutoscaler_autoscalerBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAutoscalerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAutoscaler_autoscalerBasicExample(context),
			},
			{
				ResourceName:      "google_compute_autoscaler.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeAutoscaler_autoscalerBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_autoscaler" "foobar" {
  name   = "my-autoscaler%{random_suffix}"
  zone   = "us-central1-f"
  target = google_compute_instance_group_manager.foobar.self_link

  autoscaling_policy {
    max_replicas    = 5
    min_replicas    = 1
    cooldown_period = 60

    cpu_utilization {
      target = 0.5
    }
  }
}

resource "google_compute_instance_template" "foobar" {
  name           = "my-instance-template%{random_suffix}"
  machine_type   = "n1-standard-1"
  can_ip_forward = false

  tags = ["foo", "bar"]

  disk {
    source_image = data.google_compute_image.debian_9.self_link
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_target_pool" "foobar" {
  name = "my-target-pool%{random_suffix}"
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "my-igm%{random_suffix}"
  zone = "us-central1-f"

  version {
    instance_template  = google_compute_instance_template.foobar.self_link
    name               = "primary"
  }

  target_pools       = [google_compute_target_pool.foobar.self_link]
  base_instance_name = "foobar"
}

data "google_compute_image" "debian_9" {
  family  = "debian-9"
  project = "debian-cloud"
}
`, context)
}

func testAccCheckComputeAutoscalerDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_autoscaler" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeAutoscaler still exists at %s", url)
		}
	}

	return nil
}
