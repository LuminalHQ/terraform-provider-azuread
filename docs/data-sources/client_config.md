---
subcategory: "Base"
---

# Data Source: azuread_client_config

Use this data source to access the configuration of the AzureAD provider.

## Example Usage

```hcl
data "azuread_client_config" "current" {
}

output "account_id" {
  value = data.azuread_client_config.current.client_id
}
```

## Argument Reference

There are no arguments available for this data source.

## Attributes Reference

* `client_id` is set to the Client ID (Application ID).
* `object_id` is set to the Object ID of the authenticated principal.
* `tenant_id` is set to the Tenant ID.
