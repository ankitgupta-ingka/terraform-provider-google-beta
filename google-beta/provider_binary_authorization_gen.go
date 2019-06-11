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

import "github.com/hashicorp/terraform/helper/schema"

// If the base path has changed as a result of your PR, make sure to update
// the provider_reference page!
var BinaryAuthorizationDefaultBasePath = "https://binaryauthorization.googleapis.com/v1beta1/"
var BinaryAuthorizationCustomEndpointEntryKey = "binary_authorization_custom_endpoint"
var BinaryAuthorizationCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_BINARY_AUTHORIZATION_CUSTOM_ENDPOINT",
	}, BinaryAuthorizationDefaultBasePath),
}

var GeneratedBinaryAuthorizationResourcesMap = map[string]*schema.Resource{
	"google_binary_authorization_attestor": resourceBinaryAuthorizationAttestor(),
	"google_binary_authorization_policy":   resourceBinaryAuthorizationPolicy(),
}
