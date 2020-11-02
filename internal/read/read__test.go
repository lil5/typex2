package read

import (
	"io/ioutil"
	"testing"

	"github.com/lil5/typex2/internal/utils"
)

func TestRead(t *testing.T) {
	data, err := ioutil.ReadFile("../../examples/index.ts")

	if err != nil {
		t.FailNow()
	}

	pkg, err := LoadPackage("../../examples")

	if err != nil {
		t.FailNow()
	}

	sut := MapPackage(pkg)

	if sut == nil {
		t.FailNow()
	}

	for name := range *sut {
		if !utils.ContainsAll(string(data), []string{name}) {
			t.Errorf("Does not contain %v", name)
		}
	}

}
