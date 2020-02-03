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

func resourceCloudTasksQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudTasksQueueCreate,
		Read:   resourceCloudTasksQueueRead,
		Update: resourceCloudTasksQueueUpdate,
		Delete: resourceCloudTasksQueueDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudTasksQueueImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the queue`,
			},
			"app_engine_routing_override": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Overrides for task-level appEngineRouting. These settings apply only
to App Engine tasks in this queue`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `App instance.

By default, the task is sent to an instance which is available when the task is attempted.`,
						},
						"service": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `App service.

By default, the task is sent to the service which is the default service when the task is attempted.`,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `App version.

By default, the task is sent to the version which is the default version when the task is attempted.`,
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The host that the task is sent to.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The queue name.`,
			},
			"rate_limits": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Description: `Rate limits for task dispatches.

The queue's actual dispatch rate is the result of:

* Number of tasks in the queue
* User-specified throttling: rateLimits, retryConfig, and the queue's state.
* System throttling due to 429 (Too Many Requests) or 503 (Service
  Unavailable) responses from the worker, high error rates, or to
  smooth sudden large traffic spikes.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_concurrent_dispatches": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `The maximum number of concurrent tasks that Cloud Tasks allows to
be dispatched for this queue. After this threshold has been
reached, Cloud Tasks stops dispatching tasks until the number of
concurrent requests decreases.`,
						},
						"max_dispatches_per_second": {
							Type:     schema.TypeFloat,
							Computed: true,
							Optional: true,
							Description: `The maximum rate at which tasks are dispatched from this queue.

If unspecified when the queue is created, Cloud Tasks will pick the default.`,
						},
						"max_burst_size": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: `The max burst size.

Max burst size limits how fast tasks in queue are processed when many tasks are
in the queue and the rate is high. This field allows the queue to have a high
rate so processing starts shortly after a task is enqueued, but still limits
resource usage when many tasks are enqueued in a short period of time.`,
						},
					},
				},
			},
			"retry_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Settings that determine the retry behavior.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `Number of attempts per task.

Cloud Tasks will attempt the task maxAttempts times (that is, if
the first attempt fails, then there will be maxAttempts - 1
retries). Must be >= -1.

If unspecified when the queue is created, Cloud Tasks will pick
the default.

-1 indicates unlimited attempts.`,
						},
						"max_backoff": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `A task will be scheduled for retry between minBackoff and
maxBackoff duration after it fails, if the queue's RetryConfig
specifies that the task should be retried.`,
						},
						"max_doublings": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `The time between retries will double maxDoublings times.

A task's retry interval starts at minBackoff, then doubles maxDoublings times,
then increases linearly, and finally retries retries at intervals of maxBackoff
up to maxAttempts times.`,
						},
						"max_retry_duration": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `If positive, maxRetryDuration specifies the time limit for
retrying a failed task, measured from when the task was first
attempted. Once maxRetryDuration time has passed and the task has
been attempted maxAttempts times, no further attempts will be
made and the task will be deleted.

If zero, then the task age is unlimited.`,
						},
						"min_backoff": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `A task will be scheduled for retry between minBackoff and
maxBackoff duration after it fails, if the queue's RetryConfig
specifies that the task should be retried.`,
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

func resourceCloudTasksQueueCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandCloudTasksQueueName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	appEngineRoutingOverrideProp, err := expandCloudTasksQueueAppEngineRoutingOverride(d.Get("app_engine_routing_override"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_routing_override"); !isEmptyValue(reflect.ValueOf(appEngineRoutingOverrideProp)) && (ok || !reflect.DeepEqual(v, appEngineRoutingOverrideProp)) {
		obj["appEngineRoutingOverride"] = appEngineRoutingOverrideProp
	}
	rateLimitsProp, err := expandCloudTasksQueueRateLimits(d.Get("rate_limits"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rate_limits"); !isEmptyValue(reflect.ValueOf(rateLimitsProp)) && (ok || !reflect.DeepEqual(v, rateLimitsProp)) {
		obj["rateLimits"] = rateLimitsProp
	}
	retryConfigProp, err := expandCloudTasksQueueRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retry_config"); !isEmptyValue(reflect.ValueOf(retryConfigProp)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}

	url, err := replaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Queue: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Queue: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Queue %q: %#v", d.Id(), res)

	return resourceCloudTasksQueueRead(d, meta)
}

func resourceCloudTasksQueueRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudTasksQueue %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}

	if err := d.Set("name", flattenCloudTasksQueueName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}
	if err := d.Set("app_engine_routing_override", flattenCloudTasksQueueAppEngineRoutingOverride(res["appEngineRoutingOverride"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}
	if err := d.Set("rate_limits", flattenCloudTasksQueueRateLimits(res["rateLimits"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}
	if err := d.Set("retry_config", flattenCloudTasksQueueRetryConfig(res["retryConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}

	return nil
}

func resourceCloudTasksQueueUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	appEngineRoutingOverrideProp, err := expandCloudTasksQueueAppEngineRoutingOverride(d.Get("app_engine_routing_override"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_routing_override"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, appEngineRoutingOverrideProp)) {
		obj["appEngineRoutingOverride"] = appEngineRoutingOverrideProp
	}
	rateLimitsProp, err := expandCloudTasksQueueRateLimits(d.Get("rate_limits"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rate_limits"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, rateLimitsProp)) {
		obj["rateLimits"] = rateLimitsProp
	}
	retryConfigProp, err := expandCloudTasksQueueRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retry_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}

	url, err := replaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Queue %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("app_engine_routing_override") {
		updateMask = append(updateMask, "appEngineRoutingOverride")
	}

	if d.HasChange("rate_limits") {
		updateMask = append(updateMask, "rateLimits")
	}

	if d.HasChange("retry_config") {
		updateMask = append(updateMask, "retryConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Queue %q: %s", d.Id(), err)
	}

	return resourceCloudTasksQueueRead(d, meta)
}

func resourceCloudTasksQueueDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Queue %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Queue")
	}

	log.Printf("[DEBUG] Finished deleting Queue %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudTasksQueueImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/queues/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudTasksQueueName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

// service, version, and instance are input-only. host is output-only.
func flattenCloudTasksQueueAppEngineRoutingOverride(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] = original["host"]
	if override, ok := d.GetOk("app_engine_routing_override"); ok && len(override.([]interface{})) > 0 {
		transformed["service"] = d.Get("app_engine_routing_override.0.service")
		transformed["version"] = d.Get("app_engine_routing_override.0.version")
		transformed["instance"] = d.Get("app_engine_routing_override.0.instance")
	}
	return []interface{}{transformed}
}

func flattenCloudTasksQueueRateLimits(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["max_dispatches_per_second"] =
		flattenCloudTasksQueueRateLimitsMaxDispatchesPerSecond(original["maxDispatchesPerSecond"], d, config)
	transformed["max_concurrent_dispatches"] =
		flattenCloudTasksQueueRateLimitsMaxConcurrentDispatches(original["maxConcurrentDispatches"], d, config)
	transformed["max_burst_size"] =
		flattenCloudTasksQueueRateLimitsMaxBurstSize(original["maxBurstSize"], d, config)
	return []interface{}{transformed}
}
func flattenCloudTasksQueueRateLimitsMaxDispatchesPerSecond(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudTasksQueueRateLimitsMaxConcurrentDispatches(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudTasksQueueRateLimitsMaxBurstSize(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudTasksQueueRetryConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["max_attempts"] =
		flattenCloudTasksQueueRetryConfigMaxAttempts(original["maxAttempts"], d, config)
	transformed["max_retry_duration"] =
		flattenCloudTasksQueueRetryConfigMaxRetryDuration(original["maxRetryDuration"], d, config)
	transformed["min_backoff"] =
		flattenCloudTasksQueueRetryConfigMinBackoff(original["minBackoff"], d, config)
	transformed["max_backoff"] =
		flattenCloudTasksQueueRetryConfigMaxBackoff(original["maxBackoff"], d, config)
	transformed["max_doublings"] =
		flattenCloudTasksQueueRetryConfigMaxDoublings(original["maxDoublings"], d, config)
	return []interface{}{transformed}
}
func flattenCloudTasksQueueRetryConfigMaxAttempts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudTasksQueueRetryConfigMaxRetryDuration(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudTasksQueueRetryConfigMinBackoff(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudTasksQueueRetryConfigMaxBackoff(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudTasksQueueRetryConfigMaxDoublings(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func expandCloudTasksQueueName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{location}}/queues/{{name}}")
}

func expandCloudTasksQueueAppEngineRoutingOverride(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandCloudTasksQueueAppEngineRoutingOverrideService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandCloudTasksQueueAppEngineRoutingOverrideVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedInstance, err := expandCloudTasksQueueAppEngineRoutingOverrideInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !isEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedHost, err := expandCloudTasksQueueAppEngineRoutingOverrideHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	return transformed, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimits(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxDispatchesPerSecond, err := expandCloudTasksQueueRateLimitsMaxDispatchesPerSecond(original["max_dispatches_per_second"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDispatchesPerSecond); val.IsValid() && !isEmptyValue(val) {
		transformed["maxDispatchesPerSecond"] = transformedMaxDispatchesPerSecond
	}

	transformedMaxConcurrentDispatches, err := expandCloudTasksQueueRateLimitsMaxConcurrentDispatches(original["max_concurrent_dispatches"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxConcurrentDispatches); val.IsValid() && !isEmptyValue(val) {
		transformed["maxConcurrentDispatches"] = transformedMaxConcurrentDispatches
	}

	transformedMaxBurstSize, err := expandCloudTasksQueueRateLimitsMaxBurstSize(original["max_burst_size"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBurstSize); val.IsValid() && !isEmptyValue(val) {
		transformed["maxBurstSize"] = transformedMaxBurstSize
	}

	return transformed, nil
}

func expandCloudTasksQueueRateLimitsMaxDispatchesPerSecond(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimitsMaxConcurrentDispatches(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimitsMaxBurstSize(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxAttempts, err := expandCloudTasksQueueRetryConfigMaxAttempts(original["max_attempts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxAttempts); val.IsValid() && !isEmptyValue(val) {
		transformed["maxAttempts"] = transformedMaxAttempts
	}

	transformedMaxRetryDuration, err := expandCloudTasksQueueRetryConfigMaxRetryDuration(original["max_retry_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetryDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["maxRetryDuration"] = transformedMaxRetryDuration
	}

	transformedMinBackoff, err := expandCloudTasksQueueRetryConfigMinBackoff(original["min_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinBackoff); val.IsValid() && !isEmptyValue(val) {
		transformed["minBackoff"] = transformedMinBackoff
	}

	transformedMaxBackoff, err := expandCloudTasksQueueRetryConfigMaxBackoff(original["max_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBackoff); val.IsValid() && !isEmptyValue(val) {
		transformed["maxBackoff"] = transformedMaxBackoff
	}

	transformedMaxDoublings, err := expandCloudTasksQueueRetryConfigMaxDoublings(original["max_doublings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDoublings); val.IsValid() && !isEmptyValue(val) {
		transformed["maxDoublings"] = transformedMaxDoublings
	}

	return transformed, nil
}

func expandCloudTasksQueueRetryConfigMaxAttempts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxRetryDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMinBackoff(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxBackoff(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxDoublings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
