package idcloudhost

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bapung/idcloudhost-go-client-library/idcloudhost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDisk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDiskCreate,
		ReadContext:   resourceDiskRead,
		UpdateContext: resourceDiskUpdate,
		DeleteContext: resourceDiskDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
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
				Required: true,
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
			"vm_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*idcloudhost.APIClient)
	var diags diag.Diagnostics

	vmUUID := d.Get("vm_uuid").(string)
	diskSize := d.Get("size").(int)

	diskApi := c.Disk
	diskApi.Bind(vmUUID)
	err := diskApi.Create(diskSize)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create new Disk",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}

	diskId := fmt.Sprintf("%s/%s", vmUUID, diskApi.Disk.UUID)
	d.SetId(diskId)

	resourceDiskRead(ctx, d, m)

	return diags
}

func resourceDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*idcloudhost.APIClient)
	var diags diag.Diagnostics
	ids := strings.Split(d.Id(), "/")
	vmUUID := ids[0]
	diskUUID := ids[1]
	diskApi := c.Disk
	diskApi.Bind(vmUUID)
	err := diskApi.Get(diskUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Disk",
			Detail:   "",
		})
		return diags
	}
	// assign state using d.Set() here

	return diags
}

func resourceDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*idcloudhost.APIClient)
	diskUUID := d.Id()
	vmUUID := d.Get("vm_uuid").(string)
	diskApi := c.Disk
	diskApi.Bind(vmUUID)
	err := diskApi.Delete(diskUUID)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
