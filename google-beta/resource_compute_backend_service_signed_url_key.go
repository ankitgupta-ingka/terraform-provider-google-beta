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

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeBackendServiceSignedUrlKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendServiceSignedUrlKeyCreate,
		Read:   resourceComputeBackendServiceSignedUrlKeyRead,
		Delete: resourceComputeBackendServiceSignedUrlKeyDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"backend_service": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"key_value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
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

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/backendServices/{{backend_service}}/addSignedUrlKey")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackendServiceSignedUrlKey: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BackendServiceSignedUrlKey: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating BackendServiceSignedUrlKey",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create BackendServiceSignedUrlKey: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating BackendServiceSignedUrlKey %q: %#v", d.Id(), res)

	return resourceComputeBackendServiceSignedUrlKeyRead(d, meta)
}

func resourceComputeBackendServiceSignedUrlKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/backendServices/{{backend_service}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
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

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackendServiceSignedUrlKey: %s", err)
	}

	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("name", flattenComputeBackendServiceSignedUrlKeyName(res["keyName"], d)); err != nil {
		return fmt.Errorf("Error reading BackendServiceSignedUrlKey: %s", err)
	}

	return nil
}

func resourceComputeBackendServiceSignedUrlKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lockName, err := replaceVars(d, config, "signedUrlKey/{{project}}/backendServices/{{backend_service}}/")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/backendServices/{{backend_service}}/deleteSignedUrlKey?keyName={{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BackendServiceSignedUrlKey %q", d.Id())
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BackendServiceSignedUrlKey")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting BackendServiceSignedUrlKey",
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

	// Final nested resource is either a list of resources we need to filter
	// or just the resource itself, which we return.
	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		return v.(map[string]interface{}), nil
	default:
		return nil, fmt.Errorf("invalid value for cdnPolicy.signedUrlKeyNames: %v", v)
	}

	items := v.([]interface{})
	for _, vRaw := range items {
		// If only an id is given in parent resource,
		// construct a resource map for that id KV pair.
		item := map[string]interface{}{"keyName": vRaw}
		itemIdV, err := expandComputeBackendServiceSignedUrlKeyName(d.Get("name"), d, meta.(*Config))
		if err != nil {
			return nil, err
		}
		actualIdV := flattenComputeBackendServiceSignedUrlKeyName(item["keyName"], d)
		log.Printf("[DEBUG] Checking if item's keyName (%#v) is equal to resource's (%#v)", itemIdV, actualIdV)
		if !reflect.DeepEqual(itemIdV, actualIdV) {
			continue
		}
		return item, nil
	}
	return nil, nil
}
