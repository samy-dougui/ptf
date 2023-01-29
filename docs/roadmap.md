---
layout: default
title: Roadmap
nav_order: 7
has_children: false
---

# Roadmap

Here are the list of the features that will be implemented in the future:

## Main features

- Supports of variables in policies
- Supports of HCL2 functions
- Supports more filtering capabilities
- Creation of meta-policy (policies on the metadata of the plan)
- Creation of PTF Policy modules (with support for import)

## Smaller features

- Supports of pretty name for policy
- For raw mode in CLI, supports of a flag "output" to output the raw results in a json file.
- Refactoring of the pretty output (in particular for the attribute query with multiple attributes)
- Support policy filtering on CLI

## Internal

- Refactoring of the logger
- Refactoring of the file loader
- Better error management for types in operators package