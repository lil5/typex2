package read

import (
	"errors"

	"golang.org/x/tools/go/packages"
)

func LoadPackages(path string) ([]*packages.Package, error) {

	c := packages.Config{
		Mode:  packages.NeedTypes | packages.NeedTypesInfo,
		Tests: false,
	}

	pkgs, err := packages.Load(&c, path)

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			err := errors.New(pkg.Errors[0].Msg)
			return nil, err
		}
	}

	return pkgs, err
}
