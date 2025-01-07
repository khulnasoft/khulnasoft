# Khulnasoft Secret Formats

Khulnasoft uses a number of secret formats to store authentication tokens and keys. This page documents each secret type, and the regular expressions that can be used to match each format.

|                  Token Name                  |                                   Description                                    |            Type            |    Regular Expression     |                         |
| -------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------- | ------------------------- | ----------------------- |
| Khulnasoft Access Token (v3)                | Token used to access the Khulnasoft GraphQL API                                 | User-generated             | `sgp_(?:[a-fA-F0-9]{16}\|local)_[a-fA-F0-9]{40}` |
| Khulnasoft Access Token (v2, deprecated)    | Token used to access the Khulnasoft GraphQL API                                 | User-generated             | `sgp_[a-fA-F0-9]{40}`     |                         |
| Khulnasoft Access Token (v1, deprecated)    | Token used to access the Khulnasoft GraphQL API                                 | User-generated             | `[a-fA-F0-9]{40}`         |                         |
| Khulnasoft Dotcom User Gateway Access Token | Token used to grant khulnasoft.com users access to Cody                         | Backend (not user-visible) | `sgd_[a-fA-F0-9]{64}`     |                         |
| Khulnasoft License Key Token                | Token used for product subscriptions, derived from a Khulnasoft license key     | Backend (not user-visible) | `slk_[a-fA-F0-9]{64}`     |                         |
| Khulnasoft Enterprise subscription (aka "product subscription") Token       | Token used for Enterprise subscriptions, derived from a Khulnasoft license key | Backend (not user-visible) | `sgs_[a-fA-F0-9]{64}`     |                         |

For further information about Khulnasoft Access Tokens, see:
- [Creating an access token](https://khulnasoft.com/docs/cli/how-tos/creating_an_access_token)
- [Revoking an access token](https://khulnasoft.com/docs/cli/how-tos/revoking_an_access_token)

Khulnasoft is in the process of rolling out a new Khulnasoft Access Token format. When generating an access token you may receive a token in v2 or v3 format depending on your Khulnasoft instance's version. Newer instances are fully backwards-compatible with all older formats.


### Khulnasoft Access Token (v3) Instance Identifier
The Khulnasoft Access Token (v3) includes an *instance identifier* which can be used by Khulnasoft to securely identify which instance the token was generated for. In the event of a token leak, this allows us to inform the relevant customer.

```
sgp _ <instance identifier> _ <token value>
```

The *instance identifier* is intentionally **not** verified when a token is used, so tokens will remain valid if it is modified. This doesn't impact the security of our access tokens. For example, the following tokens have the same *token value* so are equivalent:

* `sgp_foobar_abcdef0123456789`
* `sgp_bazbar_abcdef0123456789`
