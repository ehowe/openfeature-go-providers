package shared

import (
	"github.com/open-feature/go-sdk/openfeature"
)

type Flag struct {
	Default  string
	Enabled  bool
	Name     string
	Type     string
	Value    interface{}
	Variant  string
	Variants map[string]interface{}
}

func (f *Flag) Evaluate(evalCtx map[string]interface{}) (string, openfeature.Reason, interface{}, error) {
	if !f.Enabled {
		return "", openfeature.DisabledReason, f.Value, openfeature.NewParseErrorResolutionError("flag is disabled")
	}

	if f.Variants == nil {
		return "value", openfeature.TargetingMatchReason, f.Value, nil
	}

	resolvedVariant := f.Variants[f.Variant]

	if resolvedVariant == nil {
		return "", openfeature.Reason(openfeature.TargetingKeyMissingCode), nil, nil
	}

	return f.Variant, openfeature.TargetingMatchReason, resolvedVariant, nil
}
