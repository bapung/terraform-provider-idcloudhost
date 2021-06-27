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
			"auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("IDCLOUDHOST_AUTH_TOKEN", nil),
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "jkt01",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"idcloudhost_vm":       resourceVirtualMachine(),
			"idcloudhost_vm_disks": resourceDisk(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"idcloudhost_vms": dataSourceVirtualMachine(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	authToken := d.Get("auth_token").(string)
	region := d.Get("region").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if authToken != "" {
		c, err := idcloudhost.NewClient(authToken, region)
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
