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

func GetStructTagJSON(t *types.Struct, i int) (tag string, omitempty bool) {
	f := t.Field(i)
	tagGo := f.Name()

	st := reflect.StructTag(t.Tag(i))
	tag, omitempty, found := readTag(st, "json")
	if !found {
		tag, omitempty, found = readTag(st, "form")
		if !found {
			tag, omitempty, _ = readTag(st, "uri")
		}
	}
	if tag == "" {
		tag = tagGo
	}

	return tag, omitempty
}

func readTag(st reflect.StructTag, key string) (tag string, omitempty, found bool) {
	var tagContents string
	tagContents, found = st.Lookup(key)
	if !found {
		return "", false, false
	}
	tagSplit := strings.Split(tagContents, ",")
	omitempty = strings.Contains(tagContents, "omitempty")
	tag = tagSplit[0]
	return
}

func TurnTypeOptional(t types.Type) types.Type {
	switch t.(type) {
	case *types.Pointer:
		return t
	default:
		return types.NewPointer(t)
	}
}
