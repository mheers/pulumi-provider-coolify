# nginx-hello

Minimal end-to-end source example for the Coolify provider. It adopts an
existing Coolify server and project, creates an `nginx:latest` Docker image
application, and waits for a deployment to finish when run explicitly.

This example is not deployed by the provider build or test commands. No live
Pulumi stack configuration is committed. Running `pulumi up` is an explicit
opt-in operation against the Coolify project configured below.

```bash
pulumi stack init nginx-hello
cp Pulumi.nginx-hello.example.yaml Pulumi.nginx-hello.yaml
# Edit Pulumi.nginx-hello.yaml with your own values.
```

From the provider repository root, build the provider first:

```bash
make build
cd examples/nginx-hello
PATH="../../bin:$PATH" pulumi preview --stack nginx-hello
PATH="../../bin:$PATH" pulumi up --stack nginx-hello
```
