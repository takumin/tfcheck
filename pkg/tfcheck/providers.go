package tfcheck

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func (v *version) parseTerraformProviders(body *hclwrite.Body) (map[string]Provider, error) {
	res := make(map[string]Provider)

	for _, block := range body.Blocks() {
		if block.Type() != "required_providers" {
			continue
		}

		for key, attr := range block.Body().Attributes() {
			provider := Provider{}

			value, err := getAttributeValue(attr)
			if err != nil {
				return nil, err
			}

			if value.Type().IsObjectType() {
				m := value.AsValueMap()

				for objk, objv := range m {
					switch objk {
					case "source":
						provider.Source = objv.AsString()
					case "version":
						provider.Version = objv.AsString()
					default:
						continue
					}
				}
			}

			res[key] = provider
		}
	}

	return res, nil
}
