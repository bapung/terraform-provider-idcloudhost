terraform {
  required_providers {
    idcloudhost = {
      version = "0.2"
      source  = "github.com/bapung/idcloudhost"
    }
  }
}

data "idcloudhost_vms" "all" {}

# Returns all coffees
output "all_vms" {
  value = data.idcloudhost_vms.all
}