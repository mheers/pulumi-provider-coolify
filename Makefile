PROVIDER_VERSION ?= 0.1.0-dev
PROVIDER_MODULE := github.com/mheers/pulumi-provider-coolify/provider

.PHONY: provider schema sdk-go build test plugin-install clean

provider:
	mkdir -p bin
	(cd provider && go build -ldflags "-X $(PROVIDER_MODULE)/pkg/version.Version=$(PROVIDER_VERSION)" -o ../bin/pulumi-resource-coolify ./cmd/pulumi-resource-coolify)

schema:
	mkdir -p bin
	(cd provider && go build -o ../bin/pulumi-tfgen-coolify ./cmd/pulumi-tfgen-coolify)
	./bin/pulumi-tfgen-coolify schema --out provider/cmd/pulumi-resource-coolify

sdk-go: schema
	pulumi package gen-sdk provider/cmd/pulumi-resource-coolify/schema.json --language go --out sdk --local --version $(PROVIDER_VERSION)
	(cd sdk && go mod tidy)

build: provider sdk-go

plugin-install: provider
	pulumi plugin install resource coolify $(PROVIDER_VERSION) --file bin/pulumi-resource-coolify

test:
	(cd provider && go test ./...)
	(cd provider/shim && go test ./...)
	(cd sdk && go test ./...)
	(cd examples/nginx-hello && go test ./...)

clean:
	rm -rf bin .make .pulumi
