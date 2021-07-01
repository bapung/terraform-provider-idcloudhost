package idcloudhost

import (
	"context"
	"fmt"
	"time"

	"github.com/bapung/idcloudhost-go-client-library/idcloudhost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFloatingIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFloatingIPCreate,
		ReadContext:   resourceFloatingIPRead,
		UpdateContext: resourceFloatingIPUpdate,
		DeleteContext: resourceFloatingIPDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			// ID             int    `json:"id,omitempty"`
			// Address        string `json:"address,omitempty"`
			// UserID         int    `json:"user_id,omitempty"`
			// BillingAccount int    `json:"billing_account_id"`
			// Type           string `json:"type,omitempty"`
			// NetworkID      string `json:"network_id,omitempty"`
			// Name           string `json:"name"`
			// Enabled        bool   `json:"enabled,omitempty"`
			// CreatedAt      string `json:"created_at,omitempty"`
			// UpdatedAt      string `json:"updated_at,omitempty"`
			// AssignedTo     string `json:"assigned_to,omitempty"`
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"billing_account_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assigned_to": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}
func resourceFloatingIPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*idcloudhost.APIClient)
	fipApi := c.FloatingIP

	billingAccountId := d.Get("billing_account_id").(int)
	name := d.Get("name").(string)
	assignedUuid := d.Get("assigned_to").(string)
	err := fipApi.Create(name, billingAccountId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{})
		return diags
	}
	d.SetId(fipApi.FloatingIP.Address)
	if assignedUuid != "" {
		err := fipApi.Assign(fipApi.FloatingIP.Address, assignedUuid)
		if err != nil {
			resourceFloatingIPDelete(ctx, d, m)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Floating IP",
				Detail:   fmt.Sprint(err),
			})
			return diags
		}
	}
	resourceFloatingIPRead(ctx, d, m)

	return diags
}

func resourceFloatingIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*idcloudhost.APIClient)
	var diags diag.Diagnostics
	ipAddress := d.Id()
	fipApi := c.FloatingIP
	err := fipApi.Get(ipAddress)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Floating IP",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}

	err = setFloatingIP(d, fipApi.FloatingIP)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Floating IP",
			Detail:   fmt.Sprint(err),
		})
		return diags
	}

	return diags
}

func resourceFloatingIPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*idcloudhost.APIClient)
	fipApi := c.FloatingIP
	ipAddress := d.Id()
	if d.HasChanges("billing_account_id", "name") {
		billingAccountId := d.Get("billing_account_id").(int)
		name := d.Get("name").(string)
		err := fipApi.Update(name, billingAccountId, ipAddress)
		if err != nil {
			diags = append(diags, diag.Diagnostic{})
			return diags
		}
	}

	if d.HasChange("assigned_to") {
		var err error
		assignedUuid := d.Get("assigned_to").(string)
		if assignedUuid != "" {
			err = fipApi.Assign(fipApi.FloatingIP.Address, assignedUuid)
		} else {
			err = fipApi.Unassign(fipApi.FloatingIP.Address)
		}
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update Floating IP",
				Detail:   fmt.Sprintf("cannot (un)assign to specified UUID: %s", assignedUuid),
			})
			return diags
		}
	}

	resourceFloatingIPRead(ctx, d, m)

	return diags
}

func resourceFloatingIPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*idcloudhost.APIClient)
	ipAddress := d.Id()
	fipApi := c.FloatingIP
	err := fipApi.Delete(ipAddress)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
