package idcloudhost

import (
	"context"
	"fmt"
	"time"

	idcloudhostAPI "github.com/bapung/idcloudhost-go-client-library/idcloudhost/api"
	idcloudhostVM "github.com/bapung/idcloudhost-go-client-library/idcloudhost/vm"
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"billing_account_id": {
				Type:     schema.TypeInt,
				Required: true,
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
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"private_ipv4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
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
	c := m.(*idcloudhostAPI.APIClient)
	var diags diag.Diagnostics

	newVM := &idcloudhostVM.NewVM{
		Backup:          d.Get("backup").(bool),
		BillingAccount:  d.Get("billing_account_id").(int), //should be automatically assigned to "default" billing account if not specified
		Description:     d.Get("description").(string),
		Disks:           d.Get("disks").(int),
		Name:            d.Get("name").(string),
		OSName:          d.Get("os_name").(string),
		OSVersion:       d.Get("os_version").(string),
		InitialPassword: d.Get("initial_password").(string),
		PublicKey:       d.Get("public_key").(string),
		SourceReplica:   d.Get("source_replica").(string),
		SourceUUID:      d.Get("source_uuid").(string),
		Username:        d.Get("username").(string),
		VCPU:            d.Get("vcpu").(int),
		Memory:          d.Get("memory").(int),
		ReservePublicIP: false,
	}

	vmApi := c.VM
	if err := vmApi.Create(*newVM); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create new VM",
			Detail:   fmt.Sprint(err),
		})

		return diags
	}

	d.SetId(vmApi.VM.UUID)

	resourceVirtualMachineRead(ctx, d, m)

	return diags
}

func resourceVirtualMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*idcloudhostAPI.APIClient)
	var diags diag.Diagnostics
	uuid := d.Id()
	vmApi := c.VM
	err := vmApi.Get(uuid)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get VM",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}

	err = setVmResource(d, &vmApi.VM)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get VM",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}

	return diags
}

func resourceVirtualMachineUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var isSomethingChanged = true
	c := m.(*idcloudhostAPI.APIClient)
	vmApi := c.VM
	uuid := d.Id()

	if d.HasChanges("name", "vcpu", "memory") {
		isSomethingChanged = true
		updatedVM := &idcloudhostVM.VM{
			UUID:   uuid,
			Name:   d.Get("name").(string),
			VCPU:   d.Get("vcpu").(int),
			Memory: d.Get("memory").(int),
		}
		err := vmApi.Get(uuid)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to modify VM",
				Detail:   "cannot fetch VM state for update, cannot update resource",
			})
			return diags
		}
		if d.HasChanges("vcpu", "memory") && vmApi.VM.Status != "stopped" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to modify VM",
				Detail:   "Updating vcpu and ram requires VM to be stopped",
			})
			return diags
		}
		err = vmApi.Modify(*updatedVM)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to modify VM",
				Detail:   fmt.Sprint(err),
			})
			return diags
		}
		err = setVmResource(d, &vmApi.VM)
		if err != nil {
			return diags
		}
	}

	if d.HasChange("backup") {
		isSomethingChanged = true
		err := vmApi.ToggleAutoBackup(uuid)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to modify VM",
				Detail:   "Unable to toggle auto backup",
			})
			return diags
		}
		err = setVmResource(d, &vmApi.VM)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to modify VM",
				Detail:   fmt.Sprint(err),
			})
			return diags
		}
	}
	if isSomethingChanged {
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceVirtualMachineRead(ctx, d, m)
}

func resourceVirtualMachineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*idcloudhostAPI.APIClient)
	uuid := d.Id()
	vmApi := c.VM
	err := vmApi.Delete(uuid)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
