# nginx-hello

Minimal end-to-end source example for the Coolify provider. It adopts an
existing Coolify server and project, creates an `nginx:latest` Docker image
application, and waits for a deployment to finish when run explicitly.

The example supports both credential modes:

- Host-stack mode reads the endpoint and token from an existing Pulumi stack.
- Direct mode uses the `nginx:endpoint` and secret `nginx:token` config values.

This example is not deployed by the provider build or test commands. No live
Pulumi stack configuration is committed. Running `pulumi up` is an explicit
opt-in operation against the Coolify project configured below.

```bash
pulumi stack init nginx-hello
# Host-stack mode:
cp Pulumi.nginx-hello.example.yaml Pulumi.nginx-hello.yaml
# Edit Pulumi.nginx-hello.yaml with your own values.
# Or direct mode:
cp Pulumi.nginx-hello.direct.example.yaml Pulumi.nginx-hello.yaml
# Edit the endpoint and Coolify identifiers, then set the token securely:
pulumi config set --secret nginx:token your-coolify-api-token
```

From the provider repository root, build the provider first:

```bash
make build
cd examples/nginx-hello
PATH="../../bin:$PATH" pulumi preview --stack nginx-hello
PATH="../../bin:$PATH" pulumi up --stack nginx-hello
```
