// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNFallbackToStringEncoding2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFallbackToStringEncoding(ctx context.Context, v interface{}) (FallbackToStringEncoding, error) {
	tmp, err := graphql.UnmarshalString(v)
	res := FallbackToStringEncoding(tmp)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNFallbackToStringEncoding2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFallbackToStringEncoding(ctx context.Context, sel ast.SelectionSet, v FallbackToStringEncoding) graphql.Marshaler {
	res := graphql.MarshalString(string(v))
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

// endregion ***************************** type.gotpl *****************************
