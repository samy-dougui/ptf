rule "my_first_rule" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attributes = "metadata.hdi_version"
    values     = "2013-09-01"
  }
  severity      = "warning"
  error_message = ""
}