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

func TestAccPubsubSubscription_pubsubSubscriptionPullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPubsubSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionPullExample(context),
			},
			{
				ResourceName:      "google_pubsub_subscription.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionPullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "example-topic-%{random_suffix}"
}

resource "google_pubsub_subscription" "example" {
  name  = "example-subscription-%{random_suffix}"
  topic = "${google_pubsub_topic.example.name}"

  labels = {
    foo = "bar"
  }

  # 20 minutes
  message_retention_duration = "1200s"
  retain_acked_messages = true

  ack_deadline_seconds = 20

  expiration_policy {
    ttl = "300000.5s"
  }
}
`, context)
}

func testAccCheckPubsubSubscriptionDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_pubsub_subscription" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://pubsub.googleapis.com/v1/projects/{{project}}/subscriptions/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("PubsubSubscription still exists at %s", url)
		}
	}

	return nil
}
