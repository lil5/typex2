package generate

import (
	"fmt"
	"go/types"
	"strings"
)

func GetClassDeps(t *types.Struct, supportMultiInheritance bool) (string, error) {
	deps := GetStructDeps(t)
	var dep string
	var err error

	if supportMultiInheritance {
		dep = strings.Join(*deps, ", ")
	} else {
		if len(*deps) > 1 {
			err = fmt.Errorf("This language does not support multipule inheritance.")
		}

		if len(*deps) > 0 {
			dep = (*deps)[0]
		}
	}

	return dep, err
}
