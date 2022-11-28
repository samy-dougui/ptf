# tftest

## Port attributes definition

- Policy
    - name
    - type/target (applied on resources or on config/meta)
    - severity (warning or error)
    - error_message ("go string template" with two variables: expected value / value received)
    - disabled (default: false)
    - passed (default: true): will contain the information if the policy has been validated against the target
    - filter
    - condition
- Resource
    - address
    - mode (?) (terraform resource or data: should only focus on terraform resources)
    - type
    - provider
    - values (merge of after_values and after_sensitive_values)
    - action (no-op, deleted, etc...)
- Plan Configuration
    - variables
    - terraform version
    - providers configuration (list of provider configuration):
        - provider name
        - version_constraint

## core

The main method should be "validate" (core.Validate) that takes as input a list of policy, a list of resource (type
Resource), and a configuration (type Config)
This method should return a list of dictionary (json) (type Policy output) with keys:
- name: policy_name
- validated: true or false
- resource_list (only if validated is false)
  - resource_address
  - expected attribute
  - attribute received
  - error message
