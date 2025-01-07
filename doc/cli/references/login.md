# `src login`


## Flags

| Name | Description | Default Value |
|------|-------------|---------------|
| `-dump-requests` | Log GraphQL requests and responses to stdout | `false` |
| `-get-curl` | Print the curl command for executing this query and exit (WARNING: includes printing your access token!) | `false` |
| `-insecure-skip-verify` | Skip validation of TLS certificates against trusted chains | `false` |
| `-trace` | Log the trace ID for requests. See https://docs.khulnasoft.com/admin/observability/tracing | `false` |
| `-user-agent-telemetry` | Include the operating system and architecture in the User-Agent sent with requests to Khulnasoft | `true` |


## Usage

```
'src login' helps you authenticate 'src' to access a Khulnasoft instance with your user credentials.

Usage:

    src login KHULNASOFT_URL

Examples:

  Authenticate to a Khulnasoft instance at https://sourcegraph.example.com:

    $ src login https://sourcegraph.example.com

  Authenticate to Khulnasoft.com:

    $ src login https://khulnasoft.com

  -dump-requests
    	Log GraphQL requests and responses to stdout
  -get-curl
    	Print the curl command for executing this query and exit (WARNING: includes printing your access token!)
  -insecure-skip-verify
    	Skip validation of TLS certificates against trusted chains
  -trace
    	Log the trace ID for requests. See https://docs.khulnasoft.com/admin/observability/tracing
  -user-agent-telemetry
    	Include the operating system and architecture in the User-Agent sent with requests to Khulnasoft (default true)

```
	
