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

Detailed setup and usage instructions, including direct credentials and
optional host-stack integration, are in `docs/USAGE.md`.

## Nginx Example

`examples/nginx-hello` is a source example only. It is not run by `make build`
or `make test`, and no live Pulumi stack configuration is committed. It can
either read Coolify credentials from an existing host stack or use direct
endpoint and token configuration without a host stack. In both modes it adopts
an existing Coolify server/project and deploys `nginx:latest` without requiring
a Git deploy key.

Copy one of the example configurations, replace every placeholder, and review
the preview before applying it:

```bash
make build
cd examples/nginx-hello
pulumi stack init nginx-hello
cp Pulumi.nginx-hello.example.yaml Pulumi.nginx-hello.yaml
# Host-stack mode: edit Pulumi.nginx-hello.yaml with your host stack and
# Coolify identifiers.
# Direct mode: use Pulumi.nginx-hello.direct.example.yaml instead and provide
# the endpoint and Coolify identifiers, then set the token with:
# pulumi config set --secret nginx:token your-coolify-api-token
PATH="../../bin:$PATH" pulumi preview --stack nginx-hello
PATH="../../bin:$PATH" pulumi up --stack nginx-hello
```

The example is intentionally opt-in. Do not run `pulumi up` unless you want to
create and deploy the demonstration application in the configured Coolify
project.

The `http://` domain is intentional: Caddy terminates public TLS and forwards
to Coolify's internal proxy. The stack exports the public `https://` URL.

## License

The bridge is Apache-2.0 licensed. The generated provider is a derived work of
the upstream Coolify Terraform provider, which is MPL-2.0 licensed.
