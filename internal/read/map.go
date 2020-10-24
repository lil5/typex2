package read

import (
	"fmt"
	"go/token"
	"strings"

	"golang.org/x/tools/go/packages"
)

func MapPackages(pkgs []*packages.Package) []*StructMap {
	var sm []*StructMap

	for _, p := range pkgs {
		sm = append(sm, mapPackage(p))
	}

	return sm
}

// filter packages that are not exported
func mapPackage(pkg *packages.Package) *StructMap {
	scope := pkg.Types.Scope()

	sm := make(StructMap)

	for _, name := range scope.Names() {
		if !isExported(name) {
			fmt.Println("isExported: false")
			continue
		}
		obj := scope.Lookup(name)
		path := obj.Pkg().Path() + "." + name

		t := obj.Type()

		sm[path] = t.Underlying()

		// fmt.Printf("type: %s, path: %s\n", t.String(), path)
	}

	return &sm
}

func isExported(s string) bool {
	n := strings.ReplaceAll(s, ".", "/")
	i := strings.LastIndex(n, "/")
	return token.IsExported(n[i+1:])
}
