policy "policy_1" {
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "network_rules.[*].default_action"
    operator  = "="
    values    = "Allow"
  }
  severity      = "error"
  error_message = ""
  disabled      = false
}

policy "policy_2" {
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "foo"
    operator  = "="
    values    = "bar2"
  }
  severity      = "error"
  error_message = ""
  disabled      = false
}

policy "policy_3" {
  target = "resource"
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "foo"
    operator  = ">="
    values    = "bar2"
  }
  severity      = "error"
  error_message = ""
  disabled      = false
}