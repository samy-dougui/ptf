policy "azurem_storage_account_network_rules" {
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "network_rules.[*].default_action"
    operator  = "="
    values    = "Allow"
  }
  severity = "error"
}

policy "azure_storage_container_name_pattern" {
  name = "Azure storage account naming convention"
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "name"
    operator  = "re"
    values    = "([aA-zZ]+)_([aA-zZ]+)_([aA-zZ]+)"
  }
  severity = "warning"
  disabled = false
}