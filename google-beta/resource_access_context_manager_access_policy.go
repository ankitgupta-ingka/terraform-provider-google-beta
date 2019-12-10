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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAccessContextManagerAccessPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerAccessPolicyCreate,
		Read:   resourceAccessContextManagerAccessPolicyRead,
		Update: resourceAccessContextManagerAccessPolicyUpdate,
		Delete: resourceAccessContextManagerAccessPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerAccessPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The parent of this AccessPolicy in the Cloud Resource Hierarchy.
Format: organizations/{organization_id}`,
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human readable title. Does not affect behavior.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was created in UTC.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Resource name of the AccessPolicy. Format: {policy_id}`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was updated in UTC.`,
			},
		},
	}
}

func resourceAccessContextManagerAccessPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	parentProp, err := expandAccessContextManagerAccessPolicyParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	titleProp, err := expandAccessContextManagerAccessPolicyTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AccessPolicy: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AccessPolicy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = accessContextManagerOperationWaitTime(
		config, res, "Creating AccessPolicy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// Remove ID to show resource wasn't created.
		d.SetId("")
		return fmt.Errorf("Error waiting to create AccessPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AccessPolicy %q: %#v", d.Id(), res)

	// The operation for this resource contains the generated name that we need
	// in order to perform a READ. We need to access the object inside of it as
	// a map[string]interface, so let's do that.

	resp := res["response"].(map[string]interface{})
	name := GetResourceNameFromSelfLink(resp["name"].(string))
	log.Printf("[DEBUG] Setting AccessPolicy name, id to %s", name)
	d.Set("name", name)
	d.SetId(name)

	return resourceAccessContextManagerAccessPolicyRead(d, meta)
}

func resourceAccessContextManagerAccessPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerAccessPolicy %q", d.Id()))
	}

	if err := d.Set("name", flattenAccessContextManagerAccessPolicyName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("create_time", flattenAccessContextManagerAccessPolicyCreateTime(res["createTime"], d)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenAccessContextManagerAccessPolicyUpdateTime(res["updateTime"], d)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("parent", flattenAccessContextManagerAccessPolicyParent(res["parent"], d)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("title", flattenAccessContextManagerAccessPolicyTitle(res["title"], d)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}

	return nil
}

func resourceAccessContextManagerAccessPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerAccessPolicyTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AccessPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("title") {
		updateMask = append(updateMask, "title")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AccessPolicy %q: %s", d.Id(), err)
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Updating AccessPolicy",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerAccessPolicyRead(d, meta)
}

func resourceAccessContextManagerAccessPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AccessPolicy %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AccessPolicy")
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Deleting AccessPolicy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AccessPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerAccessPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
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

func flattenAccessContextManagerAccessPolicyName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenAccessContextManagerAccessPolicyCreateTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessPolicyUpdateTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessPolicyParent(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessPolicyTitle(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandAccessContextManagerAccessPolicyParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessPolicyTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
