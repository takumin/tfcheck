package tfcheck

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func getAttributeValue(attr *hclwrite.Attribute) (cty.Value, error) {
	src := attr.Expr().BuildTokens(nil).Bytes()

	expr, diags := hclsyntax.ParseExpression(src, "generated_by_attributeToValue", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return cty.NilVal, fmt.Errorf("Failed to hclsyntax.ParseExpression(): %w", diags)
	}

	v, diags := expr.Value(&hcl.EvalContext{})
	if diags.HasErrors() {
		return cty.NilVal, fmt.Errorf("Failed to expr.Value(): %w", diags)
	}

	return v, nil
}
