policy "azurem_storage_account_network_rules" {
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "network_rules.[*].default_action"
    operator  = "="
    values    = "Deny"
  }
  severity      = "error"
  error_message = ""
  disabled      = false
}
