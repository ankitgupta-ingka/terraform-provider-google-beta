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
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceStorageBucketAccessControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageBucketAccessControlCreate,
		Read:   resourceStorageBucketAccessControlRead,
		Update: resourceStorageBucketAccessControlUpdate,
		Delete: resourceStorageBucketAccessControlDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageBucketAccessControlImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"entity": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"OWNER", "READER", "WRITER", ""}, false),
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_team": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"team": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"editors", "owners", "viewers", ""}, false),
						},
					},
				},
			},
		},
	}
}

func resourceStorageBucketAccessControlCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageBucketAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageBucketAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	roleProp, err := expandStorageBucketAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BucketAccessControl: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BucketAccessControl: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{bucket}}/{{entity}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BucketAccessControl %q: %#v", d.Id(), res)

	return resourceStorageBucketAccessControlRead(d, meta)
}

func resourceStorageBucketAccessControlRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("StorageBucketAccessControl %q", d.Id()))
	}

	if err := d.Set("bucket", flattenStorageBucketAccessControlBucket(res["bucket"], d)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("domain", flattenStorageBucketAccessControlDomain(res["domain"], d)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("email", flattenStorageBucketAccessControlEmail(res["email"], d)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("entity", flattenStorageBucketAccessControlEntity(res["entity"], d)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("entity_id", flattenStorageBucketAccessControlEntityId(res["entityId"], d)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("project_team", flattenStorageBucketAccessControlProjectTeam(res["projectTeam"], d)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}
	if err := d.Set("role", flattenStorageBucketAccessControlRole(res["role"], d)); err != nil {
		return fmt.Errorf("Error reading BucketAccessControl: %s", err)
	}

	return nil
}

func resourceStorageBucketAccessControlUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageBucketAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageBucketAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	roleProp, err := expandStorageBucketAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BucketAccessControl %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating BucketAccessControl %q: %s", d.Id(), err)
	}

	return resourceStorageBucketAccessControlRead(d, meta)
}

func resourceStorageBucketAccessControlDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BucketAccessControl %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BucketAccessControl")
	}

	log.Printf("[DEBUG] Finished deleting BucketAccessControl %q: %#v", d.Id(), res)
	return nil
}

func resourceStorageBucketAccessControlImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<bucket>[^/]+)/(?P<entity>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{bucket}}/{{entity}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenStorageBucketAccessControlBucket(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenStorageBucketAccessControlDomain(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageBucketAccessControlEmail(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageBucketAccessControlEntity(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageBucketAccessControlEntityId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageBucketAccessControlProjectTeam(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project_number"] =
		flattenStorageBucketAccessControlProjectTeamProjectNumber(original["projectNumber"], d)
	transformed["team"] =
		flattenStorageBucketAccessControlProjectTeamTeam(original["team"], d)
	return []interface{}{transformed}
}
func flattenStorageBucketAccessControlProjectTeamProjectNumber(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageBucketAccessControlProjectTeamTeam(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageBucketAccessControlRole(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandStorageBucketAccessControlBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlEntity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
