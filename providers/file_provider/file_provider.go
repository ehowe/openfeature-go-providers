package file_provider

import (
	"context"
	"errors"

	"github.com/open-feature/go-sdk/openfeature"
)

// FileProvider implements the FeatureProvider interface and provides functions for evaluating flags
type FileProvider struct {
	file file
}

const (
	ReasonStatic = "static"

	ErrorTypeMismatch = "type mismatch"
	ErrorParse        = "parse error"
	ErrorFlagNotFound = "flag not found"
)

func NewFileProvider(path string, format string, deepKeys []string) FileProvider {
	return FileProvider{
		file: file{
			path:     path,
			format:   format,
			deepKeys: deepKeys,
		},
	}
}

// Metadata returns the metadata of the provider
func (p FileProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "file-flag-evaluator",
	}
}

// Hooks returns hooks
func (p FileProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}

// BooleanEvaluation returns a boolean flag
func (p FileProvider) BooleanEvaluation(ctx context.Context, flagKey string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	res := p.resolveFlag(flagKey, defaultValue, evalCtx)
	v, ok := res.Value.(bool)
	if !ok {
		return openfeature.BoolResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewTypeMismatchResolutionError(""),
				Reason:          openfeature.ErrorReason,
			},
		}
	}

	return openfeature.BoolResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: res.ProviderResolutionDetail,
	}
}

// StringEvaluation returns a string flag
func (p FileProvider) StringEvaluation(ctx context.Context, flagKey string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	res := p.resolveFlag(flagKey, defaultValue, evalCtx)
	v, ok := res.Value.(string)
	if !ok {
		return openfeature.StringResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewTypeMismatchResolutionError(""),
				Reason:          openfeature.ErrorReason,
			},
		}
	}

	return openfeature.StringResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: res.ProviderResolutionDetail,
	}
}

// IntEvaluation returns an int flag
func (p FileProvider) IntEvaluation(ctx context.Context, flagKey string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	res := p.resolveFlag(flagKey, defaultValue, evalCtx)
	v, ok := res.Value.(float64)
	if !ok {
		return openfeature.IntResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewTypeMismatchResolutionError(""),
				Reason:          openfeature.ErrorReason,
			},
		}
	}

	return openfeature.IntResolutionDetail{
		Value:                    int64(v),
		ProviderResolutionDetail: res.ProviderResolutionDetail,
	}
}

// FloatEvaluation returns a float flag
func (p FileProvider) FloatEvaluation(ctx context.Context, flagKey string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	res := p.resolveFlag(flagKey, defaultValue, evalCtx)
	v, ok := res.Value.(float64)
	if !ok {
		return openfeature.FloatResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewTypeMismatchResolutionError(""),
				Reason:          openfeature.ErrorReason,
			},
		}
	}

	return openfeature.FloatResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: res.ProviderResolutionDetail,
	}
}

// ObjectEvaluation returns an object flag
func (p FileProvider) ObjectEvaluation(ctx context.Context, flagKey string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return p.resolveFlag(flagKey, defaultValue, evalCtx)
}

func (p FileProvider) resolveFlag(flagKey string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	// fetch the stored flag from environment variables
	res, err := p.file.fetchFlag(flagKey)
	if err != nil {
		var e openfeature.ResolutionError
		if !errors.As(err, &e) {
			e = openfeature.NewGeneralResolutionError(err.Error())
		}

		return openfeature.InterfaceResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: e,
				Reason:          openfeature.ErrorReason,
			},
		}
	}
	// evaluate the stored flag to return the variant, reason, value and error
	variant, reason, value, err := res.Evaluate(evalCtx)
	if err != nil {
		var e openfeature.ResolutionError
		if !errors.As(err, &e) {
			e = openfeature.NewGeneralResolutionError(err.Error())
		}
		return openfeature.InterfaceResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: e,
				Reason:          openfeature.ErrorReason,
			},
		}
	}

	// return the type naive ResolutionDetail structure
	return openfeature.InterfaceResolutionDetail{
		Value: value,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Variant: variant,
			Reason:  reason,
		},
	}
}
