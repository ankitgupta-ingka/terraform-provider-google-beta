// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//     ***     DIFF TEST DIFF TEST    ***
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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccPubsubTopicIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopicIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s roles/viewer", getTestProjectFromEnv(), fmt.Sprintf("example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccPubsubTopicIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s roles/viewer", getTestProjectFromEnv(), fmt.Sprintf("example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPubsubTopicIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccPubsubTopicIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPubsubTopicIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopicIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s", getTestProjectFromEnv(), fmt.Sprintf("example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubTopicIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }
}

resource "google_pubsub_topic_iam_member" "foo" {
  project = "${google_pubsub_topic.example.project}"
  topic = "${google_pubsub_topic.example.name}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccPubsubTopicIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_pubsub_topic_iam_policy" "foo" {
  project = "${google_pubsub_topic.example.project}"
  topic = "${google_pubsub_topic.example.name}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccPubsubTopicIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }
}

resource "google_pubsub_topic_iam_binding" "foo" {
  project = "${google_pubsub_topic.example.project}"
  topic = "${google_pubsub_topic.example.name}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccPubsubTopicIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }
}

resource "google_pubsub_topic_iam_binding" "foo" {
  project = "${google_pubsub_topic.example.project}"
  topic = "${google_pubsub_topic.example.name}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
