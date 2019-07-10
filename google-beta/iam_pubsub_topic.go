package google

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var PubsubTopicIamSchema = map[string]*schema.Schema{
	"project": {
		Type:             schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew:         true,
	},
	"topic": {
		Type:             schema.TypeString,
		Required: true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type PubsubTopicIamUpdater struct {
	project string
	topic string
	d       *schema.ResourceData
	Config  *Config
}

func PubsubTopicIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)
	
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	// While this may be overridden by the "project" value from getImportIdQualifiers below,
	// setting project here ensures the value is set even if the value set in config is the short
	// name or otherwise doesn't include the project.
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/topics/(?P<topic>[^/]+)","(?P<project>[^/]+)/(?P<topic>[^/]+)","(?P<topic>[^/]+)"}, d, config, d.Get("topic").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &PubsubTopicIamUpdater{
		project: values["project"],
		topic: values["topic"],
		d:       d,
		Config:  config,
	}
	d.SetId(u.GetResourceId())

	return u, nil
}

func PubsubTopicIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)
	
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/topics/(?P<topic>[^/]+)","(?P<project>[^/]+)/(?P<topic>[^/]+)","(?P<topic>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
    values[k] = v
	}

	u := &PubsubTopicIamUpdater{
		project: values["project"],
		topic: values["topic"],
		d:       d,
		Config:  config,
	}
	d.Set("topic", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *PubsubTopicIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url := u.qualifyTopicUrl("getIamPolicy")

	policy, err := sendRequest(u.Config, "GET", url, nil)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *PubsubTopicIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url := u.qualifyTopicUrl("setIamPolicy")
	
	_, err = sendRequestWithTimeout(u.Config, "POST", url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *<= resource_name -%>IamUpdater) qualifyTopicUrl(methodIdentifier string) string {
	return fmt.Sprintf("https://pubsub.googleapis.com/v1/%s:%s", fmt.Sprintf("projects/%s/topics/%s", u.project, u.topic), methodIdentifier)
}

func (u *PubsubTopicIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/topics/%s", u.project, u.topic)
}

func (u *PubsubTopicIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-pubsub-topic-%s", u.GetResourceId())
}

func (u *PubsubTopicIamUpdater) DescribeResource() string {
	return fmt.Sprintf("pubsub topic %q", u.GetResourceId())
}
