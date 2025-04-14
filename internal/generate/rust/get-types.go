package rust

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/lil5/typex2/internal/generate"
)

func getClassFields(ac *ArgCounter, t *types.Struct, indent int) string {
	if t.NumFields() == 0 {
		return ""
	}
	var s string
	indent += 1
	for i := range t.NumFields() {
		field := t.Field(i)
		tag, omitempty := generate.GetStructTagJSON(t, i)
		if tag == "-" {
			continue
		}
		t := field.Type()
		if omitempty {
			t = generate.TurnTypeOptional(t)
		}
		s += generate.IndentStr(indent)
		s += tag + ": " + getTypeContent(ac, t) + ",\n"
	}
	return s
}

func getTypeContent(ac *ArgCounter, t types.Type) string {
	var s string
	switch tt := t.(type) {
	case *types.Chan, *types.Signature:
		s = ac.append()
	case *types.Interface:
		s = "Map<String, " + ac.append() + ">"
	case *types.Basic:
		s = getBasicType(ac, tt)
	case *types.Array:
		s = getArrayType(ac, tt)
	case *types.Slice:
		s = getSliceType(ac, tt)
	case *types.Map:
		s = getMapType(ac, tt)
	case *types.Named:
		// fmt.Printf("'%v'\n\n", tt.String())
		s = getNamedType(tt)
	case *types.Pointer:
		s += "Option<"
		s += getTypeContent(ac, tt.Elem())
		s += ">"
	//case *types.Struct:
	// recursivily create classes
	// s will return the new class name
	// s = getStructType(tt)
	default:
		s = ac.append()
	}

	return s
}

func getMapType(ac *ArgCounter, t *types.Map) string {
	s := "Map<"
	s += getMapKey(ac, t.Key())
	s += ", "
	s += getTypeContent(ac, t.Elem())
	s += ">"

	return s
}

func getArrayType(ac *ArgCounter, t *types.Array) string {
	s := "Vec<"
	s += getTypeContent(ac, t.Elem())
	s += ">"

	return s
}

func getSliceType(ac *ArgCounter, t *types.Slice) string {
	s := "Vec<"
	s += getTypeContent(ac, t.Elem())
	s += ">"

	return s
}

func getNamedType(t *types.Named) string {
	s := t.String()
	switch s {
	case "time.Time":
		return "String"
	}
	return generate.GetName(s)
}

func getBasicType(ac *ArgCounter, t *types.Basic) string {
	// fmt.Printf("t: %v\n", t)
	switch t.Kind() {
	case types.Bool:
		return "bool"
	case types.Int:
		return "usize"
	case types.Int64:
		return "i64"
	case types.Int8:
		return "i8"
	case types.Int16:
		return "i16"
	case types.Int32:
		return "i32"
	case types.Uint:
		return "usize"
	case types.Uint8:
		return "u8"
	case types.Uint16:
		return "u16"
	case types.Uint32:
		return "u32"
	case types.Uint64:
		return "u64"
	case types.Float32:
		return "f32"
	case types.Float64,
		types.Complex64:
		return "f64"
	case types.Complex128,
		types.String:
		return "String"
	default:
		return ac.append()
	}
}

func getMapKey(ac *ArgCounter, t types.Type) string {
	switch tt := t.(type) {
	case *types.Basic:
		return getBasicType(ac, tt)
	case *types.Map:
		return getMapKey(ac, tt.Key())
	case *types.Named:
		return getMapKey(ac, tt.Underlying())
	}
	return "String"
}

type ArgCounter struct {
	list []string
}

func (ac *ArgCounter) append() string {
	next := fmt.Sprintf("T%d", len(ac.list))
	ac.list = append(ac.list, next)
	return next
}

func (ac *ArgCounter) typeParams() string {
	if len(ac.list) == 0 {
		return ""
	}
	return fmt.Sprintf("<%s>", strings.Join(ac.list, ", "))
}
