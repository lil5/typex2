package rust

import (
	"fmt"
	"go/types"
	"regexp"
	"sort"
	"strings"

	"github.com/lil5/typex2/internal/generate"
	"github.com/lil5/typex2/internal/utils"
)

func GenerateRust(tm *utils.StructMap) (*strings.Builder, error) {
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
	s.WriteString("use serde::{Serialize, Deserialize};\n\n")
	for _, n := range keys {
		ac := &ArgCounter{}
		t := (*tm)[n]
		indent = 0
		gp := publicPrefix(n)
		s.WriteString("#[derive(Serialize, Deserialize)]\n")
		switch tt := t.(type) {
		case *types.Struct:
			// generate type content
			gc := getClassFields(ac, tt, indent)

			// generate class
			dep, _ := generate.GetClassDeps(tt, false)
			gc1, gc2 := buildClass(ac, n, dep)
			s.WriteString(gp + gc1 + gc + gc2)
		default:
			// // generate type content
			gc := getTypeContent(ac, tt)
			// // generate type alias
			gt1, gt2 := buildTypeAlias(ac, n)

			s.WriteString(gp + gt1 + gc + gt2)
		}
	}

	return &s, nil
}

var reFirstLetterUpper = regexp.MustCompile("^[A-Z]")

func publicPrefix(name string) string {
	if reFirstLetterUpper.MatchString(name) {
		return "pub "
	}
	return "pub mod "
}
