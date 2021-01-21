---
subcategory: "Domains"
---

# Data Source: azuread_domains

Use this data source to access information about existing Domains within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Directory.Read.All` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_domains" "aad_domains" {}

output "domains" {
  value = data.azuread_domains.aad_domains.domains
}
```

## Argument Reference

* `include_unverified` - (Optional) Set to `true` if unverified Azure AD domains should be included. Defaults to `false`.
* `only_default` - (Optional) Set to `true` to only return the default domain.
* `only_initial` - (Optional) Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.

~> **NOTE:** If `include_unverified` is set to `true` you cannot specify `only_default` or `only_initial`. Additionally, you cannot combine `only_default` with `only_initial`.

## Attributes Reference

* `domains` - A list of domains. Each `domain` object provides the attributes documented below.

`domain` object exports the following:

* `authentication_type` - The authentication type of the domain (Managed or Federated).
* `domain_name` - The name of the domain.
* `is_default` - `True` if this is the default domain that is used for user creation.
* `is_initial` - `True` if this is the initial domain created by Azure Active Directory.
* `is_verified` - `True` if the domain has completed domain ownership verification.
