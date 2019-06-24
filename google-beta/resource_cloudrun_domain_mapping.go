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
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCloudrunDomainMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudrunDomainMappingCreate,
		Read:   resourceCloudrunDomainMappingRead,
		Update: resourceCloudrunDomainMappingUpdate,
		Delete: resourceCloudrunDomainMappingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudrunDomainMappingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"annotations": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"labels": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"generation": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"resource_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"self_link": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"spec": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"CERTIFICATE_MODE_UNSPECIFIED", "NONE", "AUTOMATIC", ""}, false),
						},
						"force_override": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"route_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"conditions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"reason": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"mapped_route_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"observed_generation": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"resource_records": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"rrdata": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"RECORD_TYPE_UNSPECIFIED", "A", "AAAA", "CNAME", ""}, false),
									},
								},
							},
						},
					},
				},
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

func resourceCloudrunDomainMappingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	statusProp, err := expandCloudrunDomainMappingStatus(d.Get("status"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(statusProp)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}
	apiVersionProp, err := expandCloudrunDomainMappingApiVersion(d.Get("api_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_version"); !isEmptyValue(reflect.ValueOf(apiVersionProp)) && (ok || !reflect.DeepEqual(v, apiVersionProp)) {
		obj["apiVersion"] = apiVersionProp
	}
	specProp, err := expandCloudrunDomainMappingSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !isEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	metadataProp, err := expandCloudrunDomainMappingMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := replaceVars(d, config, "{{CloudrunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DomainMapping: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating DomainMapping: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DomainMapping %q: %#v", d.Id(), res)

	return resourceCloudrunDomainMappingRead(d, meta)
}

func resourceCloudrunDomainMappingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{CloudrunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudrunDomainMapping %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	if err := d.Set("status", flattenCloudrunDomainMappingStatus(res["status"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("api_version", flattenCloudrunDomainMappingApiVersion(res["apiVersion"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("spec", flattenCloudrunDomainMappingSpec(res["spec"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("metadata", flattenCloudrunDomainMappingMetadata(res["metadata"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	return nil
}

func resourceCloudrunDomainMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	statusProp, err := expandCloudrunDomainMappingStatus(d.Get("status"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}
	apiVersionProp, err := expandCloudrunDomainMappingApiVersion(d.Get("api_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, apiVersionProp)) {
		obj["apiVersion"] = apiVersionProp
	}
	specProp, err := expandCloudrunDomainMappingSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	metadataProp, err := expandCloudrunDomainMappingMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := replaceVars(d, config, "{{CloudrunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DomainMapping %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating DomainMapping %q: %s", d.Id(), err)
	}

	return resourceCloudrunDomainMappingRead(d, meta)
}

func resourceCloudrunDomainMappingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{CloudrunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DomainMapping %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "DomainMapping")
	}

	log.Printf("[DEBUG] Finished deleting DomainMapping %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudrunDomainMappingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/domainmappings/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)", "(?P<location>[^/]+)/(?P<name>[^/]+)"}, d, config); err != nil {
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

func flattenCloudrunDomainMappingStatus(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["conditions"] =
		flattenCloudrunDomainMappingStatusConditions(original["conditions"], d)
	transformed["observed_generation"] =
		flattenCloudrunDomainMappingStatusObservedGeneration(original["observedGeneration"], d)
	transformed["resource_records"] =
		flattenCloudrunDomainMappingStatusResourceRecords(original["resourceRecords"], d)
	transformed["mapped_route_name"] =
		flattenCloudrunDomainMappingStatusMappedRouteName(original["mappedRouteName"], d)
	return []interface{}{transformed}
}
func flattenCloudrunDomainMappingStatusConditions(v interface{}, d *schema.ResourceData) interface{} {
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
			"message": flattenCloudrunDomainMappingStatusConditionsMessage(original["message"], d),
			"status":  flattenCloudrunDomainMappingStatusConditionsStatus(original["status"], d),
			"reason":  flattenCloudrunDomainMappingStatusConditionsReason(original["reason"], d),
			"type":    flattenCloudrunDomainMappingStatusConditionsType(original["type"], d),
		})
	}
	return transformed
}
func flattenCloudrunDomainMappingStatusConditionsMessage(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingStatusConditionsStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingStatusConditionsReason(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingStatusConditionsType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingStatusObservedGeneration(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudrunDomainMappingStatusResourceRecords(v interface{}, d *schema.ResourceData) interface{} {
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
			"type":   flattenCloudrunDomainMappingStatusResourceRecordsType(original["type"], d),
			"rrdata": flattenCloudrunDomainMappingStatusResourceRecordsRrdata(original["rrdata"], d),
			"name":   flattenCloudrunDomainMappingStatusResourceRecordsName(original["name"], d),
		})
	}
	return transformed
}
func flattenCloudrunDomainMappingStatusResourceRecordsType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingStatusResourceRecordsRrdata(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingStatusResourceRecordsName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingStatusMappedRouteName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingApiVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingSpec(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["force_override"] =
		flattenCloudrunDomainMappingSpecForceOverride(original["forceOverride"], d)
	transformed["route_name"] =
		flattenCloudrunDomainMappingSpecRouteName(original["routeName"], d)
	transformed["certificate_mode"] =
		flattenCloudrunDomainMappingSpecCertificateMode(original["certificateMode"], d)
	return []interface{}{transformed}
}
func flattenCloudrunDomainMappingSpecForceOverride(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingSpecRouteName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingSpecCertificateMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingMetadata(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["labels"] =
		flattenCloudrunDomainMappingMetadataLabels(original["labels"], d)
	transformed["generation"] =
		flattenCloudrunDomainMappingMetadataGeneration(original["generation"], d)
	transformed["resource_version"] =
		flattenCloudrunDomainMappingMetadataResourceVersion(original["resourceVersion"], d)
	transformed["self_link"] =
		flattenCloudrunDomainMappingMetadataSelfLink(original["selfLink"], d)
	transformed["uid"] =
		flattenCloudrunDomainMappingMetadataUid(original["uid"], d)
	transformed["namespace"] =
		flattenCloudrunDomainMappingMetadataNamespace(original["namespace"], d)
	transformed["annotations"] =
		flattenCloudrunDomainMappingMetadataAnnotations(original["annotations"], d)
	transformed["name"] =
		flattenCloudrunDomainMappingMetadataName(original["name"], d)
	return []interface{}{transformed}
}
func flattenCloudrunDomainMappingMetadataLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingMetadataGeneration(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudrunDomainMappingMetadataResourceVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingMetadataSelfLink(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingMetadataUid(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingMetadataNamespace(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingMetadataAnnotations(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudrunDomainMappingMetadataName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandCloudrunDomainMappingStatus(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConditions, err := expandCloudrunDomainMappingStatusConditions(original["conditions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConditions); val.IsValid() && !isEmptyValue(val) {
		transformed["conditions"] = transformedConditions
	}

	transformedObservedGeneration, err := expandCloudrunDomainMappingStatusObservedGeneration(original["observed_generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedObservedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["observedGeneration"] = transformedObservedGeneration
	}

	transformedResourceRecords, err := expandCloudrunDomainMappingStatusResourceRecords(original["resource_records"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceRecords); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceRecords"] = transformedResourceRecords
	}

	transformedMappedRouteName, err := expandCloudrunDomainMappingStatusMappedRouteName(original["mapped_route_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMappedRouteName); val.IsValid() && !isEmptyValue(val) {
		transformed["mappedRouteName"] = transformedMappedRouteName
	}

	return transformed, nil
}

func expandCloudrunDomainMappingStatusConditions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMessage, err := expandCloudrunDomainMappingStatusConditionsMessage(original["message"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMessage); val.IsValid() && !isEmptyValue(val) {
			transformed["message"] = transformedMessage
		}

		transformedStatus, err := expandCloudrunDomainMappingStatusConditionsStatus(original["status"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStatus); val.IsValid() && !isEmptyValue(val) {
			transformed["status"] = transformedStatus
		}

		transformedReason, err := expandCloudrunDomainMappingStatusConditionsReason(original["reason"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedReason); val.IsValid() && !isEmptyValue(val) {
			transformed["reason"] = transformedReason
		}

		transformedType, err := expandCloudrunDomainMappingStatusConditionsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudrunDomainMappingStatusConditionsMessage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusConditionsStatus(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusConditionsReason(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusConditionsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusObservedGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusResourceRecords(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandCloudrunDomainMappingStatusResourceRecordsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedRrdata, err := expandCloudrunDomainMappingStatusResourceRecordsRrdata(original["rrdata"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRrdata); val.IsValid() && !isEmptyValue(val) {
			transformed["rrdata"] = transformedRrdata
		}

		transformedName, err := expandCloudrunDomainMappingStatusResourceRecordsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudrunDomainMappingStatusResourceRecordsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusResourceRecordsRrdata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusResourceRecordsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingStatusMappedRouteName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingApiVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedForceOverride, err := expandCloudrunDomainMappingSpecForceOverride(original["force_override"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedForceOverride); val.IsValid() && !isEmptyValue(val) {
		transformed["forceOverride"] = transformedForceOverride
	}

	transformedRouteName, err := expandCloudrunDomainMappingSpecRouteName(original["route_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRouteName); val.IsValid() && !isEmptyValue(val) {
		transformed["routeName"] = transformedRouteName
	}

	transformedCertificateMode, err := expandCloudrunDomainMappingSpecCertificateMode(original["certificate_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateMode); val.IsValid() && !isEmptyValue(val) {
		transformed["certificateMode"] = transformedCertificateMode
	}

	return transformed, nil
}

func expandCloudrunDomainMappingSpecForceOverride(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingSpecRouteName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingSpecCertificateMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandCloudrunDomainMappingMetadataLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedGeneration, err := expandCloudrunDomainMappingMetadataGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	transformedResourceVersion, err := expandCloudrunDomainMappingMetadataResourceVersion(original["resource_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceVersion"] = transformedResourceVersion
	}

	transformedSelfLink, err := expandCloudrunDomainMappingMetadataSelfLink(original["self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["selfLink"] = transformedSelfLink
	}

	transformedUid, err := expandCloudrunDomainMappingMetadataUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !isEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	transformedNamespace, err := expandCloudrunDomainMappingMetadataNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	transformedAnnotations, err := expandCloudrunDomainMappingMetadataAnnotations(original["annotations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnnotations); val.IsValid() && !isEmptyValue(val) {
		transformed["annotations"] = transformedAnnotations
	}

	transformedName, err := expandCloudrunDomainMappingMetadataName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandCloudrunDomainMappingMetadataLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudrunDomainMappingMetadataGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingMetadataResourceVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingMetadataSelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingMetadataUid(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingMetadataNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudrunDomainMappingMetadataAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudrunDomainMappingMetadataName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
