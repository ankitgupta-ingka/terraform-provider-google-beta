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

func TestAccAppEngineStandardAppVersion_appEngineStandardAppVersionExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppEngineStandardAppVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineStandardAppVersion_appEngineStandardAppVersionExample(context),
			},
			{
				ResourceName:      "google_app_engine_standard_app_version.name",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppEngineStandardAppVersion_appEngineStandardAppVersionExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name       = "tf-test-project"
  project_id = "test-project-%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_app_engine_application" "app" {
  project     = "${google_project.my_project.project_id}"
  location_id = "us-central"
}


resource "google_app_engine_app_version" "name" {
  version_id = "v31"
  project = "${google_app_engine_application.app.project}/services/default"
  runtime = "nodejs10"
  api_config = {}
  endpoints_api_service = {}
  entrypoint {
    shell = "node index.js"
  }
`, context)
}

func testAccCheckAppEngineStandardAppVersionDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_app_engine_standard_app_version" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}/versions/{{versionId}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("AppEngineStandardAppVersion still exists at %s", url)
		}
	}

	return nil
}
