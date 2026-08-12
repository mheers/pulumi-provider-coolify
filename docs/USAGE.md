# Coolify Provider Usage

This provider manages Coolify resources through Pulumi. It can be used as a
standalone provider or from a Pulumi program that obtains its credentials from
another host stack.

The host stack is optional. The provider itself does not require a host stack.

## Prerequisites

- Go 1.26.5 or newer
- Pulumi CLI
- A reachable Coolify instance
- A Coolify API token

## Build And Install

Build the provider and generated Go SDK from the repository root:

```bash
make build
```

Install the development provider into the local Pulumi plugin cache:

```bash
make plugin-install
```

Run the test suites:

```bash
make test
```

## Direct Provider Configuration

Use direct configuration when the Pulumi program already has the Coolify
endpoint and API token. The token should be stored as a Pulumi secret.

```go
package main

import (
	"github.com/mheers/pulumi-provider-coolify/sdk/go/coolify"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "coolify")
		provider, err := coolify.NewProvider(ctx, "coolify", &coolify.ProviderArgs{
			Endpoint: pulumi.StringPtr(cfg.Require("endpoint")),
			Token:    cfg.RequireSecret("token").ToStringPtrOutput(),
		})
		if err != nil {
			return err
		}

		_, err = coolify.GetProjects(ctx, &coolify.GetProjectsArgs{}, pulumi.Provider(provider))
		return err
	})
}
```

Configure the stack:

```bash
pulumi config set coolify:endpoint https://coolify.example.com
pulumi config set --secret coolify:token your-coolify-api-token
```

The provider also supports environment variables when no explicit provider
arguments are supplied:

```bash
export COOLIFY_ENDPOINT=https://coolify.example.com
export COOLIFY_TOKEN=your-coolify-api-token
```

Additional provider settings include `caCert`, `insecure`, Cloudflare Access
credentials, and retry limits. Prefer `caCert` over `insecure` for custom TLS
certificates.

## Host-Stack Configuration

Use a `pulumi.StackReference` when another Pulumi stack owns the Coolify host
and exports its endpoint and API token. The referenced stack must export values
matching the keys used by the program.

```go
hostStack, err := pulumi.NewStackReference(ctx, "coolify-host", &pulumi.StackReferenceArgs{
	Name: pulumi.String("org/host-project/host-stack"),
})
if err != nil {
	return err
}

endpoint := hostStack.GetStringOutput(pulumi.String("coolify:endpoint")).ToStringPtrOutput()
token := hostStack.GetOutput(pulumi.String("coolify:token")).ApplyT(func(value any) *string {
	valueString, ok := value.(string)
	if !ok {
		return nil
	}
	return &valueString
}).(pulumi.StringPtrOutput)

provider, err := coolify.NewProvider(ctx, "coolify", &coolify.ProviderArgs{
	Endpoint: endpoint,
	Token:    token,
})
```

The `nginx-hello` example uses the host-stack output names
`showboat:coolify:dashboard` and `showboat:coolify:apiToken`. Change those keys
when integrating with a different host stack.

## Nginx Example

The source example in `examples/nginx-hello` adopts an existing Coolify project
and server, creates an nginx Docker-image application, and waits for its
deployment to complete. It does not create the Coolify server or project.

Initialize a stack and build the provider:

```bash
make build
cd examples/nginx-hello
pulumi stack init nginx-hello
```

### Host-Stack Mode

Copy the host-stack configuration and set the host stack, project, server, and
domain values:

```bash
cp Pulumi.nginx-hello.example.yaml Pulumi.nginx-hello.yaml
```

This mode reads the endpoint and token from the configured host stack.

### Direct Mode

Copy the direct configuration and set the endpoint, project, server, and
domain values:

```bash
cp Pulumi.nginx-hello.direct.example.yaml Pulumi.nginx-hello.yaml
pulumi config set --secret nginx:token your-coolify-api-token
```

This mode does not create a `StackReference`. It reads the endpoint from
`nginx:endpoint` and the token from the secret `nginx:token`.

Run either mode with:

```bash
PATH="../../bin:$PATH" pulumi preview --stack nginx-hello
PATH="../../bin:$PATH" pulumi up --stack nginx-hello
```

The example is intentionally opt-in. Do not run `pulumi up` unless you want to
create and deploy the demonstration application in the configured Coolify
project.

## Resource Selection

The provider exposes Coolify resources for applications, databases, services,
projects, servers, storage, environment variables, deployments, backups, and
data sources for reading existing resources.

Use data sources such as `GetProject` and `GetServer` when adopting existing
Coolify objects. Use resource constructors such as
`NewApplicationDockerImage` and `NewDeployment` when Pulumi should manage
their lifecycle.

## Configuration Safety

- Store API tokens with `pulumi config set --secret`.
- Do not commit populated `Pulumi.*.yaml` files.
- Prefer `caCert` for self-signed or internal certificate authorities.
- Use `insecure` only for temporary development diagnostics.
- Review the preview before applying changes to a Coolify project.
