rule "azure storage container metadata" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "metadata.hdi_version"
    operator   = "="
    values     = 2013
  }
  severity      = "error"
  error_message = ""
}