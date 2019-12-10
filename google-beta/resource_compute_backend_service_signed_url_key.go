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
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceComputeBackendServiceSignedUrlKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendServiceSignedUrlKeyCreate,
		Read:   resourceComputeBackendServiceSignedUrlKeyRead,
		Delete: resourceComputeBackendServiceSignedUrlKeyDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"backend_service": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The backend service this signed URL key belongs.`,
			},
			"key_value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `128-bit key value used for signing the URL. The key value must be a
valid RFC 4648 Section 5 base64url encoded string.`,
				Sensitive: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
				Description:  `Name of the signed URL key.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceComputeBackendServiceSignedUrlKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	keyNameProp, err := expandComputeBackendServiceSignedUrlKeyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(keyNameProp)) && (ok || !reflect.DeepEqual(v, keyNameProp)) {
		obj["keyName"] = keyNameProp
	}
	keyValueProp, err := expandComputeBackendServiceSignedUrlKeyKeyValue(d.Get("key_value"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key_value"); !isEmptyValue(reflect.ValueOf(keyValueProp)) && (ok || !reflect.DeepEqual(v, keyValueProp)) {
		obj["keyValue"] = keyValueProp
	}
	backendServiceProp, err := expandComputeBackendServiceSignedUrlKeyBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(backendServiceProp)) && (ok || !reflect.DeepEqual(v, backendServiceProp)) {
		obj["backendService"] = backendServiceProp
	}

	lockName, err := replaceVars(d, config, "signedUrlKey/{{project}}/backendServices/{{backend_service}}/")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendServices/{{backend_service}}/addSignedUrlKey")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackendServiceSignedUrlKey: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BackendServiceSignedUrlKey: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/backendServices/{{backend_service}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating BackendServiceSignedUrlKey",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// Remove ID to show resource wasn't created.
		d.SetId("")
		return fmt.Errorf("Error waiting to create BackendServiceSignedUrlKey: %s", err)
	}

	log.Printf("[DEBUG] Finished creating BackendServiceSignedUrlKey %q: %#v", d.Id(), res)

	return resourceComputeBackendServiceSignedUrlKeyRead(d, meta)
}

func resourceComputeBackendServiceSignedUrlKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendServices/{{backend_service}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeBackendServiceSignedUrlKey %q", d.Id()))
	}

	res, err = flattenNestedComputeBackendServiceSignedUrlKey(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeBackendServiceSignedUrlKey because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackendServiceSignedUrlKey: %s", err)
	}

	if err := d.Set("name", flattenComputeBackendServiceSignedUrlKeyName(res["keyName"], d)); err != nil {
		return fmt.Errorf("Error reading BackendServiceSignedUrlKey: %s", err)
	}

	return nil
}

func resourceComputeBackendServiceSignedUrlKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "signedUrlKey/{{project}}/backendServices/{{backend_service}}/")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendServices/{{backend_service}}/deleteSignedUrlKey?keyName={{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BackendServiceSignedUrlKey %q", d.Id())

	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BackendServiceSignedUrlKey")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting BackendServiceSignedUrlKey",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BackendServiceSignedUrlKey %q: %#v", d.Id(), res)
	return nil
}

func flattenComputeBackendServiceSignedUrlKeyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeBackendServiceSignedUrlKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceSignedUrlKeyKeyValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceSignedUrlKeyBackendService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("backendServices", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for backend_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func flattenNestedComputeBackendServiceSignedUrlKey(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["cdnPolicy"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})

	v, ok = res["signedUrlKeyNames"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value cdnPolicy.signedUrlKeyNames. Actual value: %v", v)
	}

	_, item, err := resourceComputeBackendServiceSignedUrlKeyFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeBackendServiceSignedUrlKeyFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName, err := expandComputeBackendServiceSignedUrlKeyName(d.Get("name"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		// List response only contains the ID - construct a response object.
		item := map[string]interface{}{
			"keyName": itemRaw,
		}

		itemName := flattenComputeBackendServiceSignedUrlKeyName(item["keyName"], d)
		if !reflect.DeepEqual(itemName, expectedName) {
			log.Printf("[DEBUG] Skipping item with keyName= %#v, looking for %#v)", itemName, expectedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
