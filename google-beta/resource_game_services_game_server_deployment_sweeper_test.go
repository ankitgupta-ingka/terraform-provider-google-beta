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
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("GameServicesGameServerDeployment", &resource.Sweeper{
		Name: "GameServicesGameServerDeployment",
		F:    testSweepGameServicesGameServerDeployment,
	})
}

// At the time of writing, the CI only passes us-central1 as the region
func testSweepGameServicesGameServerDeployment(region string) error {
	resourceName := "GameServicesGameServerDeployment"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	// Setup variables to replace in list template
	d := &ResourceDataMock{
		FieldsInSchema: map[string]interface{}{
			"project":  config.Project,
			"region":   region,
			"location": region,
			"zone":     "-",
		},
	}

	listTemplate := strings.Split("https://gameservices.googleapis.com/v1beta/projects/{{project}}/locations/{{location}}/gameServerDeployments", "?")[0]
	listUrl, err := replaceVars(d, config, listTemplate)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error preparing sweeper list url: %s", err)
		return nil
	}

	res, err := sendRequest(config, "GET", config.Project, listUrl, nil)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", listUrl, err)
		return nil
	}

	resourceList, ok := res["gameServerDeployments"]
	if !ok {
		log.Printf("[INFO][SWEEPER_LOG] Nothing found in response.")
		return nil
	}

	rl := resourceList.([]interface{})

	log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
	// Keep count of items that aren't sweepable for logging.
	nonPrefixCount := 0
	for _, ri := range rl {
		obj := ri.(map[string]interface{})
		var name string
		// Id detected in the delete URL, attempt to use id.
		if obj["id"] != nil {
			name = GetResourceNameFromSelfLink(obj["id"].(string))
		} else if obj["name"] != nil {
			name = GetResourceNameFromSelfLink(obj["name"].(string))
		} else {
			log.Printf("[INFO][SWEEPER_LOG] %s resource name and id were nil", resourceName)
			return nil
		}
		// Skip resources that shouldn't be sweeped
		if !isSweepableTestResource(name) {
			nonPrefixCount++
			continue
		}

		deleteTemplate := "https://gameservices.googleapis.com/v1beta/projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}"
		deleteUrl, err := replaceVars(d, config, deleteTemplate)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error preparing delete url: %s", err)
			return nil
		}
		deleteUrl = deleteUrl + name

		// Don't wait on operations as we may have a lot to delete
		_, err = sendRequest(config, "DELETE", config.Project, deleteUrl, nil)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", deleteUrl, err)
		} else {
			log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, name)
		}
	}

	if nonPrefixCount > 0 {
		log.Printf("[INFO][SWEEPER_LOG] %d items were non-sweepable and skipped.", nonPrefixCount)
	}

	return nil
}
