package tfcheck

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty/gocty"
)

func (v *version) parseTerraformModules(body *hclwrite.Body, label string) (map[string]Module, error) {
	res := make(map[string]Module)

	source, err := getTerraformModulesAttribute(body, "source")
	if err != nil {
		return nil, fmt.Errorf("Failed to getAttributeValue(): %w", err)
	}

	version, err := getTerraformModulesAttribute(body, "version")
	if err != nil {
		return nil, fmt.Errorf("Failed to getAttributeValue(): %w", err)
	}

	if source != "" && version != "" {
		res[label] = Module{
			Source:  source,
			Version: version,
		}
	}

	return res, nil
}

func getTerraformModulesAttribute(body *hclwrite.Body, key string) (string, error) {
	var res string

	if attr := body.GetAttribute(key); attr != nil {
		v, err := getAttributeValue(attr)
		if err != nil {
			return "", err
		}

		if err := gocty.FromCtyValue(v, &res); err != nil {
			return "", fmt.Errorf("Failed to gocty.FromCtyValue(): %w", err)
		}
	}

	return res, nil
}
