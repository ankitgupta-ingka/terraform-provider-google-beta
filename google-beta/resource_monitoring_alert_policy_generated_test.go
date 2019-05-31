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

func TestAccMonitoringAlertPolicy_monitoringAlertPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitoringAlertPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringAlertPolicy_monitoringAlertPolicyBasicExample(context),
			},
			{
				ResourceName:      "google_monitoring_alert_policy.alert_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringAlertPolicy_monitoringAlertPolicyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_alert_policy" "alert_policy" {
  display_name = "My Alert Policy-%{random_suffix}"
  combiner = "OR"
  conditions {
    display_name = "test condition"
    condition_threshold {
      filter = "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\""
      duration = "60s"
      comparison = "COMPARISON_GT"
      aggregations {
        alignment_period = "60s"
        per_series_aligner = "ALIGN_RATE"
      }
    }
  }

  user_labels = {
    foo = "bar"
  }
}
`, context)
}

func testAccCheckMonitoringAlertPolicyDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_monitoring_alert_policy" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://monitoring.googleapis.com/v3/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("MonitoringAlertPolicy still exists at %s", url)
		}
	}

	return nil
}
