package singlefile

import "github.com/stealthmodesoft/gqlgen/codegen/testserver/singlefile/otherpkg"

type (
	WrappedScalar = otherpkg.Scalar
	WrappedStruct otherpkg.Struct
	WrappedMap    otherpkg.Map
	WrappedSlice  otherpkg.Slice
)
