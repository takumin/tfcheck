package tfcheck

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Parser interface {
	Parse() (Versions, error)
}

type Versions map[string]Version

type Version struct {
	IsSet     bool
	Terraform string
	Providers map[string]Provider
	Modules   map[string]Module
}

type Provider struct {
	Source  string
	Version string
}

type Module struct {
	Source  string
	Version string
}

type version struct {
	filepaths []string
	versions  Versions
}

func NewParser(filepaths []string) Parser {
	return &version{
		filepaths: filepaths,
		versions:  make(map[string]Version, len(filepaths)),
	}
}

func (v *version) Parse() (Versions, error) {
	for _, filepath := range v.filepaths {
		v.versions[filepath] = Version{}

		raw, err := ioutil.ReadFile(filepath)
		if err != nil {
			return nil, fmt.Errorf("Falied to ioutil.ReadFile(): %#w", err)
		}

		hclfile, diags := hclwrite.ParseConfig(raw, filepath, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			return nil, fmt.Errorf("Falied to hclwrite.ParseConfig(): %#w", diags)
		}

		if err := v.parseTerraform(filepath, hclfile.Body()); err != nil {
			return nil, fmt.Errorf("Falied to parse(): %#w", err)
		}
	}

	return v.versions, nil
}

func (v *version) parseTerraform(filepath string, body *hclwrite.Body) error {
	ver := Version{}

	for _, block := range body.Blocks() {
		if block.Type() == "terraform" {
			ver.IsSet = true

			requiredTerraform, err := v.parseTerraformVersion(block.Body())
			if err != nil {
				return err
			}
			ver.Terraform = requiredTerraform

			requiredProviders, err := v.parseTerraformProviders(block.Body())
			if err != nil {
				return err
			}
			ver.Providers = requiredProviders
		}

		if block.Type() == "module" {
			if ver.Modules == nil {
				ver.Modules = make(map[string]Module)
			}

			requiredModules, err := v.parseTerraformModules(block.Body(), block.Labels()[0])
			if err != nil {
				return err
			}

			for mk, mv := range requiredModules {
				ver.IsSet = true
				ver.Modules[mk] = mv
			}
		}
	}

	if ver.IsSet {
		v.versions[filepath] = ver
	}

	return nil
}
