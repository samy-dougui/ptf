rule "my_goal_rule" {
  filter {
    type = "azurerm_resource_group" # if more than one filter, should be an "and"
  }
  condition {
    attribute = ""
    operator = "" # <, <=, >, >=, !=, =, in, not in, re (regex)
    values = "" # list of values (if operator = 'in') or single value
  }
  severity      = "error" # For now, it's always an error, could be in the future a warning
  error_message = "" # Could be nice to customise it, should be a default one for now
}