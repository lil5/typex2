package read

import (
	"errors"
	"fmt"

	"golang.org/x/tools/go/packages"
)

func LoadPackage(path string) (*packages.Package, error) {

	c := packages.Config{
		Mode:  packages.NeedTypes | packages.NeedTypesInfo,
		Tests: false,
	}

	pkgs, err := packages.Load(&c, path)
	if err != nil {
		return nil, fmt.Errorf("unable to run package.Load: %w", err)
	}

	if pkgsLen := len(pkgs); pkgsLen != 1 {
		err := fmt.Errorf("path must include only one package\ncurrent amount: %d", pkgsLen)
		return nil, err
	}

	pkg := pkgs[0]

	if len(pkg.Errors) > 0 {
		err := errors.New(pkg.Errors[0].Msg)
		return nil, fmt.Errorf("unable to load package: %w", err)
	}

	return pkg, err
}
