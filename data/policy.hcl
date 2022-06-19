policy "azure_storage_container_name_pattern" {
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
  disabled      = true
}

policy "azurem_storage_account_blob_properties" {
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "name"
    operator  = "re"
    values    = "([aA-zZ]+)"
  }
  severity      = "error"
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
  disabled      = true
}

policy "dbt_file_size" {
  filter {
    type = "databricks_dbfs_file"
  }
  condition {
    attribute = "file_size"
    operator  = "="
    values    = 0
  }
  disabled = true
}
