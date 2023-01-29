---
layout: default
title: Operators
nav_order: 6
has_children: false
---

# Operators

PTF supports the following operators:

- "=": Equality
- ">": Strict Superiority
- ">=": Superiority
- "<": Strict Inferiority
- "<=": Inferiority
- "re": Regex
- "in": Inclusion
- "not in": Not Inclusion

## Equality

The Equality Operator supports the following types:

- Integer
- Float
- String
- Boolean

> The equality between dictionaries will be added in a following release.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = "="
  values    = "Accepted_value"
}
```

## Strict Superiority

The Strict Superiority Operator only supports Integer and Float types.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = ">"
  values    = 1.5
}
```

## Superiority

The Superiority Operator only supports Integer and Float types.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = ">="
  values    = 1.5
}
```

## Strict Inferiority

The Strict Inferiority Operator only supports Integer and Float types.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = "<"
  values    = 1.5
}
```

## Inferiority

The Inferiority Operator only supports Integer and Float types.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = "<="
  values    = 1.5
}
```

## Regex

The Regex Operator only supports String types.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = "re"
  values    = "expected_regex"
}
```

## Inclusion

The Inclusion Operator supports the List Type. All the elements of the provided list should have the same type.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = "in"
  values    = ["accepted_value_1", "accepted_value_2"]
}
```

## Not Inclusion

The Not Inclusion Operator support the List Type. All the elements of the provided list should have the same type.

### Example

```terraform
condition {
  attribute = "attribute_1"
  operator  = "not in"
  values    = ["not_accepted_value_1", "not_accepted_value_2"]
}
```