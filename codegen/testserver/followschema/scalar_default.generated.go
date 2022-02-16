// Code generated by github.com/stealthmodesoft/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"strconv"

	"github.com/stealthmodesoft/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _EmbeddedDefaultScalar_value(ctx context.Context, field graphql.CollectedField, obj *EmbeddedDefaultScalar) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "EmbeddedDefaultScalar",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Value, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalODefaultScalarImplementation2ᚖstring(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var embeddedDefaultScalarImplementors = []string{"EmbeddedDefaultScalar"}

func (ec *executionContext) _EmbeddedDefaultScalar(ctx context.Context, sel ast.SelectionSet, obj *EmbeddedDefaultScalar) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, embeddedDefaultScalarImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EmbeddedDefaultScalar")
		case "value":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._EmbeddedDefaultScalar_value(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNDefaultScalarImplementation2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNDefaultScalarImplementation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalODefaultScalarImplementation2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalString(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalODefaultScalarImplementation2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalString(*v)
	return res
}

// endregion ***************************** type.gotpl *****************************
