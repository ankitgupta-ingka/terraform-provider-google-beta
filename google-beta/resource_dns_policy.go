// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//     ***     DIFF TEST DIFF TEST    ***
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
	"bytes"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDNSPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSPolicyCreate,
		Read:   resourceDNSPolicyRead,
		Update: resourceDNSPolicyUpdate,
		Delete: resourceDNSPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDNSPolicyImport,
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
				Description: `User assigned name for this policy.`,
			},
			"alternative_name_server_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Sets an alternative name server for the associated networks.
When specified, all DNS queries are forwarded to a name server that you choose.
Names such as .internal are not available when an alternative name server is specified.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_name_servers": {
							Type:     schema.TypeSet,
							Required: true,
							Description: `Sets an alternative name server for the associated networks. When specified,
all DNS queries are forwarded to a name server that you choose. Names such as .internal
are not available when an alternative name server is specified.`,
							Elem: dnsPolicyAlternativeNameServerConfigTargetNameServersSchema(),
							Set: func(v interface{}) int {
								raw := v.(map[string]interface{})
								if address, ok := raw["ipv4_address"]; ok {
									hashcode.String(address.(string))
								}
								var buf bytes.Buffer
								schema.SerializeResourceForHash(&buf, raw, dnsPolicyAlternativeNameServerConfigTargetNameServersSchema())
								return hashcode.String(buf.String())
							},
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A textual description field. Defaults to 'Managed by Terraform'.`,
				Default:     "Managed by Terraform",
			},
			"enable_inbound_forwarding": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Allows networks bound to this policy to receive DNS queries sent
by VMs or applications over VPN connections. When enabled, a
virtual IP address will be allocated from each of the sub-networks
that are bound to this policy.`,
			},
			"enable_logging": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Controls whether logging is enabled for the networks bound to this policy.
Defaults to no logging if not set.`,
			},
			"networks": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `List of network names specifying networks to which this policy is applied.`,
				Elem:        dnsPolicyNetworksSchema(),
				Set: func(v interface{}) int {
					raw := v.(map[string]interface{})
					if url, ok := raw["network_url"]; ok {
						return selfLinkNameHash(url)
					}
					var buf bytes.Buffer
					schema.SerializeResourceForHash(&buf, raw, dnsPolicyNetworksSchema())
					return hashcode.String(buf.String())
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

func dnsPolicyAlternativeNameServerConfigTargetNameServersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ipv4_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `IPv4 address to forward to.`,
			},
		},
	}
}

func dnsPolicyNetworksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_url": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The fully qualified URL of the VPC network to bind to.
This should be formatted like
'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'`,
			},
		},
	}
}

func resourceDNSPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	alternativeNameServerConfigProp, err := expandDNSPolicyAlternativeNameServerConfig(d.Get("alternative_name_server_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("alternative_name_server_config"); !isEmptyValue(reflect.ValueOf(alternativeNameServerConfigProp)) && (ok || !reflect.DeepEqual(v, alternativeNameServerConfigProp)) {
		obj["alternativeNameServerConfig"] = alternativeNameServerConfigProp
	}
	descriptionProp, err := expandDNSPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableInboundForwardingProp, err := expandDNSPolicyEnableInboundForwarding(d.Get("enable_inbound_forwarding"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_inbound_forwarding"); ok || !reflect.DeepEqual(v, enableInboundForwardingProp) {
		obj["enableInboundForwarding"] = enableInboundForwardingProp
	}
	enableLoggingProp, err := expandDNSPolicyEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
		obj["enableLogging"] = enableLoggingProp
	}
	nameProp, err := expandDNSPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networksProp, err := expandDNSPolicyNetworks(d.Get("networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("networks"); !isEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Policy: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Policy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/policies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Policy %q: %#v", d.Id(), res)

	return resourceDNSPolicyRead(d, meta)
}

func resourceDNSPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DNSPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	if err := d.Set("alternative_name_server_config", flattenDNSPolicyAlternativeNameServerConfig(res["alternativeNameServerConfig"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("description", flattenDNSPolicyDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("enable_inbound_forwarding", flattenDNSPolicyEnableInboundForwarding(res["enableInboundForwarding"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("enable_logging", flattenDNSPolicyEnableLogging(res["enableLogging"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("name", flattenDNSPolicyName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("networks", flattenDNSPolicyNetworks(res["networks"], d)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	return nil
}

func resourceDNSPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	d.Partial(true)

	if d.HasChange("alternative_name_server_config") || d.HasChange("description") || d.HasChange("enable_inbound_forwarding") || d.HasChange("enable_logging") || d.HasChange("networks") {
		obj := make(map[string]interface{})

		alternativeNameServerConfigProp, err := expandDNSPolicyAlternativeNameServerConfig(d.Get("alternative_name_server_config"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("alternative_name_server_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, alternativeNameServerConfigProp)) {
			obj["alternativeNameServerConfig"] = alternativeNameServerConfigProp
		}
		descriptionProp, err := expandDNSPolicyDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
			obj["description"] = descriptionProp
		}
		enableInboundForwardingProp, err := expandDNSPolicyEnableInboundForwarding(d.Get("enable_inbound_forwarding"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("enable_inbound_forwarding"); ok || !reflect.DeepEqual(v, enableInboundForwardingProp) {
			obj["enableInboundForwarding"] = enableInboundForwardingProp
		}
		enableLoggingProp, err := expandDNSPolicyEnableLogging(d.Get("enable_logging"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
			obj["enableLogging"] = enableLoggingProp
		}
		networksProp, err := expandDNSPolicyNetworks(d.Get("networks"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("networks"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networksProp)) {
			obj["networks"] = networksProp
		}

		url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
		if err != nil {
			return err
		}
		_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
		}

		d.SetPartial("alternative_name_server_config")
		d.SetPartial("description")
		d.SetPartial("enable_inbound_forwarding")
		d.SetPartial("enable_logging")
		d.SetPartial("networks")
	}

	d.Partial(false)

	return resourceDNSPolicyRead(d, meta)
}

func resourceDNSPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	// if networks are attached, they need to be detached before the policy can be deleted
	if d.Get("networks.#").(int) > 0 {
		patched := make(map[string]interface{})
		patched["networks"] = nil

		url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequestWithTimeout(config, "PATCH", project, url, patched, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
		}
	}
	log.Printf("[DEBUG] Deleting Policy %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Policy")
	}

	log.Printf("[DEBUG] Finished deleting Policy %q: %#v", d.Id(), res)
	return nil
}

func resourceDNSPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/policies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/policies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDNSPolicyAlternativeNameServerConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_name_servers"] =
		flattenDNSPolicyAlternativeNameServerConfigTargetNameServers(original["targetNameServers"], d)
	return []interface{}{transformed}
}
func flattenDNSPolicyAlternativeNameServerConfigTargetNameServers(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(func(v interface{}) int {
		raw := v.(map[string]interface{})
		if address, ok := raw["ipv4_address"]; ok {
			hashcode.String(address.(string))
		}
		var buf bytes.Buffer
		schema.SerializeResourceForHash(&buf, raw, dnsPolicyAlternativeNameServerConfigTargetNameServersSchema())
		return hashcode.String(buf.String())
	}, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"ipv4_address": flattenDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(original["ipv4Address"], d),
		})
	}
	return transformed
}
func flattenDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSPolicyDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSPolicyEnableInboundForwarding(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSPolicyEnableLogging(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSPolicyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDNSPolicyNetworks(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(func(v interface{}) int {
		raw := v.(map[string]interface{})
		if url, ok := raw["network_url"]; ok {
			return selfLinkNameHash(url)
		}
		var buf bytes.Buffer
		schema.SerializeResourceForHash(&buf, raw, dnsPolicyNetworksSchema())
		return hashcode.String(buf.String())
	}, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"network_url": flattenDNSPolicyNetworksNetworkUrl(original["networkUrl"], d),
		})
	}
	return transformed
}
func flattenDNSPolicyNetworksNetworkUrl(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandDNSPolicyAlternativeNameServerConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetNameServers, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServers(original["target_name_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNameServers); val.IsValid() && !isEmptyValue(val) {
		transformed["targetNameServers"] = transformedTargetNameServers
	}

	return transformed, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpv4Address, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(original["ipv4_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpv4Address); val.IsValid() && !isEmptyValue(val) {
			transformed["ipv4Address"] = transformedIpv4Address
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyEnableInboundForwarding(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyEnableLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSPolicyNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSPolicyNetworksNetworkUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
