# Contributing to the Pulumi ecosystem

Do you want to contribute to Pulumi? Awesome! We are so happy to have you.
We have a few tips and housekeeping items to help you get up and running.

## Code of Conduct

Please make sure to read and observe our [Code of Conduct](./CODE-OF-CONDUCT.md)

## Community Expectations

Please read about our [contribution guidelines here.](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md#communications)

## Setting up your development environment

### Pulumi prerequisites

Please refer to the [main Pulumi repo](https://github.com/pulumi/pulumi/)'s [CONTRIBUTING.md file](
https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md#developing) for details on how to get set up with Pulumi.

## Committing Generated Code

You must generate and check in the SDKs on each pull request containing a code change, e.g. adding a new resource to `resources.go`.

1. Run `make build` from the root of this repository
1. Open a pull request containing all changes
1. *Note:* If a large number of seemingly-unrelated diffs are produced by `make build` (for example, lots of changes to comments unrelated to the change you are making), ensure that the latest dependencies for the provider are installed by running `go mod tidy` in the `provider/` directory of this repository.

## Running Integration Tests

The default `make test` command runs Go tests and does not contact a Coolify
server. The `examples/nginx-hello` program is an opt-in integration example;
running `pulumi up` creates a real application and deployment in the configured
Coolify project. Always review `pulumi preview` first.
