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
	"google.golang.org/api/appengine/v1"
)

func resourceAppEngineApplicationUrlDispatchRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineApplicationUrlDispatchRulesCreate,
		Read:   resourceAppEngineApplicationUrlDispatchRulesRead,
		Update: resourceAppEngineApplicationUrlDispatchRulesUpdate,
		Delete: resourceAppEngineApplicationUrlDispatchRulesDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineApplicationUrlDispatchRulesImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dispatch_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service": {
							Type:     schema.TypeString,
							Required: true,
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

func resourceAppEngineApplicationUrlDispatchRulesCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	dispatchRulesProp, err := expandAppEngineApplicationUrlDispatchRulesDispatchRules(d.Get("dispatch_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dispatch_rules"); !isEmptyValue(reflect.ValueOf(dispatchRulesProp)) && (ok || !reflect.DeepEqual(v, dispatchRulesProp)) {
		obj["dispatchRules"] = dispatchRulesProp
	}

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ApplicationUrlDispatchRules: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ApplicationUrlDispatchRules: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &appengine.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := appEngineOperationWaitTime(
		config.clientAppEngine, op, project, "Creating ApplicationUrlDispatchRules",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ApplicationUrlDispatchRules: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating ApplicationUrlDispatchRules %q: %#v", d.Id(), res)

	return resourceAppEngineApplicationUrlDispatchRulesRead(d, meta)
}

func resourceAppEngineApplicationUrlDispatchRulesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AppEngineApplicationUrlDispatchRules %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ApplicationUrlDispatchRules: %s", err)
	}

	if err := d.Set("dispatch_rules", flattenAppEngineApplicationUrlDispatchRulesDispatchRules(res["dispatchRules"], d)); err != nil {
		return fmt.Errorf("Error reading ApplicationUrlDispatchRules: %s", err)
	}

	return nil
}

func resourceAppEngineApplicationUrlDispatchRulesUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	dispatchRulesProp, err := expandAppEngineApplicationUrlDispatchRulesDispatchRules(d.Get("dispatch_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dispatch_rules"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dispatchRulesProp)) {
		obj["dispatchRules"] = dispatchRulesProp
	}

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ApplicationUrlDispatchRules %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ApplicationUrlDispatchRules %q: %s", d.Id(), err)
	}

	op := &appengine.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = appEngineOperationWaitTime(
		config.clientAppEngine, op, project, "Updating ApplicationUrlDispatchRules",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceAppEngineApplicationUrlDispatchRulesRead(d, meta)
}

func resourceAppEngineApplicationUrlDispatchRulesDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ApplicationUrlDispatchRules %q", d.Id())

	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ApplicationUrlDispatchRules")
	}

	op := &appengine.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = appEngineOperationWaitTime(
		config.clientAppEngine, op, project, "Deleting ApplicationUrlDispatchRules",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ApplicationUrlDispatchRules %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineApplicationUrlDispatchRulesImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRules(v interface{}, d *schema.ResourceData) interface{} {
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
			"domain":  flattenAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(original["domain"], d),
			"path":    flattenAppEngineApplicationUrlDispatchRulesDispatchRulesPath(original["path"], d),
			"service": flattenAppEngineApplicationUrlDispatchRulesDispatchRulesService(original["service"], d),
		})
	}
	return transformed
}
func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesService(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRules(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !isEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedPath, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedService, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
