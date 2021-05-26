package idcloudhost

import (
	"context"
	"log"

	"github.com/bapung/idcloudhost-go-client-library/idcloudhost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVirtualMachineCreate,
		ReadContext:   resourceVirtualMachineRead,
		UpdateContext: resourceVirtualMachineUpdate,
		DeleteContext: resourceVirtualMachineDelete,
		Schema: map[string]*schema.Schema{
			"backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"billing_account": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disks": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"initial_password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_ipv4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reserve_public_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"source_replica": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pool": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"replica": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"shared": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcpu": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceVirtualMachineCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*idcloudhost.APIClient)
	var diags diag.Diagnostics

	newVM := map[string]interface{}{
		"backup":            d.Get("backup"),
		"billing_account":   d.Get("billing_account"), //should be automatically handled if the auth token is valid, to do
		"description":       d.Get("description"),
		"disks":             d.Get("disks"),
		"os_name":           d.Get("os_name"),
		"os_version":        d.Get("os_version"),
		"password":          d.Get("initial_password"),
		"public_key":        d.Get("public_key"),
		"source_replica":    d.Get("source_replica"),
		"source_uuid":       d.Get("source_uuid"),
		"username":          d.Get("username"),
		"vcpu":              d.Get("vcpu"),
		"ram":               d.Get("memory"),
		"reserve_public_ip": d.Get("reserve_public_ip"),
	}

	vmApi := c.APIs["vm"].(*idcloudhost.VirtualMachineAPI)
	if err := vmApi.Create(newVM); err != nil {
		log.Fatal(err)
	}

	d.SetId(vmApi.VMMap["uuid"].(string))

	resourceVirtualMachineRead(ctx, d, m)

	return diags
}

func resourceVirtualMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*idcloudhost.APIClient)
	var diags diag.Diagnostics
	uuid := d.Id()
	vmApi := c.APIs["vm"].(*idcloudhost.VirtualMachineAPI)
	err := vmApi.Get(uuid)
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range vmApi.VMMap {
		err := d.Set(k, v)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVirtualMachineUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceVirtualMachineRead(ctx, d, m)
}

func resourceVirtualMachineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
