rule "dir_rule" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "name"
    operator  = "re"
    values    = "([aA-zZ]+)_([aA-zZ]+)_([aA-zZ]+)"
  }
  severity      = "error"
  error_message = ""
}

rule "azure_storage_container_metadata" {
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
}

rule "azure_tag" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "tags"
    operator  = "="
    values    = {
      "environment" : "testfsdjfsdkj",
      "id" : 1.5,
      "key_missing": "value",
      "is_prod": true
    }
  }
}
