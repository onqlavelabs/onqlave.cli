
## Setup

Init configuration

```bash
  go run main.go config init
```

## Authentication

Sign up new account

```bash
  go run main.go auth signup [email] --tenant_name [tenant_name] --full_name[full_name]
```

Login

```bash
  go run main.go auth login [email] --tenant_name [tenant_name]
```

## User commands

List users

```bash
  go run main.go user list
```

## API Key commands

API key base info

```bash
  go run main.go key base
```

List api keys

```bash
  go run main.go key list
```

Describe an api key

```bash
  go run main.go key describe [keyid]
```

Add an api key

```bash
  go run main.go key add --key_application_id [appid] --key_arx_id [arxid] --key_application_technology [client|server]
```

Delete an api key

```bash
  go run main.go key delete [keyid]
```

## Arx commands

Arx base info

```bash
  go run main.go arx base
```

List arx

```bash
  go run main.go arx list
```

Describe a arx

```bash
  go run main.go arx describe [arxid]
```

Add a arx

```bash
  go run main.go arx add [arx_name] --arx_provider [gcp|aws|azure] --arx_purpose [development|testing|staging|production] --arx_region [au|us|sg|gb] --arx_encryption_method [aes-gcm-128|cha-cha-20-poly-1305|aes-gcm-256] --arx_rotation_cycle [monthly|3-monthly|6-monthly|annually] --arx_owner [userid] --arx_spend_limit [int] --arx_is_default=[true|false]
```

Edit a arx

```bash
  go run main.go arx edit [arxid] --arx_name [arx_name] --arx_region [au|us|sg|gb] --arx_rotation_cycle [monthly|3-monthly|6-monthly|annually] --arx_owner [userid] --arx_spend_limit [int] --arx_is_default=[true|false]
```

Delete a arx

```bash
  go run main.go arx delete [arxid]
```

Seal a arx

```bash
  go run main.go arx seal [arxid]
```

Unseal a arx

```bash
  go run main.go arx unseal [arxid]
```

Set a arx as default

```bash
  go run main.go arx default [arxid]
```

Reinitialize a arx

```bash
  go run main.go arx retry [arxid]
```
