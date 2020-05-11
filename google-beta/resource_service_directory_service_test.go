package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccServiceDirectoryService_serviceDirectoryServiceUpdateExample(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	location := "us-central1"
	testId := fmt.Sprintf("tf-test-example-service%s", randString(t, 10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDirectoryServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryService_basic(location, testId),
			},
			{
				ResourceName:      "google_service_directory_service.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_service.example",
				// {{project}}/{{location}}/{{namespace_id}}/{{service_id}}
				ImportStateId:     fmt.Sprintf("%s/%s/%s/%s", project, location, testId, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_service.example",
				// {{location}}/{{namespace_id}}/{{service_id}}
				ImportStateId:     fmt.Sprintf("%s/%s/%s", location, testId, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccServiceDirectoryService_update(location, testId),
			},
			{
				ResourceName:      "google_service_directory_service.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccServiceDirectoryService_basic(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"
}

resource "google_service_directory_service" "example" {
  service_id = "%s"
  namespace  = google_service_directory_namespace.example.id
}
`, testId, location, testId)
}

func testAccServiceDirectoryService_update(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"
}

resource "google_service_directory_service" "example" {
  service_id = "%s"
  namespace  = google_service_directory_namespace.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}
`, testId, location, testId)
}
