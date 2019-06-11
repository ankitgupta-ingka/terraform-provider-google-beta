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

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAccessContextManagerAccessLevel() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerAccessLevelCreate,
		Read:   resourceAccessContextManagerAccessLevelRead,
		Update: resourceAccessContextManagerAccessLevelUpdate,
		Delete: resourceAccessContextManagerAccessLevelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerAccessLevelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(360 * time.Second),
			Update: schema.DefaultTimeout(360 * time.Second),
			Delete: schema.DefaultTimeout(360 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"basic": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"conditions": {
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_policy": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allowed_device_management_levels": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_encryption_statuses": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"os_constraints": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"minimum_version": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"os_type": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice([]string{"OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", ""}, false),
															},
														},
													},
												},
												"require_screen_lock": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
									"ip_subnetworks": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"members": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"negate": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"required_access_levels": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"combining_function": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"AND", "OR", ""}, false),
							Default:      "AND",
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAccessContextManagerAccessLevelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerAccessLevelTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerAccessLevelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	basicProp, err := expandAccessContextManagerAccessLevelBasic(d.Get("basic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic"); !isEmptyValue(reflect.ValueOf(basicProp)) && (ok || !reflect.DeepEqual(v, basicProp)) {
		obj["basic"] = basicProp
	}
	parentProp, err := expandAccessContextManagerAccessLevelParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerAccessLevelName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceAccessContextManagerAccessLevelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1/{{parent}}/accessLevels")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AccessLevel: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AccessLevel: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := accessContextManagerOperationWaitTime(
		config, res, "Creating AccessLevel",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create AccessLevel: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating AccessLevel %q: %#v", d.Id(), res)

	return resourceAccessContextManagerAccessLevelRead(d, meta)
}

func resourceAccessContextManagerAccessLevelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerAccessLevel %q", d.Id()))
	}

	if err := d.Set("title", flattenAccessContextManagerAccessLevelTitle(res["title"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}
	if err := d.Set("description", flattenAccessContextManagerAccessLevelDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}
	if err := d.Set("basic", flattenAccessContextManagerAccessLevelBasic(res["basic"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}
	if err := d.Set("name", flattenAccessContextManagerAccessLevelName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}

	return nil
}

func resourceAccessContextManagerAccessLevelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerAccessLevelTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerAccessLevelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	basicProp, err := expandAccessContextManagerAccessLevelBasic(d.Get("basic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, basicProp)) {
		obj["basic"] = basicProp
	}

	obj, err = resourceAccessContextManagerAccessLevelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AccessLevel %q: %#v", d.Id(), obj)

	updateMask := []string{}

	if d.HasChange("title") {
		updateMask = append(updateMask, "title")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("basic") {
		updateMask = append(updateMask, "basic")
	}

	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AccessLevel %q: %s", d.Id(), err)
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Updating AccessLevel",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerAccessLevelRead(d, meta)
}

func resourceAccessContextManagerAccessLevelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AccessLevel %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AccessLevel")
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Deleting AccessLevel",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AccessLevel %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerAccessLevelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}
	stringParts := strings.Split(d.Get("name").(string), "/")
	d.Set("parent", fmt.Sprintf("%s/%s", stringParts[0], stringParts[1]))
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerAccessLevelTitle(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasic(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["combining_function"] =
		flattenAccessContextManagerAccessLevelBasicCombiningFunction(original["combiningFunction"], d)
	transformed["conditions"] =
		flattenAccessContextManagerAccessLevelBasicConditions(original["conditions"], d)
	return []interface{}{transformed}
}
func flattenAccessContextManagerAccessLevelBasicCombiningFunction(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil || v.(string) == "" {
		return "AND"
	}
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditions(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"ip_subnetworks":         flattenAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(original["ipSubnetworks"], d),
			"required_access_levels": flattenAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(original["requiredAccessLevels"], d),
			"members":                flattenAccessContextManagerAccessLevelBasicConditionsMembers(original["members"], d),
			"negate":                 flattenAccessContextManagerAccessLevelBasicConditionsNegate(original["negate"], d),
			"device_policy":          flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicy(original["devicePolicy"], d),
		})
	}
	return transformed
}
func flattenAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsMembers(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsNegate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["require_screen_lock"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(original["requireScreenLock"], d)
	transformed["allowed_encryption_statuses"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(original["allowedEncryptionStatuses"], d)
	transformed["allowed_device_management_levels"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(original["allowedDeviceManagementLevels"], d)
	transformed["os_constraints"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(original["osConstraints"], d)
	return []interface{}{transformed}
}
func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"minimum_version": flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(original["minimumVersion"], d),
			"os_type":         flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(original["osType"], d),
		})
	}
	return transformed
}
func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandAccessContextManagerAccessLevelTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCombiningFunction, err := expandAccessContextManagerAccessLevelBasicCombiningFunction(original["combining_function"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCombiningFunction); val.IsValid() && !isEmptyValue(val) {
		transformed["combiningFunction"] = transformedCombiningFunction
	}

	transformedConditions, err := expandAccessContextManagerAccessLevelBasicConditions(original["conditions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConditions); val.IsValid() && !isEmptyValue(val) {
		transformed["conditions"] = transformedConditions
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelBasicCombiningFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpSubnetworks, err := expandAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(original["ip_subnetworks"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpSubnetworks); val.IsValid() && !isEmptyValue(val) {
			transformed["ipSubnetworks"] = transformedIpSubnetworks
		}

		transformedRequiredAccessLevels, err := expandAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(original["required_access_levels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRequiredAccessLevels); val.IsValid() && !isEmptyValue(val) {
			transformed["requiredAccessLevels"] = transformedRequiredAccessLevels
		}

		transformedMembers, err := expandAccessContextManagerAccessLevelBasicConditionsMembers(original["members"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMembers); val.IsValid() && !isEmptyValue(val) {
			transformed["members"] = transformedMembers
		}

		transformedNegate, err := expandAccessContextManagerAccessLevelBasicConditionsNegate(original["negate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNegate); val.IsValid() && !isEmptyValue(val) {
			transformed["negate"] = transformedNegate
		}

		transformedDevicePolicy, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicy(original["device_policy"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDevicePolicy); val.IsValid() && !isEmptyValue(val) {
			transformed["devicePolicy"] = transformedDevicePolicy
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsMembers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsNegate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequireScreenLock, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(original["require_screen_lock"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireScreenLock); val.IsValid() && !isEmptyValue(val) {
		transformed["requireScreenLock"] = transformedRequireScreenLock
	}

	transformedAllowedEncryptionStatuses, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(original["allowed_encryption_statuses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedEncryptionStatuses); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedEncryptionStatuses"] = transformedAllowedEncryptionStatuses
	}

	transformedAllowedDeviceManagementLevels, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(original["allowed_device_management_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedDeviceManagementLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedDeviceManagementLevels"] = transformedAllowedDeviceManagementLevels
	}

	transformedOsConstraints, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(original["os_constraints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOsConstraints); val.IsValid() && !isEmptyValue(val) {
		transformed["osConstraints"] = transformedOsConstraints
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMinimumVersion, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(original["minimum_version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMinimumVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["minimumVersion"] = transformedMinimumVersion
		}

		transformedOsType, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(original["os_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOsType); val.IsValid() && !isEmptyValue(val) {
			transformed["osType"] = transformedOsType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceAccessContextManagerAccessLevelEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}
