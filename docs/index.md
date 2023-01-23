---
layout: default
title: PTF
nav_order: 1
has_children: false
---

# PTF

PTF (Policy Terraform) is a Policy-as-Code framework to control your Terraform deployment.

## Why PTF ?

PTF allows you to easily define policy using the same language as Terraform.


## Why not OPA ?

Open Policy Agent (OPA) is, as the documentation says, `an open source, general-purpose policy engine that unifies policy enforcement across the stack`. OPA defines policies as code in a language called Rego.

As great as OPA is, it uses a completely different language than HCL which makes the adoption of Policy as Code more complicated for Terraform. The goal of PTF was to create policies in the same language as Terraform to make it easier for teams to create their own policies.

## Why not Sentinel ?

Sentinel is, as the documentation says, `is a language and framework for policy built to be embedded in existing software to enable fine-grained, logic-based policy decisions. A policy describes under what circumstances certain behaviors are allowed. Sentinel is an enterprise feature of HashiCorp Consul, Nomad, Terraform, and Vault.` Sentinel defines policy as code but the key thing here is that it's an enterprise feature which makes the adoption of policy as code harder for companies that do not use Terraform Enterprise.