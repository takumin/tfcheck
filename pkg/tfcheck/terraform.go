package tfcheck

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty/gocty"
)

func (v *version) parseTerraformVersion(body *hclwrite.Body) (string, error) {
	var res string

	if attr := body.GetAttribute("required_version"); attr != nil {
		val, err := getAttributeValue(attr)
		if err != nil {
			return "", err
		}

		if err := gocty.FromCtyValue(val, &res); err != nil {
			return "", fmt.Errorf("Failed to gocty.FromCtyValue(): %w", err)
		}
	}

	return res, nil
}
