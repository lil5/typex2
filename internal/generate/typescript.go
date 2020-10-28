package generate

import (
	"fmt"
	"go/token"
	"go/types"
	"sort"
	"strings"

	"github.com/lil5/typex2/internal/utils"
)

func GenerateTypescript(tm *utils.StructMap) (*string, error) {
	if tm == nil {
		return nil, fmt.Errorf("tm pointer is nil")
	}

	// sort map to help any vcs
	keys := make([]string, 0, len(*tm))
	for k := range *tm {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var indent int
	var s string
	for _, n := range keys {
		t := (*tm)[n]
		indent = 1
		if token.IsExported(n) {
			s += "export "
		}

		switch tt := t.(type) {
		case *types.Struct:
			// generate type content
			gc := getStructFields(tt, &indent)

			// generate interface
			deps := getStructDeps(tt)
			gi1, gi2 := buildInterface(n, deps)
			s += gi1 + gc + gi2
		default:
			// generate type content
			gc := getTypeContent(tt, &indent)
			// generate type alias
			gt1, gt2 := buildTypeAlias(n)

			s += gt1 + gc + gt2
		}
	}

	return &s, nil
}

// run on type struct
func buildInterface(name string, deps *[]string) (string, string) {
	s := "interface " + name + " "

	if deps != nil && len(*deps) > 0 {
		s += "extends "
		s += strings.Join(*deps, ", ")
		s += " "
	}

	s += "{\n"
	s2 := "}\n\n"

	return s, s2
}

// TODO: add pointer check and run basic type or interface

func buildTypeAlias(name string) (string, string) {
	s1 := "type " + name + " = "

	s2 := "\n\n"
	return s1, s2
}

func getTypeContent(t types.Type, indent *int) string {
	var s string
	switch tt := t.(type) {
	case *types.Interface, *types.Chan, *types.Signature:
		s = "any"
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
	s += indentStr(indent)

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
			name := getStructTagJSON(t, i)

			s += indentStr(indent)
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
	return getName(t.String())
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
		fmt.Printf("basic t: %v\n", t.Info())
		return "ANY"
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
