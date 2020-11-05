package dart

import (
	"fmt"
	"go/types"
	"sort"

	"github.com/lil5/typex2/internal/generate"
	"github.com/lil5/typex2/internal/utils"
)

func GenerateDart(tm *utils.StructMap) (*string, error) {
	if tm == nil {
		return nil, fmt.Errorf("tm pointer is nil")
	}

	// sort map to help any vcs
	keys := make([]string, 0, len(*tm))
	for k := range *tm {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var s string
	for _, n := range keys {
		t := (*tm)[n]
		switch tt := t.(type) {
		case *types.Struct:
			// generate type content
			gc := ""
			//gc := getStructFields(tt, &indent)

			// generate interface
			dep, _ := generate.GetClassDeps(tt, false)
			gc1, gc2 := buildClass(n, dep)
			s += gc1 + gc + gc2
		default:
			// // generate type content
			// gc := getTypeContent(tt, &indent)
			// // generate type alias
			// gt1, gt2 := buildTypeAlias(n)

			// s += gt1 + gc + gt2
		}
	}

	return &s, nil
}