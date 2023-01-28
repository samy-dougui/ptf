---
layout: default
title: PTF
nav_order: 1
has_children: false
---

# PTF

PTF (Policy Terraform) is a Policy-as-Code framework to control your Terraform deployment.

# Installation

To install PTF, follow [this page](./installation.md).

# How does it work?

Before applying the changes, Terraform creates a plan and shows you all the action it will do. Before applying these
changes, you need to make sure they are compliant with your standards. That's where PTF comes into play: it will sit
between your Terraform plan and your Terraform apply, and make sure the Terraform plan respects the policies you have
defined.

# Basic Usage

Once you have PTF installed on your machine, you can now control your Terraform plan.

To obtain your plan, you can follow this procedure:

1. `terraform init`
2. `terraform plan --out tfplan.binary`
3. `terraform show -json tfplan.binary > tfplan.json`

Once you have your Terraform plan, you can control if it's compliant by
executing: `ptf control -p tfplan.json --chdir policies/`
This will control that your `tfplan.json` respects all the policies defined in the directory `policies/`. To learn how
to define a policy, head [here](./policy.md)

If your checks are successful, you can then apply the changes: `terraform apply tfplan.binary`

# Why PTF ?

PTF allows you to easily define policy using the same language as Terraform (HCL) and allows you to define your own
policy.

## Why not OPA ?

Open Policy Agent (OPA) is, as the documentation
says, `an open source, general-purpose policy engine that unifies policy enforcement across the stack`. OPA defines
policies as code in a language called Rego.

As great as OPA is, it uses a completely different language than HCL which makes the adoption of Policy as Code more
complicated for Terraform. The goal of PTF was to create policies in the same language as Terraform to make it easier
for teams to create their own policies.

## Why not Sentinel ?

Sentinel is, as the documentation
says, `is a language and framework for policy built to be embedded in existing software to enable fine-grained, logic-based policy decisions. A policy describes under what circumstances certain behaviors are allowed. Sentinel is an enterprise feature of HashiCorp Consul, Nomad, Terraform, and Vault.`
It has been developed by Hashicorp, the creators of Terraform.

Sentinel defines policy as code but the key thing here is hat it's an enterprise feature which makes the adoption of
policy as code harder for companies that do not use Terraform Enterprise.
