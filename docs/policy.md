---
layout: default
title: Policy
nav_order: 5
has_children: false
---

# Policy

## What is a policy

In PTF, a policy is composed of two main elements:

- a filter block, which allows you to select a subset of the resources.
    - _As of right now, it's only possible to filter base on the type of the resource._
- a condition block, which allows you to test the value of the resource attribute. This condition will be applied to all
  the resources that respects the filter criteria. It's composed of three elements:
    - `attribute`: the attribute of the resource you'd like to check.
    - `operator`: the operator that is going to done between the attribute and values. The list of operators supported
      is [here](./operators.md)
    - `values`: the values the attribute is supposed to verify, in respect to the operator.

A policy has other attributes:

- `name`: the name of the policy
- `severity`: level of the alert if the policy is not respected (warning or error) [**default to error**]
- `disabled`: define if the policy is active or not [**default to false**]
- `error_message`: message displayed if the policy is not validated
- `target`: the target of the policy (resource or meta) [**default to resource**]

## Policy schema

To define a policy, we use HCL (Hashicorp Configuration language).

A policy should be defined in a file with the `.hcl` extension and follow this schema:

```
policy "policy_name" {
  target = "resource_or_meta"
  filter {
    type = "type_of_the_resource"
  }
  condition {
    attribute = "attribute_you_want_to_verify"
    operator  = "operation_you_want_to_apply"
    values    = "values"
  }
  severity      = "error_or_warn"
  error_message = "error_message"
  disabled = "is_the_policy_active" 
}
```

## Example

### Example 1

```terraform
policy "azurem_storage_account_network_rules" {
  filter {
    type = "azurerm_storage_account"
  }
  condition {
    attribute = "network_rules.[*].default_action"
    operator  = "="
    values    = "Deny"
  }
}
```

The policy `azurem_storage_account_network_rules` makes sure that for all azure storage account resources, the default
action of every network rule is "Deny".

### Example 2

```terraform
policy "azure_storage_container_name_pattern" {
  filter {
    type = "azurerm_storage_container"
  }
  condition {
    attribute = "name"
    operator  = "re"
    values    = "([aA-zZ]+)_([aA-zZ]+)_([aA-zZ]+)"
  }
}
```

The policy `azure_storage_container_name_pattern` makes sure that all azure storage container resources have a name that
respects the regex `([aA-zZ]+)_([aA-zZ]+)_([aA-zZ]+)`

## Attribute

PTF allows you to query nested attributes. This section explains how.

### Non-nested attribute

```json
{
  "foo": "bar"
}
```

To get the value of `foo` (bar), the attribute "attribute" of the policy should be `foo`.

### Nested attribute

```json
{
  "foo": {
    "nested_foo": "nested_bar"
  }
}
```

To get the value of the attribute `nested_foo` (nested_bar), the attribute "attribute" of the policy should
be `foo.nested_foo`.

PTF allows you to query attributes as nested as you want.

### List attribute

```json
{
  "distinct_foo": [
    {
      "foo_1": "bar_1"
    },
    {
      "foo_2": "bar_2"
    }
  ],
  "same_foo": [
    {
      "foo_3": "bar_3"
    },
    {
      "foo_3": "bar_4"
    }
  ]
}
```

Some Terraform providers return list of attributes, PTF allows you to query them like so:

- To retrieve the value of `foo_1` (bar_1), the attribute "attribute" of the policy should be `distinct_foo.[0].foo_1`
- To retrieve the value of both `foo_3` ([bar_3, bar_4]), the attribute "attribute" of the policy should
  be `same_foo.[*].foo_3`
    - You can then make sure that all the foo_3 attributes respects your policy.

This can result in complex queries like:

```json
{
  "distinct_foo": [
    {
      "foo_1": [
        {
          "foo_2": "bar_1"
        },
        {
          "foo_2": "bar_3"
        }
      ]
    },
    {
      "foo_3": "bar_3"
    }
  ]
}
```

To retrieve all the values of foo_2, this would be: `distinct_foo.[0].foo_1.[*].foo_2`
