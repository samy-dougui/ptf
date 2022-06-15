#rule "azure_storage_container_has_legal_hold" {
#  filter {
#    type = "azurerm_storage_container"
#  }
#  condition {
#    attribute = "has_legal_hold"
#    operator  = "="
#    values    = false
#  }
#}

rule "azure_storage_container_name" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "name"
    operator  = "not in"
    values    = ["eltcode_regex_test"]
  }
  severity = "error"
  disabled = true
}
#rule "azure_storage_container_metadata" {
#  filter {
#    type = "azurerm_storage_container"
#  }
#  condition {
#    attribute = "metadata.hdi_version"
#    operator  = "="
#    values    = "2013"
#  }
#  severity      = "warning"
#  error_message = ""
#}

#rule "azure_tag" {
#  filter {
#    type = "azurerm_storage_container"
#  }
#  condition {
#    attribute = "tags"
#    operator  = "="
#    values    = {
#      "environment" : "testfsdjfsdkj",
#      "id" : 1.5,
#      "key_missing" : "value",
#      "is_prod" : true
#    }
#  }
#}
