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
    disks = 10
    vcpu = 1
    memory = 512
    username = "example"
    initial_password = "Password123"
}

output "vm_created" {
  value = idcloudhost_vm.testvm
}
