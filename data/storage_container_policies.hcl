policy "azure_storage_container_name_pattern" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "name"
    operator  = "re"
    values    = "([aA-zZ]+)_([aA-zZ]+)_([aA-zZ]+)"
  }
  severity      = "warning"
  error_message = ""
  disabled      = false
}

policy "azure_storage_container_metadata" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "metadata.hdi_version"
    operator  = "="
    values    = "2013"
  }
  severity      = "warning"
  error_message = ""
  disabled      = false
}

policy "azure_storage_container_has_legal_hold" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "has_legal_hold"
    operator  = "="
    values    = false
  }
  disabled = false
}

policy "azure_tag" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "tags"
    operator  = "="
    values    = {
      "environment" : "prod",
      "product" : "data product",
      "team" : "champagne",
    }
  }
  severity = "warning"
  disabled = false
}
