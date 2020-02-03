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
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAccessContextManagerServicePerimeter() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerServicePerimeterCreate,
		Read:   resourceAccessContextManagerServicePerimeterRead,
		Update: resourceAccessContextManagerServicePerimeterUpdate,
		Delete: resourceAccessContextManagerServicePerimeterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerServicePerimeterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Resource name for the ServicePerimeter. The short_name component must
begin with a letter and only include alphanumeric and '_'.
Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}`,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The AccessPolicy this ServicePerimeter lives in.
Format: accessPolicies/{policy_id}`,
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human readable title. Must be unique within the Policy.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Description of the ServicePerimeter and its use. Does not affect
behavior.`,
			},
			"perimeter_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE", ""}, false),
				Description: `Specifies the type of the Perimeter. There are two types: regular and
bridge. Regular Service Perimeter contains resources, access levels,
and restricted services. Every resource can be in at most
ONE regular Service Perimeter.

In addition to being in a regular service perimeter, a resource can also
be in zero or more perimeter bridges. A perimeter bridge only contains
resources. Cross project operations are permitted if all effected
resources share some perimeter (whether bridge or regular). Perimeter
Bridge does not contain access levels or services: those are governed
entirely by the regular perimeter that resource is in.

Perimeter Bridges are typically useful when building more complex
topologies with many independent perimeters that need to share some data
with a common perimeter, but should not be able to share data among
themselves.`,
				Default: "PERIMETER_TYPE_REGULAR",
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `ServicePerimeter configuration. Specifies sets of resources,
restricted services and access levels that determine
perimeter content and boundaries.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_levels": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of AccessLevel resource names that allow resources within
the ServicePerimeter to be accessed from the internet.
AccessLevels listed must be in the same policy as this
ServicePerimeter. Referencing a nonexistent AccessLevel is a
syntax error. If no AccessLevel names are listed, resources within
the perimeter can only be accessed via GCP calls with request
origins within the perimeter. For Service Perimeter Bridge, must
be empty.

Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							AtLeastOneOf: []string{"status.0.resources", "status.0.access_levels", "status.0.restricted_services"},
						},
						"resources": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of GCP resources that are inside of the service perimeter.
Currently only projects are allowed.
Format: projects/{project_number}`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							AtLeastOneOf: []string{"status.0.resources", "status.0.access_levels", "status.0.restricted_services"},
						},
						"restricted_services": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `GCP services that are subject to the Service Perimeter
restrictions. Must contain a list of services. For example, if
'storage.googleapis.com' is specified, access to the storage
buckets inside the perimeter must meet the perimeter's access
restrictions.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							AtLeastOneOf: []string{"status.0.resources", "status.0.access_levels", "status.0.restricted_services"},
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was created in UTC.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was updated in UTC.`,
			},
		},
	}
}

func resourceAccessContextManagerServicePerimeterCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerServicePerimeterTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerServicePerimeterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	perimeterTypeProp, err := expandAccessContextManagerServicePerimeterPerimeterType(d.Get("perimeter_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("perimeter_type"); !isEmptyValue(reflect.ValueOf(perimeterTypeProp)) && (ok || !reflect.DeepEqual(v, perimeterTypeProp)) {
		obj["perimeterType"] = perimeterTypeProp
	}
	statusProp, err := expandAccessContextManagerServicePerimeterStatus(d.Get("status"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(statusProp)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}
	parentProp, err := expandAccessContextManagerServicePerimeterParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerServicePerimeterName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceAccessContextManagerServicePerimeterEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{parent}}/servicePerimeters")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServicePerimeter: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServicePerimeter: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = accessContextManagerOperationWaitTime(
		config, res, "Creating ServicePerimeter",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServicePerimeter: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServicePerimeter %q: %#v", d.Id(), res)

	return resourceAccessContextManagerServicePerimeterRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerServicePerimeter %q", d.Id()))
	}

	if err := d.Set("title", flattenAccessContextManagerServicePerimeterTitle(res["title"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("description", flattenAccessContextManagerServicePerimeterDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("create_time", flattenAccessContextManagerServicePerimeterCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("update_time", flattenAccessContextManagerServicePerimeterUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("perimeter_type", flattenAccessContextManagerServicePerimeterPerimeterType(res["perimeterType"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("status", flattenAccessContextManagerServicePerimeterStatus(res["status"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("name", flattenAccessContextManagerServicePerimeterName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}

	return nil
}

func resourceAccessContextManagerServicePerimeterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerServicePerimeterTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerServicePerimeterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	statusProp, err := expandAccessContextManagerServicePerimeterStatus(d.Get("status"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}

	obj, err = resourceAccessContextManagerServicePerimeterEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServicePerimeter %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("title") {
		updateMask = append(updateMask, "title")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("status") {
		updateMask = append(updateMask, "status")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ServicePerimeter %q: %s", d.Id(), err)
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Updating ServicePerimeter",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerServicePerimeterRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ServicePerimeter %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ServicePerimeter")
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Deleting ServicePerimeter",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServicePerimeter %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerServicePerimeterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}
	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) < 2 {
		return nil, fmt.Errorf("Error parsing parent name. Should be in form accessPolicies/{{policy_id}}/servicePerimeters/{{short_name}}")
	}
	d.Set("parent", fmt.Sprintf("%s/%s", stringParts[0], stringParts[1]))
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerServicePerimeterTitle(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterPerimeterType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil || isEmptyValue(reflect.ValueOf(v)) {
		return "PERIMETER_TYPE_REGULAR"
	}

	return v
}

func flattenAccessContextManagerServicePerimeterStatus(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resources"] =
		flattenAccessContextManagerServicePerimeterStatusResources(original["resources"], d, config)
	transformed["access_levels"] =
		flattenAccessContextManagerServicePerimeterStatusAccessLevels(original["accessLevels"], d, config)
	transformed["restricted_services"] =
		flattenAccessContextManagerServicePerimeterStatusRestrictedServices(original["restrictedServices"], d, config)
	return []interface{}{transformed}
}
func flattenAccessContextManagerServicePerimeterStatusResources(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterStatusAccessLevels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterStatusRestrictedServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAccessContextManagerServicePerimeterTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterPerimeterType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatus(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResources, err := expandAccessContextManagerServicePerimeterStatusResources(original["resources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !isEmptyValue(val) {
		transformed["resources"] = transformedResources
	}

	transformedAccessLevels, err := expandAccessContextManagerServicePerimeterStatusAccessLevels(original["access_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["accessLevels"] = transformedAccessLevels
	}

	transformedRestrictedServices, err := expandAccessContextManagerServicePerimeterStatusRestrictedServices(original["restricted_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRestrictedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["restrictedServices"] = transformedRestrictedServices
	}

	return transformed, nil
}

func expandAccessContextManagerServicePerimeterStatusResources(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatusAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatusRestrictedServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceAccessContextManagerServicePerimeterEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}
