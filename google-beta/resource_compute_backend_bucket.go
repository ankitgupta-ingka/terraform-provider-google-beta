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
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeBackendBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendBucketCreate,
		Read:   resourceComputeBackendBucketRead,
		Update: resourceComputeBackendBucketUpdate,
		Delete: resourceComputeBackendBucketDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeBackendBucketImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
			},
			"cdn_policy": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signed_url_cache_max_age_sec": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  3600,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_cdn": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeBackendBucketCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketNameProp, err := expandComputeBackendBucketBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket_name"); !isEmptyValue(reflect.ValueOf(bucketNameProp)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	cdnPolicyProp, err := expandComputeBackendBucketCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cdn_policy"); !isEmptyValue(reflect.ValueOf(cdnPolicyProp)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	descriptionProp, err := expandComputeBackendBucketDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCdnProp, err := expandComputeBackendBucketEnableCdn(d.Get("enable_cdn"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_cdn"); !isEmptyValue(reflect.ValueOf(enableCdnProp)) && (ok || !reflect.DeepEqual(v, enableCdnProp)) {
		obj["enableCdn"] = enableCdnProp
	}
	nameProp, err := expandComputeBackendBucketName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/backendBuckets")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackendBucket: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BackendBucket: %s", err)
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
		config.clientCompute, op, project, "Creating BackendBucket",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create BackendBucket: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating BackendBucket %q: %#v", d.Id(), res)

	return resourceComputeBackendBucketRead(d, meta)
}

func resourceComputeBackendBucketRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeBackendBucket %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}

	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("bucket_name", flattenComputeBackendBucketBucketName(res["bucketName"], d)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("cdn_policy", flattenComputeBackendBucketCdnPolicy(res["cdnPolicy"], d)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("creation_timestamp", flattenComputeBackendBucketCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("description", flattenComputeBackendBucketDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("enable_cdn", flattenComputeBackendBucketEnableCdn(res["enableCdn"], d)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("name", flattenComputeBackendBucketName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}

	return nil
}

func resourceComputeBackendBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketNameProp, err := expandComputeBackendBucketBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	cdnPolicyProp, err := expandComputeBackendBucketCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cdn_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	descriptionProp, err := expandComputeBackendBucketDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCdnProp, err := expandComputeBackendBucketEnableCdn(d.Get("enable_cdn"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_cdn"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableCdnProp)) {
		obj["enableCdn"] = enableCdnProp
	}
	nameProp, err := expandComputeBackendBucketName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BackendBucket %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating BackendBucket %q: %s", d.Id(), err)
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
		config.clientCompute, op, project, "Updating BackendBucket",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeBackendBucketRead(d, meta)
}

func resourceComputeBackendBucketDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BackendBucket %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BackendBucket")
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
		config.clientCompute, op, project, "Deleting BackendBucket",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BackendBucket %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeBackendBucketImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/backendBuckets/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeBackendBucketBucketName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeBackendBucketCdnPolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["signed_url_cache_max_age_sec"] =
		flattenComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(original["signedUrlCacheMaxAgeSec"], d)
	return []interface{}{transformed}
}
func flattenComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeBackendBucketCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeBackendBucketDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeBackendBucketEnableCdn(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeBackendBucketName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeBackendBucketBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSignedUrlCacheMaxAgeSec, err := expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(original["signed_url_cache_max_age_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignedUrlCacheMaxAgeSec); val.IsValid() && !isEmptyValue(val) {
		transformed["signedUrlCacheMaxAgeSec"] = transformedSignedUrlCacheMaxAgeSec
	}

	return transformed, nil
}

func expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketEnableCdn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
