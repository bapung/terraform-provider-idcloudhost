# idCloudHost Terraform Provider

# Resources

### Example Usage
```hcl
terraform {
  required_providers {
    idcloudhost = {
      version = "0.1.0"
      source  = "bapung/idcloudhost"
    }
  }
}

provider "idcloudhost" {
    auth_token = "better-to-specify-via-secure-state-or-env"
}

module "test-vm" {
  source = "./vms"

}

resource "idcloudhost_vm" "testvm" {
    name = "testvm"
    os_name = "ubuntu"
    os_version= "18.04"
    disks = 20
    vcpu = 1
    memory = 1024
    username = "example"
    initial_password = "Password123"
    billing_account_id = 1337
    backup = false
}

resource "idcloudhost_floating_ip" "testip" {
    name = "my_test_ip"
    billing_account_id = 1337
}

```
# Notes
Early work. Need to write proper unittest and accceptance test and add more resource in the future.