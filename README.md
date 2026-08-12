# Pulumi Coolify Provider

This repository contains a Pulumi provider for Coolify, bridged from
`github.com/coolify-terraform/terraform-provider-coolify` v0.1.13.

The bridge uses Pulumi's Terraform Plugin Framework adapter and exposes the
upstream provider's applications, databases, services, projects, servers,
storage, environment variables, deployments, backups, and data sources.

## Provider Development

The provider requires Go 1.26.5 or newer, Pulumi, and the generated Go SDK.

```bash
make build
make test
```

For local Pulumi programs, install the provider binary into the local plugin
cache:

```bash
make plugin-install
```

The provider imports the upstream framework provider through
`provider/shim`. This is required because the upstream provider keeps its
provider constructor under Go's `internal` import boundary.

## Nginx Example

`examples/nginx-hello` is a source example only. It is not run by `make build`
or `make test`, and no live Pulumi stack configuration is committed. If you
run it manually, it adopts an existing Coolify server/project and deploys
`nginx:latest` without requiring a Git deploy key.

Copy the example configuration, replace every placeholder, and review the
preview before applying it:

```bash
make build
cd examples/nginx-hello
pulumi stack init nginx-hello
cp Pulumi.nginx-hello.example.yaml Pulumi.nginx-hello.yaml
# Edit Pulumi.nginx-hello.yaml with your own stack and Coolify identifiers.
PATH="../../bin:$PATH" pulumi preview --stack nginx-hello
PATH="../../bin:$PATH" pulumi up --stack nginx-hello
```

The example is intentionally opt-in. Do not run `pulumi up` unless you want to
create and deploy the demonstration application in the configured Coolify
project.

The `http://` domain is intentional: Caddy terminates public TLS and forwards
to Coolify's internal proxy. The stack exports the public `https://` URL.

## Real Showboat Deployment

The real `expat-map-guide` deployment lives in the host repository at
`coolify/expat-map-guide`, not in this provider's examples. It uses this
provider through the local SDK and adopts the existing host stack through a
Pulumi `StackReference`.

The private GitLab repository needs a dedicated read-only deploy key registered
on GitLab. Production also requires the Stripe values enforced by the
application's configuration validation.

## License

The bridge is Apache-2.0 licensed. The generated provider is a derived work of
the upstream Coolify Terraform provider, which is MPL-2.0 licensed.
