package google

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSecretmanagerSecretVersion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecretmanagerSecretVersionRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destroy_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_data": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceSecretmanagerSecretVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fv, err := parseProjectFieldValue("secrets", d.Get("secret").(string), "project", d, config, false)
	if err != nil {
		return err
	}
	project := fv.Project
	d.Set("secret", fv.Name)

	var url string
	versionNum := d.Get("version")

	if versionNum != "" {
		url, err = replaceVars(d, config, "{{SecretmanagerBasePath}}projects/{{project}}/secrets/{{secret}}/versions/{{version}}")
		if err != nil {
			return err
		}
	} else {
		url, err = replaceVars(d, config, "{{SecretmanagerBasePath}}projects/{{project}}/secrets/{{secret}}/versions/latest")
		if err != nil {
			return err
		}
	}

	var version map[string]interface{}
	version, err = sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return fmt.Errorf("Error retrieving available secret manager secret versions: %s", err.Error())
	}

	secretVersionRegex := regexp.MustCompile("projects/(.+)/secrets/(.+)/versions/(.+)$")

	parts := secretVersionRegex.FindStringSubmatch(version["name"].(string))
	// should return [full string, project number, secret name, version number]
	if len(parts) != 4 {
		return fmt.Errorf("secret name, %s, does not match format, projects/{{project}}/secrets/{{secret}}/versions/{{versions}}", version["name"].(string))
	}

	log.Printf("[DEBUG] Received Google Secretmanager Version: %q", version)

	// Get the payload
	url = fmt.Sprintf("%s:access", url)
	resp, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return fmt.Errorf("Error retrieving available secret manager secret version access: %s", err.Error())
	}

	data := resp["payload"].(map[string]interface{})

	d.Set("project", project)
	d.Set("version", parts[3])
	d.Set("create_time", version["createTime"].(string))
	if version["destroyTime"] != nil {
		d.Set("destroy_time", version["destroyTime"].(string))
	}
	d.Set("name", version["name"].(string))
	d.Set("state", version["state"].(string))
	d.Set("secret_data", data["data"].(string))

	d.SetId(time.Now().UTC().String())
	return nil
}
