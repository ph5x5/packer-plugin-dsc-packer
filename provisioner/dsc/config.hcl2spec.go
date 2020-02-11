// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package dsc

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName          *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType        *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug              *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce              *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError            *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars           map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars      []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	ExecuteCommand           *string           `mapstructure:"execute_command" cty:"execute_command"`
	ConfigurationParams      map[string]string `mapstructure:"configuration_params" cty:"configuration_params"`
	MofPath                  *string           `mapstructure:"mof_path" cty:"mof_path"`
	ConfigurationFilePath    *string           `mapstructure:"configuration_file" cty:"configuration_file"`
	ManifestDir              *string           `mapstructure:"manifest_dir" cty:"manifest_dir"`
	ManifestFile             *string           `mapstructure:"manifest_file" cty:"manifest_file"`
	ConfigurationName        *string           `mapstructure:"configuration_name" cty:"configuration_name"`
	ModulePaths              []string          `mapstructure:"module_paths" cty:"module_paths"`
	ResourcePaths            []string          `mapstructure:"resource_paths" cty:"resource_paths"`
	InstallPackageManagement *bool             `mapstructure:"install_package_management" cty:"install_package_management"`
	InstallModules           map[string]string `mapstructure:"install_modules" cty:"install_modules"`
	StagingDir               *string           `mapstructure:"staging_dir" cty:"staging_dir"`
	CleanStagingDir          *bool             `mapstructure:"clean_staging_dir" cty:"clean_staging_dir"`
	WorkingDir               *string           `mapstructure:"working_dir" cty:"working_dir"`
	IgnoreExitCodes          *bool             `mapstructure:"ignore_exit_codes" cty:"ignore_exit_codes"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"execute_command":            &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"configuration_params":       &hcldec.BlockAttrsSpec{TypeName: "configuration_params", ElementType: cty.String, Required: false},
		"mof_path":                   &hcldec.AttrSpec{Name: "mof_path", Type: cty.String, Required: false},
		"configuration_file":         &hcldec.AttrSpec{Name: "configuration_file", Type: cty.String, Required: false},
		"manifest_dir":               &hcldec.AttrSpec{Name: "manifest_dir", Type: cty.String, Required: false},
		"manifest_file":              &hcldec.AttrSpec{Name: "manifest_file", Type: cty.String, Required: false},
		"configuration_name":         &hcldec.AttrSpec{Name: "configuration_name", Type: cty.String, Required: false},
		"module_paths":               &hcldec.AttrSpec{Name: "module_paths", Type: cty.List(cty.String), Required: false},
		"resource_paths":             &hcldec.AttrSpec{Name: "resource_paths", Type: cty.List(cty.String), Required: false},
		"install_package_management": &hcldec.AttrSpec{Name: "install_package_management", Type: cty.Bool, Required: false},
		"install_modules":            &hcldec.BlockAttrsSpec{TypeName: "install_modules", ElementType: cty.String, Required: false},
		"staging_dir":                &hcldec.AttrSpec{Name: "staging_dir", Type: cty.String, Required: false},
		"clean_staging_dir":          &hcldec.AttrSpec{Name: "clean_staging_dir", Type: cty.Bool, Required: false},
		"working_dir":                &hcldec.AttrSpec{Name: "working_dir", Type: cty.String, Required: false},
		"ignore_exit_codes":          &hcldec.AttrSpec{Name: "ignore_exit_codes", Type: cty.Bool, Required: false},
	}
	return s
}