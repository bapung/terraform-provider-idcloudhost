package idcloudhost

import (
	"context"
	"fmt"
	"strings"
	"time"

	idcloudhostAPI "github.com/bapung/idcloudhost-go-client-library/idcloudhost/api"
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
	c := m.(*idcloudhostAPI.APIClient)
	var diags diag.Diagnostics
	diskApi := c.Disk

	vmUUID := d.Get("vm_uuid").(string)
	diskSize := d.Get("size").(int)

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

	diskResourceId := fmt.Sprintf("%s/%s", vmUUID, diskApi.Disk.UUID)
	d.SetId(diskResourceId)

	resourceDiskRead(ctx, d, m)

	return diags
}

func resourceDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*idcloudhostAPI.APIClient)
	diskApi := c.Disk
	vmApi := c.VM

	diskResourceId := strings.Split(d.Id(), "/")
	vmUUID := diskResourceId[0]
	diskUUID := diskResourceId[1]
	diskApi.Bind(vmUUID)
	err := vmApi.Get(d.vmUUID)
	if err != nil {
		return err
	}
	err := diskApi.Get(diskUUID, &vmApi.VM.Storage)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Disk",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}
	err = setDiskResource(d, diskApi.Disk)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Disk",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}
	return diags
}

func resourceDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var newSize, oldSize int
	c := m.(*idcloudhostAPI.APIClient)
	diskApi := c.Disk

	diskResourceId := strings.Split(d.Id(), "/")
	vmUUID := diskResourceId[0]
	diskUUID := diskResourceId[1]

	if d.HasChange("vm_uuid") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update disk",
			Detail:   "Disk cannot be reassigned to other resources",
		})
		return diags
	}

	if d.HasChange("size") {
		oldSizeIface, newSizeIface := d.GetChange("size")
		newSize = newSizeIface.(int)
		oldSize = oldSizeIface.(int)
		isShrink := newSize-oldSize < 0
		if isShrink {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update disk",
				Detail:   "Disk cannot be resized, shrinking disk is not possible",
			})
			return diags
		}
	}
	diskApi.Bind(vmUUID)
	err := diskApi.Modify(diskUUID, newSize)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update Disk",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}
	err = setDiskResource(d, diskApi.Disk)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update Disk",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}
	return diags
}

func resourceDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*idcloudhostAPI.APIClient)

	diskApi := c.Disk
	diskResourceId := strings.Split(d.Id(), "/")
	vmUUID := diskResourceId[0]
	diskUUID := diskResourceId[1]

	diskApi.Bind(vmUUID)
	err := diskApi.Delete(diskUUID)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}