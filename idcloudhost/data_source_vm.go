package idcloudhost

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/bapung/idcloudhost-go-client-library/idcloudhost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*idcloudhost.APIClient)
	var diags diag.Diagnostics

	vmApi := c.APIs["vm"].(*idcloudhost.VirtualMachineAPI)
	if err := vmApi.ListAll(); err != nil {
		log.Fatal(err)
	}

	if err := d.Set("vms", &vmApi.VMListMap); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func dataSourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualMachineRead,
		Schema: map[string]*schema.Schema{
			"vms": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"billing_account": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mac": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ipv4": {
							Type:     schema.TypeString,
							Computed: true,
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
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vcpu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
