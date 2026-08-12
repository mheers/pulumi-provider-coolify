// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package coolify

import (
	_ "embed"
	"path"

	"github.com/coolify-terraform/terraform-provider-coolify/shim"
	"github.com/mheers/pulumi-provider-coolify/provider/pkg/version"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

const (
	mainPkg      = "coolify"
	mainMod      = "index"
	upstreamRepo = "github.com/coolify-terraform/terraform-provider-coolify"
	providerRepo = "https://github.com/mheers/pulumi-provider-coolify"
)

//go:embed cmd/pulumi-resource-coolify/bridge-metadata.json
var metadata []byte

// Provider returns the bridge metadata and upstream framework provider.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:       pfbridge.ShimProvider(shim.New(version.Version)()),
		Name:    mainPkg,
		Version: version.Version,

		DisplayName:       "Coolify",
		Publisher:         "mheers",
		Description:       "A Pulumi provider for managing Coolify resources.",
		PluginDownloadURL: "github://api.github.com/mheers/pulumi-provider-coolify",
		Keywords:          []string{"coolify", "paas", "category/cloud"},
		License:           "MPL-2.0",
		Homepage:          "https://coolify.io",
		Repository:        providerRepo,
		GitHubOrg:         "coolify-terraform",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"coolify_environment": {
				Fields: map[string]*tfbridge.SchemaInfo{"id": {Type: "string"}},
			},
			"coolify_application":              uuidResourceInfo(),
			"coolify_application_dockerfile":   uuidResourceInfo(),
			"coolify_application_docker_image": uuidResourceInfo(),
			"coolify_application_github_app":   uuidResourceInfo(),
			"coolify_application_private_git":  uuidResourceInfo(),
			"coolify_cloud_token":              uuidResourceInfo(),
			"coolify_database_backup":          uuidResourceInfo(),
			"coolify_database_clickhouse":      uuidResourceInfo(),
			"coolify_database_dragonfly":       uuidResourceInfo(),
			"coolify_database_keydb":           uuidResourceInfo(),
			"coolify_database_mariadb":         uuidResourceInfo(),
			"coolify_database_mongodb":         uuidResourceInfo(),
			"coolify_database_mysql":           uuidResourceInfo(),
			"coolify_database_postgresql":      uuidResourceInfo(),
			"coolify_database_redis":           uuidResourceInfo(),
			"coolify_deployment":               uuidResourceInfo(),
			"coolify_destination":              uuidResourceInfo(),
			"coolify_environment_variable":     uuidResourceInfo(),
			"coolify_github_app":               uuidResourceInfo(),
			"coolify_private_key":              uuidResourceInfo(),
			"coolify_project":                  uuidResourceInfo(),
			"coolify_scheduled_task":           uuidResourceInfo(),
			"coolify_server":                   uuidResourceInfo(),
			"coolify_server_digitalocean":      uuidResourceInfo(),
			"coolify_server_hetzner":           uuidResourceInfo(),
			"coolify_server_vultr":             uuidResourceInfo(),
			"coolify_service":                  uuidResourceInfo(),
			"coolify_storage":                  uuidResourceInfo(),
			"coolify_storage_backup":           uuidResourceInfo(),
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,
			PyProject:            struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/mheers/pulumi-provider-coolify/sdk",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			RespectSchemaVersion:           true,
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("coolify_", mainMod, tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}

func uuidResourceInfo() *tfbridge.ResourceInfo {
	return &tfbridge.ResourceInfo{
		ComputeID: tfbridge.DelegateIDField("uuid", "Coolify", upstreamRepo),
	}
}
