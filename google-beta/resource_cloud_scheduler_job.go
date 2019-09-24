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
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

// Both oidc and oauth headers cannot be set
func validateAuthHeaders(diff *schema.ResourceDiff, v interface{}) error {
	httpBlock := diff.Get("http_target.0").(map[string]interface{})

	if httpBlock != nil {
		oauth := httpBlock["oauth_token"]
		oidc := httpBlock["oidc_token"]

		if oauth != nil && oidc != nil {
			if len(oidc.([]interface{})) > 0 && len(oauth.([]interface{})) > 0 {
				return fmt.Errorf("Error in http_target: only one of oauth_token or oidc_token can be specified, but not both.")
			}
		}
	}

	return nil
}

func authHeaderDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// If generating an `oauth_token` and `scope` is not provided in the configuration,
	// the default "https://www.googleapis.com/auth/cloud-platform" scope will be used.
	// Similarly, if generating an `oidc_token` and `audience` is not provided in the
	// configuration, the URI specified in target will be used. Although not in the
	// configuration, in both cases the default is returned in the object, but is not in.
	// state. We suppress the diff if the values are these defaults but are not stored in state.

	b := strings.Split(k, ".")
	if b[0] == "http_target" && len(b) > 4 {
		block := b[2]
		attr := b[4]

		if block == "oauth_token" && attr == "scope" {
			if old == canonicalizeServiceScope("cloud-platform") && new == "" {
				return true
			}
		}

		if block == "oidc_token" && attr == "audience" {
			uri := d.Get(strings.Join(b[0:2], ".") + ".uri")
			if old == uri && new == "" {
				return true
			}
		}

	}

	return false
}

func resourceCloudSchedulerJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudSchedulerJobCreate,
		Read:   resourceCloudSchedulerJobRead,
		Delete: resourceCloudSchedulerJobDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudSchedulerJobImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		CustomizeDiff: validateAuthHeaders,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"app_engine_http_target": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relative_uri": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"app_engine_routing": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"service": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"version": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"body": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"headers": {
							Type:         schema.TypeMap,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validateHttpHeaders(),
							Elem:         &schema.Schema{Type: schema.TypeString},
						},
						"http_method": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
				ConflictsWith: []string{"pubsub_target", "http_target"},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"http_target": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"body": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"headers": {
							Type:         schema.TypeMap,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validateHttpHeaders(),
							Elem:         &schema.Schema{Type: schema.TypeString},
						},
						"http_method": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"oauth_token": {
							Type:             schema.TypeList,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: authHeaderDiffSuppress,
							MaxItems:         1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scope": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"service_account_email": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"oidc_token": {
							Type:             schema.TypeList,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: authHeaderDiffSuppress,
							MaxItems:         1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audience": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"service_account_email": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
				ConflictsWith: []string{"pubsub_target", "app_engine_http_target"},
			},
			"pubsub_target": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"attributes": {
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"data": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
				ConflictsWith: []string{"app_engine_http_target", "http_target"},
			},
			"retry_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_backoff_duration": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"max_doublings": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"max_retry_duration": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"min_backoff_duration": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"retry_count": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"schedule": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "Etc/UTC",
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

func resourceCloudSchedulerJobCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandCloudSchedulerJobName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandCloudSchedulerJobDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	scheduleProp, err := expandCloudSchedulerJobSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	timeZoneProp, err := expandCloudSchedulerJobTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !isEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	retryConfigProp, err := expandCloudSchedulerJobRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retry_config"); !isEmptyValue(reflect.ValueOf(retryConfigProp)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}
	pubsubTargetProp, err := expandCloudSchedulerJobPubsubTarget(d.Get("pubsub_target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pubsub_target"); !isEmptyValue(reflect.ValueOf(pubsubTargetProp)) && (ok || !reflect.DeepEqual(v, pubsubTargetProp)) {
		obj["pubsubTarget"] = pubsubTargetProp
	}
	appEngineHttpTargetProp, err := expandCloudSchedulerJobAppEngineHttpTarget(d.Get("app_engine_http_target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_http_target"); !isEmptyValue(reflect.ValueOf(appEngineHttpTargetProp)) && (ok || !reflect.DeepEqual(v, appEngineHttpTargetProp)) {
		obj["appEngineHttpTarget"] = appEngineHttpTargetProp
	}
	httpTargetProp, err := expandCloudSchedulerJobHttpTarget(d.Get("http_target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("http_target"); !isEmptyValue(reflect.ValueOf(httpTargetProp)) && (ok || !reflect.DeepEqual(v, httpTargetProp)) {
		obj["httpTarget"] = httpTargetProp
	}

	url, err := replaceVars(d, config, "{{CloudSchedulerBasePath}}projects/{{project}}/locations/{{region}}/jobs")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Job: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Job: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Job %q: %#v", d.Id(), res)

	return resourceCloudSchedulerJobRead(d, meta)
}

func resourceCloudSchedulerJobRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{CloudSchedulerBasePath}}projects/{{project}}/locations/{{region}}/jobs/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudSchedulerJob %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}

	if err := d.Set("name", flattenCloudSchedulerJobName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}
	if err := d.Set("description", flattenCloudSchedulerJobDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}
	if err := d.Set("schedule", flattenCloudSchedulerJobSchedule(res["schedule"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}
	if err := d.Set("time_zone", flattenCloudSchedulerJobTimeZone(res["timeZone"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}
	if err := d.Set("retry_config", flattenCloudSchedulerJobRetryConfig(res["retryConfig"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}
	if err := d.Set("pubsub_target", flattenCloudSchedulerJobPubsubTarget(res["pubsubTarget"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}
	if err := d.Set("app_engine_http_target", flattenCloudSchedulerJobAppEngineHttpTarget(res["appEngineHttpTarget"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}
	if err := d.Set("http_target", flattenCloudSchedulerJobHttpTarget(res["httpTarget"], d)); err != nil {
		return fmt.Errorf("Error reading Job: %s", err)
	}

	return nil
}

func resourceCloudSchedulerJobDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudSchedulerBasePath}}projects/{{project}}/locations/{{region}}/jobs/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Job %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Job")
	}

	log.Printf("[DEBUG] Finished deleting Job %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudSchedulerJobImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/jobs/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
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

func flattenCloudSchedulerJobName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenCloudSchedulerJobDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobSchedule(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobTimeZone(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobRetryConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["retry_count"] =
		flattenCloudSchedulerJobRetryConfigRetryCount(original["retryCount"], d)
	transformed["max_retry_duration"] =
		flattenCloudSchedulerJobRetryConfigMaxRetryDuration(original["maxRetryDuration"], d)
	transformed["min_backoff_duration"] =
		flattenCloudSchedulerJobRetryConfigMinBackoffDuration(original["minBackoffDuration"], d)
	transformed["max_backoff_duration"] =
		flattenCloudSchedulerJobRetryConfigMaxBackoffDuration(original["maxBackoffDuration"], d)
	transformed["max_doublings"] =
		flattenCloudSchedulerJobRetryConfigMaxDoublings(original["maxDoublings"], d)
	return []interface{}{transformed}
}
func flattenCloudSchedulerJobRetryConfigRetryCount(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudSchedulerJobRetryConfigMaxRetryDuration(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobRetryConfigMinBackoffDuration(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobRetryConfigMaxBackoffDuration(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobRetryConfigMaxDoublings(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudSchedulerJobPubsubTarget(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["topic_name"] =
		flattenCloudSchedulerJobPubsubTargetTopicName(original["topicName"], d)
	transformed["data"] =
		flattenCloudSchedulerJobPubsubTargetData(original["data"], d)
	transformed["attributes"] =
		flattenCloudSchedulerJobPubsubTargetAttributes(original["attributes"], d)
	return []interface{}{transformed}
}
func flattenCloudSchedulerJobPubsubTargetTopicName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobPubsubTargetData(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobPubsubTargetAttributes(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobAppEngineHttpTarget(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["http_method"] =
		flattenCloudSchedulerJobAppEngineHttpTargetHttpMethod(original["httpMethod"], d)
	transformed["app_engine_routing"] =
		flattenCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(original["appEngineRouting"], d)
	transformed["relative_uri"] =
		flattenCloudSchedulerJobAppEngineHttpTargetRelativeUri(original["relativeUri"], d)
	transformed["body"] =
		flattenCloudSchedulerJobAppEngineHttpTargetBody(original["body"], d)
	transformed["headers"] =
		flattenCloudSchedulerJobAppEngineHttpTargetHeaders(original["headers"], d)
	return []interface{}{transformed}
}
func flattenCloudSchedulerJobAppEngineHttpTargetHttpMethod(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

// An `appEngineRouting` in API response is useless, so we set config values rather than api response to state.
func flattenCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	if v, ok := d.GetOk("app_engine_http_target"); ok && len(v.([]interface{})) > 0 {
		return d.Get("app_engine_http_target.0.app_engine_routing")
	}
	return nil
}

func flattenCloudSchedulerJobAppEngineHttpTargetRelativeUri(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobAppEngineHttpTargetBody(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobAppEngineHttpTargetHeaders(v interface{}, d *schema.ResourceData) interface{} {
	var headers = v.(map[string]interface{})
	if v, ok := headers["User-Agent"]; ok {
		if v.(string) == "AppEngine-Google; (+http://code.google.com/appengine)" {
			delete(headers, "User-Agent")
		} else if v.(string) == "Google-Cloud-Scheduler" {
			delete(headers, "User-Agent")
		} else {
			headers["User-Agent"] = strings.TrimSpace(strings.Replace(v.(string), "AppEngine-Google; (+http://code.google.com/appengine)", "", -1))
		}
	}
	if v, ok := headers["Content-Type"]; ok {
		if v.(string) == "application/octet-stream" {
			delete(headers, "Content-Type")
		}
	}
	for key := range headers {
		match, _ := regexp.MatchString("(X-Google-|X-AppEngine-|Content-Length).*", key)
		if match {
			delete(headers, key)
		}
	}
	return headers
}

func flattenCloudSchedulerJobHttpTarget(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["uri"] =
		flattenCloudSchedulerJobHttpTargetUri(original["uri"], d)
	transformed["http_method"] =
		flattenCloudSchedulerJobHttpTargetHttpMethod(original["httpMethod"], d)
	transformed["body"] =
		flattenCloudSchedulerJobHttpTargetBody(original["body"], d)
	transformed["headers"] =
		flattenCloudSchedulerJobHttpTargetHeaders(original["headers"], d)
	transformed["oauth_token"] =
		flattenCloudSchedulerJobHttpTargetOauthToken(original["oauthToken"], d)
	transformed["oidc_token"] =
		flattenCloudSchedulerJobHttpTargetOidcToken(original["oidcToken"], d)
	return []interface{}{transformed}
}
func flattenCloudSchedulerJobHttpTargetUri(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobHttpTargetHttpMethod(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobHttpTargetBody(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobHttpTargetHeaders(v interface{}, d *schema.ResourceData) interface{} {
	var headers = v.(map[string]interface{})
	if v, ok := headers["User-Agent"]; ok {
		if v.(string) == "AppEngine-Google; (+http://code.google.com/appengine)" {
			delete(headers, "User-Agent")
		} else if v.(string) == "Google-Cloud-Scheduler" {
			delete(headers, "User-Agent")
		} else {
			headers["User-Agent"] = strings.TrimSpace(strings.Replace(v.(string), "AppEngine-Google; (+http://code.google.com/appengine)", "", -1))
		}
	}
	if v, ok := headers["Content-Type"]; ok {
		if v.(string) == "application/octet-stream" {
			delete(headers, "Content-Type")
		}
	}
	for key := range headers {
		match, _ := regexp.MatchString("(X-Google-|X-AppEngine-|Content-Length).*", key)
		if match {
			delete(headers, key)
		}
	}
	return headers
}

func flattenCloudSchedulerJobHttpTargetOauthToken(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service_account_email"] =
		flattenCloudSchedulerJobHttpTargetOauthTokenServiceAccountEmail(original["serviceAccountEmail"], d)
	transformed["scope"] =
		flattenCloudSchedulerJobHttpTargetOauthTokenScope(original["scope"], d)
	return []interface{}{transformed}
}
func flattenCloudSchedulerJobHttpTargetOauthTokenServiceAccountEmail(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobHttpTargetOauthTokenScope(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobHttpTargetOidcToken(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service_account_email"] =
		flattenCloudSchedulerJobHttpTargetOidcTokenServiceAccountEmail(original["serviceAccountEmail"], d)
	transformed["audience"] =
		flattenCloudSchedulerJobHttpTargetOidcTokenAudience(original["audience"], d)
	return []interface{}{transformed}
}
func flattenCloudSchedulerJobHttpTargetOidcTokenServiceAccountEmail(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudSchedulerJobHttpTargetOidcTokenAudience(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandCloudSchedulerJobName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	var jobName string
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("name"); ok {
		jobName = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", project, region, v.(string))
	} else {
		err := fmt.Errorf("The name is missing for the job cannot be empty")
		return nil, err
	}

	return jobName, nil
}

func expandCloudSchedulerJobDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobSchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobTimeZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRetryCount, err := expandCloudSchedulerJobRetryConfigRetryCount(original["retry_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetryCount); val.IsValid() && !isEmptyValue(val) {
		transformed["retryCount"] = transformedRetryCount
	}

	transformedMaxRetryDuration, err := expandCloudSchedulerJobRetryConfigMaxRetryDuration(original["max_retry_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetryDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["maxRetryDuration"] = transformedMaxRetryDuration
	}

	transformedMinBackoffDuration, err := expandCloudSchedulerJobRetryConfigMinBackoffDuration(original["min_backoff_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinBackoffDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["minBackoffDuration"] = transformedMinBackoffDuration
	}

	transformedMaxBackoffDuration, err := expandCloudSchedulerJobRetryConfigMaxBackoffDuration(original["max_backoff_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBackoffDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["maxBackoffDuration"] = transformedMaxBackoffDuration
	}

	transformedMaxDoublings, err := expandCloudSchedulerJobRetryConfigMaxDoublings(original["max_doublings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDoublings); val.IsValid() && !isEmptyValue(val) {
		transformed["maxDoublings"] = transformedMaxDoublings
	}

	return transformed, nil
}

func expandCloudSchedulerJobRetryConfigRetryCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxRetryDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMinBackoffDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxBackoffDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxDoublings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopicName, err := expandCloudSchedulerJobPubsubTargetTopicName(original["topic_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopicName); val.IsValid() && !isEmptyValue(val) {
		transformed["topicName"] = transformedTopicName
	}

	transformedData, err := expandCloudSchedulerJobPubsubTargetData(original["data"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedData); val.IsValid() && !isEmptyValue(val) {
		transformed["data"] = transformedData
	}

	transformedAttributes, err := expandCloudSchedulerJobPubsubTargetAttributes(original["attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributes); val.IsValid() && !isEmptyValue(val) {
		transformed["attributes"] = transformedAttributes
	}

	return transformed, nil
}

func expandCloudSchedulerJobPubsubTargetTopicName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTargetData(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTargetAttributes(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobAppEngineHttpTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHttpMethod, err := expandCloudSchedulerJobAppEngineHttpTargetHttpMethod(original["http_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpMethod); val.IsValid() && !isEmptyValue(val) {
		transformed["httpMethod"] = transformedHttpMethod
	}

	transformedAppEngineRouting, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(original["app_engine_routing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAppEngineRouting); val.IsValid() && !isEmptyValue(val) {
		transformed["appEngineRouting"] = transformedAppEngineRouting
	}

	transformedRelativeUri, err := expandCloudSchedulerJobAppEngineHttpTargetRelativeUri(original["relative_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRelativeUri); val.IsValid() && !isEmptyValue(val) {
		transformed["relativeUri"] = transformedRelativeUri
	}

	transformedBody, err := expandCloudSchedulerJobAppEngineHttpTargetBody(original["body"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBody); val.IsValid() && !isEmptyValue(val) {
		transformed["body"] = transformedBody
	}

	transformedHeaders, err := expandCloudSchedulerJobAppEngineHttpTargetHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	return transformed, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetHttpMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedInstance, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !isEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	return transformed, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetRelativeUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetBody(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobHttpTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandCloudSchedulerJobHttpTargetUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !isEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedHttpMethod, err := expandCloudSchedulerJobHttpTargetHttpMethod(original["http_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpMethod); val.IsValid() && !isEmptyValue(val) {
		transformed["httpMethod"] = transformedHttpMethod
	}

	transformedBody, err := expandCloudSchedulerJobHttpTargetBody(original["body"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBody); val.IsValid() && !isEmptyValue(val) {
		transformed["body"] = transformedBody
	}

	transformedHeaders, err := expandCloudSchedulerJobHttpTargetHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	transformedOauthToken, err := expandCloudSchedulerJobHttpTargetOauthToken(original["oauth_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOauthToken); val.IsValid() && !isEmptyValue(val) {
		transformed["oauthToken"] = transformedOauthToken
	}

	transformedOidcToken, err := expandCloudSchedulerJobHttpTargetOidcToken(original["oidc_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOidcToken); val.IsValid() && !isEmptyValue(val) {
		transformed["oidcToken"] = transformedOidcToken
	}

	return transformed, nil
}

func expandCloudSchedulerJobHttpTargetUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetHttpMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetBody(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobHttpTargetOauthToken(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccountEmail, err := expandCloudSchedulerJobHttpTargetOauthTokenServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedScope, err := expandCloudSchedulerJobHttpTargetOauthTokenScope(original["scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScope); val.IsValid() && !isEmptyValue(val) {
		transformed["scope"] = transformedScope
	}

	return transformed, nil
}

func expandCloudSchedulerJobHttpTargetOauthTokenServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetOauthTokenScope(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetOidcToken(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccountEmail, err := expandCloudSchedulerJobHttpTargetOidcTokenServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedAudience, err := expandCloudSchedulerJobHttpTargetOidcTokenAudience(original["audience"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudience); val.IsValid() && !isEmptyValue(val) {
		transformed["audience"] = transformedAudience
	}

	return transformed, nil
}

func expandCloudSchedulerJobHttpTargetOidcTokenServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetOidcTokenAudience(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
