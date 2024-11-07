package kotlin

import (
	"fmt"
	"go/types"

	"github.com/lil5/typex2/internal/generate"
)

func getClassFields(t *types.Struct, indent int) string {
	if t.NumFields() == 0 {
		return ""
	}
	var s string
	indent += 1
	for i := 0; i < t.NumFields(); i++ {
		s = generate.IndentStr(indent)
		field := t.Field(i)
		s += "val " + field.Name() + ": " + getTypeContent(field.Type()) + ",\n"
	}
	return s
}

func getTypeContent(t types.Type) string {
	fmt.Printf("%v\n\n", t.String())
	var s string
	switch tt := t.(type) {
	case *types.Chan, *types.Signature:
		s = "Any"
	case *types.Interface:
		s = "Map<String, Any>"
	case *types.Basic:
		s = getBasicType(tt)
	case *types.Array:
		s = getArrayType(tt)
	case *types.Slice:
		s = getSliceType(tt)
	case *types.Map:
		s = getMapType(tt)
	case *types.Named:
		s = getNamedType(tt)
	case *types.Pointer:
		s = getTypeContent(tt.Elem())
		s += "?"
	//case *types.Struct:
	// recursivily create classes
	// s will return the new class name
	// s = getStructType(tt)
	default:
		s = "Any"
	}

	return s
}

func getMapType(t *types.Map) string {
	s := "Map<"

	s += getMapKey(t.Key())
	s += ", "
	s += getTypeContent(t.Elem())
	s += ">"

	return s
}

func getArrayType(t *types.Array) string {
	s := "Array<"
	s += getTypeContent(t.Elem())
	s += ">"

	return s
}

func getSliceType(t *types.Slice) string {
	s := "Array<"
	s += getTypeContent(t.Elem())
	s += ">"

	return s
}

func getNamedType(t *types.Named) string {
	return generate.GetName(t.String())
}

func getBasicType(t *types.Basic) string {
	fmt.Printf("t: %v\n", t)
	switch t.Kind() {
	case types.Bool:
		return "Boolean"
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
		return "Int"
	case types.Float32:
		return "Float"
	case types.Float64,
		types.Complex64,
		types.Complex128:
		return "Double"
	case types.String:
		return "String"
	default:
		return "Any"
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
			return "Int"
		case types.Float32:
			return "Float"
		case types.Float64,
			types.Complex64,
			types.Complex128:
			return "Double"
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
