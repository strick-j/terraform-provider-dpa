package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/strick-j/terraform-provider-dpa/dpa"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dpa.Provider,
	})
}
