package dart

import (
	"go/types"

	"github.com/lil5/typex2/internal/generate"
)

type Names = struct {
	class string
	key   string
}

func getTypeContent(t types.Type, names *Names) string {
	var s string
	switch tt := t.(type) {
	case *types.Chan, *types.Signature:
		s = "dynamic"
	case *types.Interface:
		s = "Map<String, dynamic>"
	case *types.Basic:
		s = getBasicType(tt)
	case *types.Array:
		s = getArrayType(tt, names)
	case *types.Slice:
		s = getSliceType(tt, names)
	case *types.Map:
		s = getMapType(tt, names)
	case *types.Named:
		s = getNamedType(tt)
	case *types.Pointer:
		s = getTypeContent(tt.Elem(), names)
		s += "?"
	//case *types.Struct:
	// recursivily create classes
	// s will return the new class name
	// s = getStructType(tt)
	default:
		s = "dynamic"
	}

	return s
}

func getMapType(t *types.Map, names *Names) string {
	s := "Map<"

	s += getMapKey(t.Key())
	s += ", "
	s += getTypeContent(t.Elem(), names)
	s += ">"

	return s
}

func getArrayType(t *types.Array, names *Names) string {
	s := "List<"
	s += getTypeContent(t.Elem(), names)
	s += ">"

	return s
}

func getSliceType(t *types.Slice, names *Names) string {
	s := "List<"
	s += getTypeContent(t.Elem(), names)
	s += ">"

	return s
}

func getNamedType(t *types.Named) string {
	return generate.GetName(t.String())
}

func getBasicType(t *types.Basic) string {
	switch t.Kind() {
	case types.Bool:
		return "bool"
	case types.Int,
		types.Int8,
		types.Int16,
		types.Int32,
		types.Int64,
		types.Uint,
		types.Uint8,
		types.Uint16,
		types.Uint32,
		types.Uint64:
		return "int"
	case types.Float32,
		types.Float64,
		types.Complex64,
		types.Complex128:
		return "double"
	case types.String:
		return "String"
	default:
		return "dynamic"
	}
}

func getMapKey(t types.Type) string {
	switch tt := t.(type) {
	case *types.Basic:
		switch tt.Kind() {
		case types.Int,
			types.Int8,
			types.Int16,
			types.Int32,
			types.Int64,
			types.Uint,
			types.Uint8,
			types.Uint16,
			types.Uint32,
			types.Uint64:
			return "int"
		case types.Float32,
			types.Float64,
			types.Complex64,
			types.Complex128:
			return "double"
		default:
			return "String"
		}
	case *types.Map:
		return getMapKey(tt.Key())
	case *types.Named:
		return getMapKey(tt.Underlying())
	}
	return "String"
}
