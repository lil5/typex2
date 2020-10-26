package read

import (
	"github.com/lil5/typex2/internal/utils"
	"golang.org/x/tools/go/packages"
)

// filter packages that are not exported
func MapPackage(pkg *packages.Package) *utils.StructMap {
	scope := pkg.Types.Scope()

	sm := make(utils.StructMap)

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		// This is how to get the go path + name
		// path := obj.Pkg().Path() + "." + name

		t := obj.Type()

		sm[name] = t.Underlying()

		// fmt.Printf("type: %s, path: %s\n", t.String(), path)
	}

	return &sm
}
