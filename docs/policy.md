---
layout: default
title: Policy
nav_order: 3
has_children: false
---

# Policy

## What is a policy

## How do we define a policy in PTF

```hcl
policy "policy_example" {
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "queue_properties.[*].hour_metrics.[*].retention_policy_days"
    operator  = ">="
    values    = 6
  }
}
```

## Examples
