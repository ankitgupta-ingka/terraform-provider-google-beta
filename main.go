package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-google-beta/google"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
