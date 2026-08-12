package shim

import (
	"github.com/coolify-terraform/terraform-provider-coolify/internal/provider"
	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
)

// New exposes the upstream framework provider through an allowed internal
// import path. The bridge module consumes this small shim module.
func New(version string) func() frameworkprovider.Provider {
	return provider.New(version)
}
