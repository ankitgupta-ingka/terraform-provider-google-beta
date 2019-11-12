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

func resourceFirestoreIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirestoreIndexCreate,
		Read:   resourceFirestoreIndexRead,
		Delete: resourceFirestoreIndexDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirestoreIndexImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"collection": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"array_config": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"CONTAINS", ""}, false),
						},
						"field_path": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"order": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"ASCENDING", "DESCENDING", ""}, false),
						},
					},
				},
			},
			"database": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "(default)",
			},
			"query_scope": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"COLLECTION", "COLLECTION_GROUP", ""}, false),
				Default:      "COLLECTION",
			},
			"name": {
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

func resourceFirestoreIndexCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	databaseProp, err := expandFirestoreIndexDatabase(d.Get("database"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database"); !isEmptyValue(reflect.ValueOf(databaseProp)) && (ok || !reflect.DeepEqual(v, databaseProp)) {
		obj["database"] = databaseProp
	}
	collectionProp, err := expandFirestoreIndexCollection(d.Get("collection"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collection"); !isEmptyValue(reflect.ValueOf(collectionProp)) && (ok || !reflect.DeepEqual(v, collectionProp)) {
		obj["collection"] = collectionProp
	}
	queryScopeProp, err := expandFirestoreIndexQueryScope(d.Get("query_scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("query_scope"); !isEmptyValue(reflect.ValueOf(queryScopeProp)) && (ok || !reflect.DeepEqual(v, queryScopeProp)) {
		obj["queryScope"] = queryScopeProp
	}
	fieldsProp, err := expandFirestoreIndexFields(d.Get("fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fields"); !isEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	obj, err = resourceFirestoreIndexEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Index: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Index: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := firestoreOperationWaitTime(
		config, res, project, "Creating Index",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Index: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Index %q: %#v", d.Id(), res)

	// The operation for this resource contains the generated name that we need
	// in order to perform a READ.
	metadata := res["metadata"].(map[string]interface{})
	name := metadata["index"].(string)
	log.Printf("[DEBUG] Setting Index name, id to %s", name)
	d.Set("name", name)
	d.SetId(name)

	return resourceFirestoreIndexRead(d, meta)
}

func resourceFirestoreIndexRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{FirestoreBasePath}}{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirestoreIndex %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	if err := d.Set("name", flattenFirestoreIndexName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("query_scope", flattenFirestoreIndexQueryScope(res["queryScope"], d)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("fields", flattenFirestoreIndexFields(res["fields"], d)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	return nil
}

func resourceFirestoreIndexDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{FirestoreBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Index %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Index")
	}

	err = firestoreOperationWaitTime(
		config, res, project, "Deleting Index",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Index %q: %#v", d.Id(), res)
	return nil
}

func resourceFirestoreIndexImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) != 8 {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s",
			d.Get("name"),
			"projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}",
		)
	}

	d.Set("project", fmt.Sprintf("%s", stringParts[1]))
	return []*schema.ResourceData{d}, nil
}

func flattenFirestoreIndexName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFirestoreIndexQueryScope(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFirestoreIndexFields(v interface{}, d *schema.ResourceData) interface{} {
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
			"field_path":   flattenFirestoreIndexFieldsFieldPath(original["fieldPath"], d),
			"order":        flattenFirestoreIndexFieldsOrder(original["order"], d),
			"array_config": flattenFirestoreIndexFieldsArrayConfig(original["arrayConfig"], d),
		})
	}
	return transformed
}
func flattenFirestoreIndexFieldsFieldPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFirestoreIndexFieldsOrder(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenFirestoreIndexFieldsArrayConfig(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandFirestoreIndexDatabase(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexCollection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexQueryScope(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFields(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFieldPath, err := expandFirestoreIndexFieldsFieldPath(original["field_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFieldPath); val.IsValid() && !isEmptyValue(val) {
			transformed["fieldPath"] = transformedFieldPath
		}

		transformedOrder, err := expandFirestoreIndexFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !isEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedArrayConfig, err := expandFirestoreIndexFieldsArrayConfig(original["array_config"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArrayConfig); val.IsValid() && !isEmptyValue(val) {
			transformed["arrayConfig"] = transformedArrayConfig
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirestoreIndexFieldsFieldPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsOrder(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsArrayConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceFirestoreIndexEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// We've added project / database / collection as split fields of the name, but
	// the API doesn't expect them.  Make sure we remove them from any requests.

	delete(obj, "project")
	delete(obj, "database")
	delete(obj, "collection")
	return obj, nil
}
