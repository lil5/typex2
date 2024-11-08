package typescript

import (
	"fmt"
	"go/token"
	"go/types"
	"sort"
	"strings"

	"github.com/lil5/typex2/internal/generate"
	"github.com/lil5/typex2/internal/utils"
)

func GenerateTypescript(tm *utils.StructMap) (*strings.Builder, error) {
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
	var s strings.Builder
	for _, n := range keys {
		t := (*tm)[n]
		indent = 1
		if token.IsExported(n) {
			s.WriteString("export ")
		}

		switch tt := t.(type) {
		case *types.Struct:
			// generate type content
			gc := getStructFields(tt, indent)

			// generate interface
			deps := generate.GetStructDeps(tt)
			gi1, gi2 := buildInterface(n, deps)
			s.WriteString(gi1 + gc + gi2)
		default:
			// generate type content
			gc := getTypeContent(tt, indent)
			// generate type alias
			gt1, gt2 := buildTypeAlias(n)

			s.WriteString(gt1 + gc + gt2)
		}
	}

	return &s, nil
}
