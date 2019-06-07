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

var FirestoreDefaultBasePath = "https://firestore.googleapis.com/v1/"

var FirestoreBasePathEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_FIRESTORE_BASE_PATH",
	}, FirestoreDefaultBasePath),
}

var GeneratedFirestoreResourcesMap = map[string]*schema.Resource{
	"google_firestore_index": resourceFirestoreIndex(),
}
