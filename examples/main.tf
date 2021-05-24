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