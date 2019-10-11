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
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/appengine/v1"
)

func resourceAppEngineStandardAppVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineStandardAppVersionCreate,
		Read:   resourceAppEngineStandardAppVersionRead,
		Update: resourceAppEngineStandardAppVersionUpdate,
		Delete: resourceAppEngineStandardAppVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineStandardAppVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"runtime": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"files": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"sha1_sum": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"source_url": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"zip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"files_count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"source_url": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"entrypoint": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"env_variables": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"handlers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_fail_action": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"AUTH_FAIL_ACTION_UNSPECIFIED", "AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED", ""}, false),
						},
						"login": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"LOGIN_UNSPECIFIED", "LOGIN_OPTIONAL", "LOGIN_ADMIN", "LOGIN_REQUIRED", ""}, false),
						},
						"redirect_http_response_code": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED", "REDIRECT_HTTP_RESPONSE_CODE_301", "REDIRECT_HTTP_RESPONSE_CODE_302", "REDIRECT_HTTP_RESPONSE_CODE_303", "REDIRECT_HTTP_RESPONSE_CODE_307", ""}, false),
						},
						"script": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"script_path": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"security_level": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"SECURE_UNSPECIFIED", "SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS", ""}, false),
						},
						"static_files": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_readable": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"expiration": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"http_headers": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"mime_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"require_matching_file": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"upload_path_regex": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"url_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"libraries": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"runtime_api_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"threadsafe": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"version_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"noop_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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

func resourceAppEngineStandardAppVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	idProp, err := expandAppEngineStandardAppVersionVersionId(d.Get("version_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version_id"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	runtimeProp, err := expandAppEngineStandardAppVersionRuntime(d.Get("runtime"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime"); !isEmptyValue(reflect.ValueOf(runtimeProp)) && (ok || !reflect.DeepEqual(v, runtimeProp)) {
		obj["runtime"] = runtimeProp
	}
	threadsafeProp, err := expandAppEngineStandardAppVersionThreadsafe(d.Get("threadsafe"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threadsafe"); !isEmptyValue(reflect.ValueOf(threadsafeProp)) && (ok || !reflect.DeepEqual(v, threadsafeProp)) {
		obj["threadsafe"] = threadsafeProp
	}
	runtimeApiVersionProp, err := expandAppEngineStandardAppVersionRuntimeApiVersion(d.Get("runtime_api_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_api_version"); !isEmptyValue(reflect.ValueOf(runtimeApiVersionProp)) && (ok || !reflect.DeepEqual(v, runtimeApiVersionProp)) {
		obj["runtimeApiVersion"] = runtimeApiVersionProp
	}
	handlersProp, err := expandAppEngineStandardAppVersionHandlers(d.Get("handlers"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("handlers"); !isEmptyValue(reflect.ValueOf(handlersProp)) && (ok || !reflect.DeepEqual(v, handlersProp)) {
		obj["handlers"] = handlersProp
	}
	librariesProp, err := expandAppEngineStandardAppVersionLibraries(d.Get("libraries"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("libraries"); !isEmptyValue(reflect.ValueOf(librariesProp)) && (ok || !reflect.DeepEqual(v, librariesProp)) {
		obj["libraries"] = librariesProp
	}
	envVariablesProp, err := expandAppEngineStandardAppVersionEnvVariables(d.Get("env_variables"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("env_variables"); !isEmptyValue(reflect.ValueOf(envVariablesProp)) && (ok || !reflect.DeepEqual(v, envVariablesProp)) {
		obj["envVariables"] = envVariablesProp
	}
	deploymentProp, err := expandAppEngineStandardAppVersionDeployment(d.Get("deployment"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("deployment"); !isEmptyValue(reflect.ValueOf(deploymentProp)) && (ok || !reflect.DeepEqual(v, deploymentProp)) {
		obj["deployment"] = deploymentProp
	}
	entrypointProp, err := expandAppEngineStandardAppVersionEntrypoint(d.Get("entrypoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entrypoint"); !isEmptyValue(reflect.ValueOf(entrypointProp)) && (ok || !reflect.DeepEqual(v, entrypointProp)) {
		obj["entrypoint"] = entrypointProp
	}

	lockName, err := replaceVars(d, config, "apps/{{project}}/services/{{service}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}/versions")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new StandardAppVersion: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating StandardAppVersion: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "apps/{{project}}/services/{{service}}/versions/{{version_id}}")
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
		config.clientAppEngine, op, project, "Creating StandardAppVersion",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create StandardAppVersion: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating StandardAppVersion %q: %#v", d.Id(), res)

	return resourceAppEngineStandardAppVersionRead(d, meta)
}

func resourceAppEngineStandardAppVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}/versions/{{version_id}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AppEngineStandardAppVersion %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOk("noop_on_destroy"); !ok {
		d.Set("noop_on_destroy", false)
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading StandardAppVersion: %s", err)
	}

	if err := d.Set("name", flattenAppEngineStandardAppVersionName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading StandardAppVersion: %s", err)
	}
	if err := d.Set("version_id", flattenAppEngineStandardAppVersionVersionId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading StandardAppVersion: %s", err)
	}
	if err := d.Set("runtime", flattenAppEngineStandardAppVersionRuntime(res["runtime"], d)); err != nil {
		return fmt.Errorf("Error reading StandardAppVersion: %s", err)
	}
	if err := d.Set("runtime_api_version", flattenAppEngineStandardAppVersionRuntimeApiVersion(res["runtimeApiVersion"], d)); err != nil {
		return fmt.Errorf("Error reading StandardAppVersion: %s", err)
	}
	if err := d.Set("handlers", flattenAppEngineStandardAppVersionHandlers(res["handlers"], d)); err != nil {
		return fmt.Errorf("Error reading StandardAppVersion: %s", err)
	}
	if err := d.Set("libraries", flattenAppEngineStandardAppVersionLibraries(res["libraries"], d)); err != nil {
		return fmt.Errorf("Error reading StandardAppVersion: %s", err)
	}

	return nil
}

func resourceAppEngineStandardAppVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	idProp, err := expandAppEngineStandardAppVersionVersionId(d.Get("version_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	runtimeProp, err := expandAppEngineStandardAppVersionRuntime(d.Get("runtime"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, runtimeProp)) {
		obj["runtime"] = runtimeProp
	}
	threadsafeProp, err := expandAppEngineStandardAppVersionThreadsafe(d.Get("threadsafe"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threadsafe"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, threadsafeProp)) {
		obj["threadsafe"] = threadsafeProp
	}
	runtimeApiVersionProp, err := expandAppEngineStandardAppVersionRuntimeApiVersion(d.Get("runtime_api_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_api_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, runtimeApiVersionProp)) {
		obj["runtimeApiVersion"] = runtimeApiVersionProp
	}
	handlersProp, err := expandAppEngineStandardAppVersionHandlers(d.Get("handlers"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("handlers"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, handlersProp)) {
		obj["handlers"] = handlersProp
	}
	librariesProp, err := expandAppEngineStandardAppVersionLibraries(d.Get("libraries"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("libraries"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, librariesProp)) {
		obj["libraries"] = librariesProp
	}
	envVariablesProp, err := expandAppEngineStandardAppVersionEnvVariables(d.Get("env_variables"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("env_variables"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, envVariablesProp)) {
		obj["envVariables"] = envVariablesProp
	}
	deploymentProp, err := expandAppEngineStandardAppVersionDeployment(d.Get("deployment"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("deployment"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, deploymentProp)) {
		obj["deployment"] = deploymentProp
	}
	entrypointProp, err := expandAppEngineStandardAppVersionEntrypoint(d.Get("entrypoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entrypoint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entrypointProp)) {
		obj["entrypoint"] = entrypointProp
	}

	lockName, err := replaceVars(d, config, "apps/{{project}}/services/{{service}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}/versions")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating StandardAppVersion %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating StandardAppVersion %q: %s", d.Id(), err)
	}

	op := &appengine.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = appEngineOperationWaitTime(
		config.clientAppEngine, op, project, "Updating StandardAppVersion",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceAppEngineStandardAppVersionRead(d, meta)
}

func resourceAppEngineStandardAppVersionDelete(d *schema.ResourceData, meta interface{}) error {
	if d.Get("noop_on_destroy") == true {
		log.Printf("[DEBUG] Keeping the StandardAppVersion %q", d.Id())
		return nil
	}
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "apps/{{project}}/services/{{service}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}/versions/{{version_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting StandardAppVersion %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "StandardAppVersion")
	}

	op := &appengine.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = appEngineOperationWaitTime(
		config.clientAppEngine, op, project, "Deleting StandardAppVersion",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting StandardAppVersion %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineStandardAppVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"apps/(?P<project>[^/]+)/services/(?P<service>[^/]+)/versions/(?P<version_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<service>[^/]+)/(?P<version_id>[^/]+)",
		"(?P<service>[^/]+)/(?P<version_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "apps/{{project}}/services/{{service}}/versions/{{version_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	d.Set("noop_on_destroy", false)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineStandardAppVersionName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionVersionId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionRuntime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionRuntimeApiVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlers(v interface{}, d *schema.ResourceData) interface{} {
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
			"url_regex":                   flattenAppEngineStandardAppVersionHandlersUrlRegex(original["urlRegex"], d),
			"security_level":              flattenAppEngineStandardAppVersionHandlersSecurityLevel(original["securityLevel"], d),
			"login":                       flattenAppEngineStandardAppVersionHandlersLogin(original["login"], d),
			"auth_fail_action":            flattenAppEngineStandardAppVersionHandlersAuthFailAction(original["authFailAction"], d),
			"redirect_http_response_code": flattenAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(original["redirectHttpResponseCode"], d),
			"script":                      flattenAppEngineStandardAppVersionHandlersScript(original["script"], d),
			"static_files":                flattenAppEngineStandardAppVersionHandlersStaticFiles(original["staticFiles"], d),
		})
	}
	return transformed
}
func flattenAppEngineStandardAppVersionHandlersUrlRegex(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersSecurityLevel(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersLogin(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersAuthFailAction(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersScript(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["script_path"] =
		flattenAppEngineStandardAppVersionHandlersScriptScriptPath(original["scriptPath"], d)
	return []interface{}{transformed}
}
func flattenAppEngineStandardAppVersionHandlersScriptScriptPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersStaticFiles(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["path"] =
		flattenAppEngineStandardAppVersionHandlersStaticFilesPath(original["path"], d)
	transformed["upload_path_regex"] =
		flattenAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(original["uploadPathRegex"], d)
	transformed["http_headers"] =
		flattenAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(original["httpHeaders"], d)
	transformed["mime_type"] =
		flattenAppEngineStandardAppVersionHandlersStaticFilesMimeType(original["mimeType"], d)
	transformed["expiration"] =
		flattenAppEngineStandardAppVersionHandlersStaticFilesExpiration(original["expiration"], d)
	transformed["require_matching_file"] =
		flattenAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(original["requireMatchingFile"], d)
	transformed["application_readable"] =
		flattenAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(original["applicationReadable"], d)
	return []interface{}{transformed}
}
func flattenAppEngineStandardAppVersionHandlersStaticFilesPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersStaticFilesMimeType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersStaticFilesExpiration(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionLibraries(v interface{}, d *schema.ResourceData) interface{} {
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
			"name":    flattenAppEngineStandardAppVersionLibrariesName(original["name"], d),
			"version": flattenAppEngineStandardAppVersionLibrariesVersion(original["version"], d),
		})
	}
	return transformed
}
func flattenAppEngineStandardAppVersionLibrariesName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineStandardAppVersionLibrariesVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandAppEngineStandardAppVersionVersionId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionRuntime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionThreadsafe(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionRuntimeApiVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUrlRegex, err := expandAppEngineStandardAppVersionHandlersUrlRegex(original["url_regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUrlRegex); val.IsValid() && !isEmptyValue(val) {
			transformed["urlRegex"] = transformedUrlRegex
		}

		transformedSecurityLevel, err := expandAppEngineStandardAppVersionHandlersSecurityLevel(original["security_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecurityLevel); val.IsValid() && !isEmptyValue(val) {
			transformed["securityLevel"] = transformedSecurityLevel
		}

		transformedLogin, err := expandAppEngineStandardAppVersionHandlersLogin(original["login"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLogin); val.IsValid() && !isEmptyValue(val) {
			transformed["login"] = transformedLogin
		}

		transformedAuthFailAction, err := expandAppEngineStandardAppVersionHandlersAuthFailAction(original["auth_fail_action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAuthFailAction); val.IsValid() && !isEmptyValue(val) {
			transformed["authFailAction"] = transformedAuthFailAction
		}

		transformedRedirectHttpResponseCode, err := expandAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(original["redirect_http_response_code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRedirectHttpResponseCode); val.IsValid() && !isEmptyValue(val) {
			transformed["redirectHttpResponseCode"] = transformedRedirectHttpResponseCode
		}

		transformedScript, err := expandAppEngineStandardAppVersionHandlersScript(original["script"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedScript); val.IsValid() && !isEmptyValue(val) {
			transformed["script"] = transformedScript
		}

		transformedStaticFiles, err := expandAppEngineStandardAppVersionHandlersStaticFiles(original["static_files"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStaticFiles); val.IsValid() && !isEmptyValue(val) {
			transformed["staticFiles"] = transformedStaticFiles
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineStandardAppVersionHandlersUrlRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersSecurityLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersLogin(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersAuthFailAction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersScript(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedScriptPath, err := expandAppEngineStandardAppVersionHandlersScriptScriptPath(original["script_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScriptPath); val.IsValid() && !isEmptyValue(val) {
		transformed["scriptPath"] = transformedScriptPath
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionHandlersScriptScriptPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandAppEngineStandardAppVersionHandlersStaticFilesPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedUploadPathRegex, err := expandAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(original["upload_path_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUploadPathRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["uploadPathRegex"] = transformedUploadPathRegex
	}

	transformedHttpHeaders, err := expandAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(original["http_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["httpHeaders"] = transformedHttpHeaders
	}

	transformedMimeType, err := expandAppEngineStandardAppVersionHandlersStaticFilesMimeType(original["mime_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMimeType); val.IsValid() && !isEmptyValue(val) {
		transformed["mimeType"] = transformedMimeType
	}

	transformedExpiration, err := expandAppEngineStandardAppVersionHandlersStaticFilesExpiration(original["expiration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpiration); val.IsValid() && !isEmptyValue(val) {
		transformed["expiration"] = transformedExpiration
	}

	transformedRequireMatchingFile, err := expandAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(original["require_matching_file"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireMatchingFile); val.IsValid() && !isEmptyValue(val) {
		transformed["requireMatchingFile"] = transformedRequireMatchingFile
	}

	transformedApplicationReadable, err := expandAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(original["application_readable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApplicationReadable); val.IsValid() && !isEmptyValue(val) {
		transformed["applicationReadable"] = transformedApplicationReadable
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesMimeType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesExpiration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionLibraries(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandAppEngineStandardAppVersionLibrariesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedVersion, err := expandAppEngineStandardAppVersionLibrariesVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineStandardAppVersionLibrariesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionLibrariesVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionEnvVariables(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAppEngineStandardAppVersionDeployment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedZip, err := expandAppEngineStandardAppVersionDeploymentZip(original["zip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedZip); val.IsValid() && !isEmptyValue(val) {
		transformed["zip"] = transformedZip
	}

	transformedFiles, err := expandAppEngineStandardAppVersionDeploymentFiles(original["files"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFiles); val.IsValid() && !isEmptyValue(val) {
		transformed["files"] = transformedFiles
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionDeploymentZip(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSourceUrl, err := expandAppEngineStandardAppVersionDeploymentZipSourceUrl(original["source_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSourceUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["sourceUrl"] = transformedSourceUrl
	}

	transformedFilesCount, err := expandAppEngineStandardAppVersionDeploymentZipFilesCount(original["files_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilesCount); val.IsValid() && !isEmptyValue(val) {
		transformed["filesCount"] = transformedFilesCount
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionDeploymentZipSourceUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentZipFilesCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentFiles(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSha1Sum, err := expandAppEngineStandardAppVersionDeploymentFilesSha1Sum(original["sha1_sum"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["sha1Sum"] = transformedSha1Sum
		transformedSourceUrl, err := expandAppEngineStandardAppVersionDeploymentFilesSourceUrl(original["source_url"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["sourceUrl"] = transformedSourceUrl

		m[original["name"].(string)] = transformed
	}
	return m, nil
}

func expandAppEngineStandardAppVersionDeploymentFilesSha1Sum(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentFilesSourceUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionEntrypoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedShell, err := expandAppEngineStandardAppVersionEntrypointShell(original["shell"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShell); val.IsValid() && !isEmptyValue(val) {
		transformed["shell"] = transformedShell
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionEntrypointShell(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
