terraform {
  required_providers {
    idcloudhost = {
      version = "0.2"
      source  = "github.com/bapung/idcloudhost"
    }
  }
}

provider "idcloudhost" {}

module "test-vm" {
  source = "./vms"

}

output "test-vm" {
  value = module.test-vm.all_vms.vms
}

resource "idcloudhost_vm" "testvm" {
    name = "testvm"
    os_name = "ubuntu"
    os_version= "16.04"
    disks = 20
    vcpu = 1
    memory = 2048
    username = "example"
    initial_password = "Password123"
    billing_account = 1200132376
}

output "vm_created" {
  value = idcloudhost_vm.testvm
}
