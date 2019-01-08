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
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	filestore "google.golang.org/api/file/v1beta1"
)

func resourceFilestoreInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceFilestoreInstanceCreate,
		Read:   resourceFilestoreInstanceRead,
		Update: resourceFilestoreInstanceUpdate,
		Delete: resourceFilestoreInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFilestoreInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(360 * time.Second),
			Update: schema.DefaultTimeout(360 * time.Second),
			Delete: schema.DefaultTimeout(360 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"file_shares": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capacity_gb": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modes": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"reserved_ip_range": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"ip_addresses": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"tier": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"TIER_UNSPECIFIED", "STANDARD", "PREMIUM"}, false),
			},
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_mask": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceFilestoreInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tierProp, err := expandFilestoreInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}
	labelsProp, err := expandFilestoreInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	fileSharesProp, err := expandFilestoreInstanceFileShares(d.Get("file_shares"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("file_shares"); !isEmptyValue(reflect.ValueOf(fileSharesProp)) && (ok || !reflect.DeepEqual(v, fileSharesProp)) {
		obj["fileShares"] = fileSharesProp
	}
	networksProp, err := expandFilestoreInstanceNetworks(d.Get("networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("networks"); !isEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}

	url, err := replaceVars(d, config, "https://file.googleapis.com/v1beta1/projects/{{project}}/locations/{{zone}}/instances?instanceId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{zone}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &filestore.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := filestoreOperationWaitTime(
		config.clientFilestore, op, project, "Creating Instance",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceFilestoreInstanceRead(d, meta)
}

func resourceFilestoreInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://file.googleapis.com/v1beta1/projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FilestoreInstance %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("name", flattenFilestoreInstanceName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("description", flattenFilestoreInstanceDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenFilestoreInstanceCreateTime(res["createTime"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("tier", flattenFilestoreInstanceTier(res["tier"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenFilestoreInstanceLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("file_shares", flattenFilestoreInstanceFileShares(res["fileShares"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("networks", flattenFilestoreInstanceNetworks(res["networks"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("etag", flattenFilestoreInstanceEtag(res["etag"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("update_mask", flattenFilestoreInstanceUpdateMask(res["updateMask"], d)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceFilestoreInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandFilestoreInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	fileSharesProp, err := expandFilestoreInstanceFileShares(d.Get("file_shares"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("file_shares"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fileSharesProp)) {
		obj["fileShares"] = fileSharesProp
	}

	url, err := replaceVars(d, config, "https://file.googleapis.com/v1beta1/projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("file_shares") {
		updateMask = append(updateMask, "fileShares")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &filestore.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = filestoreOperationWaitTime(
		config.clientFilestore, op, project, "Updating Instance",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceFilestoreInstanceRead(d, meta)
}

func resourceFilestoreInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://file.googleapis.com/v1beta1/projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &filestore.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = filestoreOperationWaitTime(
		config.clientFilestore, op, project, "Deleting Instance",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceFilestoreInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/locations/(?P<zone>[^/]+)/instances/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{zone}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFilestoreInstanceName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenFilestoreInstanceDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceCreateTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceTier(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceFileShares(v interface{}, d *schema.ResourceData) interface{} {
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
			"name":        flattenFilestoreInstanceFileSharesName(original["name"], d),
			"capacity_gb": flattenFilestoreInstanceFileSharesCapacityGb(original["capacityGb"], d),
		})
	}
	return transformed
}
func flattenFilestoreInstanceFileSharesName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceFileSharesCapacityGb(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenFilestoreInstanceNetworks(v interface{}, d *schema.ResourceData) interface{} {
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
			"network":           flattenFilestoreInstanceNetworksNetwork(original["network"], d),
			"modes":             flattenFilestoreInstanceNetworksModes(original["modes"], d),
			"reserved_ip_range": flattenFilestoreInstanceNetworksReservedIpRange(original["reservedIpRange"], d),
			"ip_addresses":      flattenFilestoreInstanceNetworksIpAddresses(original["ipAddresses"], d),
		})
	}
	return transformed
}
func flattenFilestoreInstanceNetworksNetwork(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceNetworksModes(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceNetworksReservedIpRange(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceNetworksIpAddresses(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceEtag(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFilestoreInstanceUpdateMask(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandFilestoreInstanceDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceTier(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFilestoreInstanceFileShares(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandFilestoreInstanceFileSharesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedCapacityGb, err := expandFilestoreInstanceFileSharesCapacityGb(original["capacity_gb"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCapacityGb); val.IsValid() && !isEmptyValue(val) {
			transformed["capacityGb"] = transformedCapacityGb
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceFileSharesName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesCapacityGb(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworks(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandFilestoreInstanceNetworksNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !isEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		transformedModes, err := expandFilestoreInstanceNetworksModes(original["modes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedModes); val.IsValid() && !isEmptyValue(val) {
			transformed["modes"] = transformedModes
		}

		transformedReservedIpRange, err := expandFilestoreInstanceNetworksReservedIpRange(original["reserved_ip_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedReservedIpRange); val.IsValid() && !isEmptyValue(val) {
			transformed["reservedIpRange"] = transformedReservedIpRange
		}

		transformedIpAddresses, err := expandFilestoreInstanceNetworksIpAddresses(original["ip_addresses"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddresses"] = transformedIpAddresses
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceNetworksNetwork(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksModes(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksReservedIpRange(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksIpAddresses(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
