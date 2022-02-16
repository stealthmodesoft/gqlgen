// Code generated by github.com/stealthmodesoft/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"strconv"
	"sync/atomic"

	"github.com/stealthmodesoft/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type OverlappingFieldsResolver interface {
	OldFoo(ctx context.Context, obj *OverlappingFields) (int, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _OverlappingFields_oneFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_twoFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_oldFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.OverlappingFields().OldFoo(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_newFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NewFoo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) _OverlappingFields_new_foo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NewFoo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var overlappingFieldsImplementors = []string{"OverlappingFields"}

func (ec *executionContext) _OverlappingFields(ctx context.Context, sel ast.SelectionSet, obj *OverlappingFields) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, overlappingFieldsImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("OverlappingFields")
		case "oneFoo":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._OverlappingFields_oneFoo(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "twoFoo":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._OverlappingFields_twoFoo(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "oldFoo":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._OverlappingFields_oldFoo(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "newFoo":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._OverlappingFields_newFoo(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "new_foo":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._OverlappingFields_new_foo(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
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

func (ec *executionContext) marshalOOverlappingFields2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐOverlappingFields(ctx context.Context, sel ast.SelectionSet, v *OverlappingFields) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._OverlappingFields(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
