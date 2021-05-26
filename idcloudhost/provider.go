package idcloudhost

import (
	"context"

	"github.com/bapung/idcloudhost-go-client-library/idcloudhost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("IDCLOUDHOST_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"idcloudhost_vm": resourceVirtualMachine(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"idcloudhost_vms": dataSourceVirtualMachine(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	authToken := d.Get("auth_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if authToken != "" {
		c, err := idcloudhost.NewClient(authToken, "jkt01")
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}
	c, err := idcloudhost.NewClient("", "")
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
