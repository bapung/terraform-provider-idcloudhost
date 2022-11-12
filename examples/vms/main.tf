terraform {
  required_providers {
    idcloudhost = {
      version = "0.2.0"
      source  = "github.com/bapung/idcloudhost"
    }
  }
}

data "idcloudhost_vms" "all" {}

output "all_vms" {
  value = data.idcloudhost_vms.all
}