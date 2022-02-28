rule "my_first_rule" {
  filter {
    type = "databricks_dbfs_file"
  }
  condition     = "my_condition"
  severity      = ""
  error_message = ""
}