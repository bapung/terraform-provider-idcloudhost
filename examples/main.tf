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

resource "idcloudhost_vm" "testvm" {
    name = "testvm"
    os_name = "ubuntu"
    os_version= "18.04"
    disks = 20
    vcpu = 1
    memory = 1024
    username = "example"
    initial_password = "Password123"
    billing_account = 1200132376
    backup = false
}

resource "idcloudhost_floating_ip" "testip" {
    name = "my_test_ip"
    billing_account_id = 1200132376
}
