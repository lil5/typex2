package generate

import (
	"go/types"
	"reflect"
	"strings"
)

// Check if json has a type for t
// and if it is worth getting the named type
func IsTranslatable(t types.Type) bool {
	switch t.(type) {
	case *types.Chan, *types.Interface, *types.Signature:
		return false
	case *types.Named:
		return IsTranslatable(t.Underlying())
	}
	return true
}

func GetStructDeps(t *types.Struct) *[]string {
	var deps []string
	for i := 0; i < t.NumFields(); i++ {
		fld := (*t).Field(i)
		if fld.Anonymous() {

			deps = append(deps, fld.Name())
		}
	}

	return &deps
}

func GetName(p string) string {
	i := strings.LastIndex(p, ".")
	return p[i+1:]
}

func IndentStr(indent int) string {
	return strings.Repeat("\t", indent)
}

func GetStructTagJSON(t *types.Struct, i int) string {
	f := t.Field(i)
	tag := f.Name()

	st := reflect.StructTag(t.Tag(i))
	tagJSON, ok := st.Lookup("json")
	if ok {
		tag = tagJSON
	}

	return tag
}
