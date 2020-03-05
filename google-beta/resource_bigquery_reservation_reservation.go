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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBigqueryReservationReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationReservationCreate,
		Read:   resourceBigqueryReservationReservationRead,
		Update: resourceBigqueryReservationReservationUpdate,
		Delete: resourceBigqueryReservationReservationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationReservationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the reservation. This field must only contain alphanumeric characters or dash.`,
			},
			"slot_capacity": {
				Type:     schema.TypeInt,
				Required: true,
				Description: `Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.`,
			},
			"ignore_idle_slots": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If false, any query using this reservation will use idle slots from other reservations within
the same admin project. If true, a query using this reservation will execute with the slot
capacity specified above at most.`,
				Default: false,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the transfer config should reside.
Examples: US, EU, asia-northeast1. The default value is US.`,
				Default: "US",
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

func resourceBigqueryReservationReservationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	slotCapacityProp, err := expandBigqueryReservationReservationSlotCapacity(d.Get("slot_capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("slot_capacity"); !isEmptyValue(reflect.ValueOf(slotCapacityProp)) && (ok || !reflect.DeepEqual(v, slotCapacityProp)) {
		obj["slotCapacity"] = slotCapacityProp
	}
	ignoreIdleSlotsProp, err := expandBigqueryReservationReservationIgnoreIdleSlots(d.Get("ignore_idle_slots"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ignore_idle_slots"); !isEmptyValue(reflect.ValueOf(ignoreIdleSlotsProp)) && (ok || !reflect.DeepEqual(v, ignoreIdleSlotsProp)) {
		obj["ignoreIdleSlots"] = ignoreIdleSlotsProp
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations?reservationId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Reservation: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Reservation: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Reservation %q: %#v", d.Id(), res)

	return resourceBigqueryReservationReservationRead(d, meta)
}

func resourceBigqueryReservationReservationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigqueryReservationReservation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	if err := d.Set("slot_capacity", flattenBigqueryReservationReservationSlotCapacity(res["slotCapacity"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("ignore_idle_slots", flattenBigqueryReservationReservationIgnoreIdleSlots(res["ignoreIdleSlots"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	return nil
}

func resourceBigqueryReservationReservationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	slotCapacityProp, err := expandBigqueryReservationReservationSlotCapacity(d.Get("slot_capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("slot_capacity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, slotCapacityProp)) {
		obj["slotCapacity"] = slotCapacityProp
	}
	ignoreIdleSlotsProp, err := expandBigqueryReservationReservationIgnoreIdleSlots(d.Get("ignore_idle_slots"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ignore_idle_slots"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ignoreIdleSlotsProp)) {
		obj["ignoreIdleSlots"] = ignoreIdleSlotsProp
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Reservation %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("slot_capacity") {
		updateMask = append(updateMask, "slotCapacity")
	}

	if d.HasChange("ignore_idle_slots") {
		updateMask = append(updateMask, "ignoreIdleSlots")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Reservation %q: %s", d.Id(), err)
	}

	return resourceBigqueryReservationReservationRead(d, meta)
}

func resourceBigqueryReservationReservationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Reservation %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Reservation")
	}

	log.Printf("[DEBUG] Finished deleting Reservation %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryReservationReservationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/reservations/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryReservationReservationSlotCapacity(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenBigqueryReservationReservationIgnoreIdleSlots(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandBigqueryReservationReservationSlotCapacity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationIgnoreIdleSlots(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
