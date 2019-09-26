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
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccContainerAnalysisNote_containerAnalysisNoteBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContainerAnalysisNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNote_containerAnalysisNoteBasicExample(context),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccContainerAnalysisNote_containerAnalysisNoteBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "test-attestor-note%{random_suffix}"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}
`, context)
}

func testAccCheckContainerAnalysisNoteDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_container_analysis_note" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ContainerAnalysisBasePath}}projects/{{project}}/notes/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ContainerAnalysisNote still exists at %s", url)
		}
	}

	return nil
}
