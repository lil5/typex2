package typescript

import (
	"fmt"
	"go/types"

	"github.com/lil5/typex2/internal/generate"
)

func getTypeContent(t types.Type, indent *int) string {
	var s string
	switch tt := t.(type) {
	case *types.Chan, *types.Signature:
		s = "unknown"
	case *types.Interface:
		s = "Record<string, any>"
	case *types.Basic:
		s = getBasicType(tt)
	case *types.Array:
		s = getArrayType(tt, indent)
	case *types.Map:
		s = getMapType(tt, indent)
	case *types.Named:
		s = getNamedType(tt)
	case *types.Pointer:
		s = getTypeContent(tt.Elem(), indent)
		s += " | null"
		s = fmt.Sprintf("(%s)", s)
	case *types.Slice:
		s = getSliceType(tt, indent)
	case *types.Struct:
		s = getStructType(tt, indent)
	default:
		s = "unknown"
	}

	return s
}

func getStructType(t *types.Struct, indent *int) string {
	if t.NumFields() == 0 {
		return "{}"
	}
	s := "{\n"
	*indent = *indent + 1
	s += generate.IndentStr(indent)

	for i := 0; i < t.NumFields(); i++ {

	}
	*indent = *indent - 1
	s += "}"
	return s
}

func getStructFields(t *types.Struct, indent *int) string {
	s := ""
	for i := 0; i < t.NumFields(); i++ {
		f := t.Field(i)
		if !f.Anonymous() {
			name := generate.GetStructTagJSON(t, i)

			s += generate.IndentStr(indent)
			s += fmt.Sprintf("%s: ", name)
			s += getTypeContent(f.Type(), indent)
			s += "\n"
		}
	}

	return s
}

func getMapType(t *types.Map, indent *int) string {
	s := "Record<"

	s += getMapKey(t.Key())
	s += ", "
	s += getTypeContent(t.Elem(), indent)
	s += ">"

	return s
}

func getSliceType(t *types.Slice, indent *int) string {
	s := getTypeContent(t.Elem(), indent)
	s += "[]"

	return s
}

func getArrayType(t *types.Array, indent *int) string {
	s := getTypeContent(t.Elem(), indent)
	s += fmt.Sprintf("[/* %d */]", t.Len())

	return s
}

func getNamedType(t *types.Named) string {
	return generate.GetName(t.String())
}

func getBasicType(t *types.Basic) string {
	switch t.Kind() {
	case types.Bool:
		return "boolean"
	case types.Int,
		types.Int8,
		types.Int16,
		types.Int32,
		types.Int64,
		types.Uint,
		types.Uint8,
		types.Uint16,
		types.Uint32,
		types.Uint64,
		types.Float32,
		types.Float64,
		types.Complex64,
		types.Complex128:
		return "number"
	case types.String:
		return "string"
	default:
		return "any"
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
			types.Uint64,
			types.Float32,
			types.Float64,
			types.Complex64,
			types.Complex128:
			return "number"
		default:
			return "string"
		}
	case *types.Map:
		return getMapKey(tt.Key())
	case *types.Named:
		return getMapKey(tt.Underlying())
	}
	return "symbol"
}
