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

func (ec *executionContext) _PtrToSliceContainer_ptrToSlice(ctx context.Context, field graphql.CollectedField, obj *PtrToSliceContainer) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "PtrToSliceContainer",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PtrToSlice, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*[]string)
	fc.Result = res
	return ec.marshalOString2ᚖᚕstringᚄ(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var ptrToSliceContainerImplementors = []string{"PtrToSliceContainer"}

func (ec *executionContext) _PtrToSliceContainer(ctx context.Context, sel ast.SelectionSet, obj *PtrToSliceContainer) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, ptrToSliceContainerImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PtrToSliceContainer")
		case "ptrToSlice":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._PtrToSliceContainer_ptrToSlice(ctx, field, obj)
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

func (ec *executionContext) marshalNPtrToSliceContainer2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToSliceContainer(ctx context.Context, sel ast.SelectionSet, v PtrToSliceContainer) graphql.Marshaler {
	return ec._PtrToSliceContainer(ctx, sel, &v)
}

func (ec *executionContext) marshalNPtrToSliceContainer2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToSliceContainer(ctx context.Context, sel ast.SelectionSet, v *PtrToSliceContainer) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._PtrToSliceContainer(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
