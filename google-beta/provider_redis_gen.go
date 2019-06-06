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

import "github.com/hashicorp/terraform/helper/schema"

var RedisDefaultBasePath = "https://redis.googleapis.com/v1beta1/"

var RedisBasePathEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_REDIS_BASE_PATH",
	}, RedisDefaultBasePath),
}

var GeneratedRedisResourcesMap = map[string]*schema.Resource{
	"google_redis_instance": resourceRedisInstance(),
}
