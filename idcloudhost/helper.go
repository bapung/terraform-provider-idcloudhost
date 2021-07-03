package idcloudhost

import (
	"encoding/json"
	"strconv"

	"github.com/bapung/idcloudhost-go-client-library/idcloudhost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func setVmResource(d *schema.ResourceData, vm *idcloudhost.VM) error {
	var storageList []map[string]interface{}
	storageJson, err := json.Marshal(vm.Storage)
	if err != nil {
		return err
	}
	err = json.Unmarshal(storageJson, &storageList)
	if err != nil {
		return err
	}
	if err := d.Set("storage", storageList); err != nil {
		return err
	}
	if err := d.Set("backup", vm.Backup); err != nil {
		return err
	}
	if err := d.Set("billing_account_id", vm.BillingAccount); err != nil {
		return err
	}
	if err := d.Set("created_at", vm.CreatedAt); err != nil {
		return err
	}
	if err := d.Set("description", vm.Description); err != nil {
		return err
	}
	if err := d.Set("hostname", vm.Hostname); err != nil {
		return err
	}
	if err := d.Set("hypervisor_id", vm.HypervisorId); err != nil {
		return err
	}
	if err := d.Set("id", strconv.Itoa(vm.Id)); err != nil {
		return err
	}
	if err := d.Set("mac", vm.MACAddress); err != nil {
		return err
	}
	if err := d.Set("memory", vm.Memory); err != nil {
		return err
	}
	if err := d.Set("name", vm.Name); err != nil {
		return err
	}
	if err := d.Set("os_name", vm.OSName); err != nil {
		return err
	}
	if err := d.Set("os_version", vm.OSVersion); err != nil {
		return err
	}
	if err := d.Set("private_ipv4", vm.PrivateIPv4); err != nil {
		return err
	}
	if err := d.Set("status", vm.Status); err != nil {
		return err
	}
	if err := d.Set("tags", vm.Tags); err != nil {
		return err
	}
	if err := d.Set("updated_at", vm.UpdatedAt); err != nil {
		return err
	}
	if err := d.Set("user_id", vm.UserId); err != nil {
		return err
	}
	if err := d.Set("username", vm.Username); err != nil {
		return err
	}
	if err := d.Set("uuid", vm.UUID); err != nil {
		return err
	}
	if err := d.Set("vcpu", vm.VCPU); err != nil {
		return err
	}
	return nil
}

func adaptVMListStructToMap(vmList *[]idcloudhost.VM) ([]map[string]interface{}, error) {
	var vmMapList []map[string]interface{}
	vmJson, err := json.Marshal(vmMapList)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(vmJson, &vmMapList)
	if err != nil {
		return nil, err
	}
	return vmMapList, nil
}

func setFloatingIP(d *schema.ResourceData, fip *idcloudhost.FloatingIP) error {
	if err := d.Set("id", strconv.Itoa(fip.ID)); err != nil {
		return err
	}
	if err := d.Set("address", fip.Address); err != nil {
		return err
	}
	if err := d.Set("user_id", fip.UserID); err != nil {
		return err
	}
	if err := d.Set("billing_account_id", fip.UserID); err != nil {
		return err
	}
	if err := d.Set("type", fip.Type); err != nil {
		return err
	}
	if err := d.Set("network_id", fip.NetworkID); err != nil {
		return err
	}
	if err := d.Set("name", fip.Name); err != nil {
		return err
	}
	if err := d.Set("enabled", fip.Enabled); err != nil {
		return err
	}
	if err := d.Set("created_at", fip.CreatedAt); err != nil {
		return err
	}
	if err := d.Set("updated_at", fip.UpdatedAt); err != nil {
		return err
	}
	if err := d.Set("assigned_to", fip.AssignedTo); err != nil {
		return err
	}

	return nil
}
