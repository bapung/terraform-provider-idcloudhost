package idcloudhost

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

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
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 20 || v > 240 {
						errs = append(errs, fmt.Errorf("%q must be between 20 and 240, got: %d", key, v))
					}
					return
				},
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
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 1024 || v > 65536 {
						errs = append(errs, fmt.Errorf("%q must be between 1024 and 66536, got: %d", key, v))
					}
					return
				},
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
							Required: true,
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
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 1 || v > 16 {
						errs = append(errs, fmt.Errorf("%q must be between 1 and 16, got: %d", key, v))
					}
					return
				},
			},
		},
	}
}

func resourceVirtualMachineCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*idcloudhost.APIClient)
	var diags diag.Diagnostics

	newVM := map[string]interface{}{
		"backup":            d.Get("backup"),
		"billing_account":   d.Get("billing_account"), //should be automatically assigned to "default" billing account if not specified
		"description":       d.Get("description"),
		"disks":             d.Get("disks"),
		"name":              d.Get("name"),
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
	c := m.(*idcloudhost.APIClient)
	vmApi := c.APIs["vm"].(*idcloudhost.VirtualMachineAPI)
	uuid := d.Id()
	propertyMap := map[string]interface{}{
		"uuid": uuid, "vcpu": nil, "ram": nil, "name": nil,
	}
	isSomethingChanged := false
	requireToFetchStatus := false
	for k, _ := range propertyMap {
		if d.HasChange(k) {
			isSomethingChanged = true
			if (k == "ram") || (k == "vcpu") {
				requireToFetchStatus = true
			}
		}
		propertyMap[k] = d.Get(k)
	}

	if requireToFetchStatus {
		if err := vmApi.Get(uuid); err != nil {
			if vmApi.VMMap["status"].(string) != "stopped" {
				err := errors.New("VM is not in stopped state, cannot update resource")
				return diag.FromErr(err)
			}
			err := errors.New("cannot fetch VM state for update, cannot update resource")
			return diag.FromErr(err)
		}
	}

	if isSomethingChanged {
		err := vmApi.Modify(propertyMap)
		if err != nil {
			return diag.FromErr(err)
		}
		for k, v := range vmApi.VMMap {
			err := d.Set(k, v)
			if err != nil {
				return diag.FromErr(err)
			}
		}
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceVirtualMachineRead(ctx, d, m)
}

func resourceVirtualMachineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
