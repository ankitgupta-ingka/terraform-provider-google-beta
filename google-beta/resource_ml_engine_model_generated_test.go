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

func TestAccMLEngineModel_mlModelBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMLEngineModelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMLEngineModel_mlModelBasicExample(context),
			},
			{
				ResourceName:      "google_ml_engine_model.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMLEngineModel_mlModelBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_ml_engine_model" "default" {
  name = "default"
  description = "My model"
  regions = ["us-central1"]
}
`, context)
}

func testAccCheckMLEngineModelDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_ml_engine_model" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{MLEngineBasePath}}projects/{{project}}/models/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("MLEngineModel still exists at %s", url)
		}
	}

	return nil
}
