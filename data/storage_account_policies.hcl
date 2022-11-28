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

policy "databricks_secret" {
  filter {
    type = "databricks_secret_scope"
  }
  condition {
    attribute = "name"
    operator  = "in"
    values    = ["container-storage-carrefour-samy-", "container-storage-carrefour-samy-2"]
  }
  severity = "warning"
}
