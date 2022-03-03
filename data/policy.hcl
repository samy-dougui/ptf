rule "my_first_rule" {
  filter {
    type = "azurerm_storage_container"
  }
  condition     = "my_first_condition"
  severity      = "warning"
  error_message = ""
}

rule "my_second_rule" {
  filter {
    type = "azurerm_resource_group"
  }
  condition     = "my_second_condition"
  severity      = "error"
  error_message = ""
}
