package read

import (
	"io/ioutil"
	"testing"

	"github.com/lil5/typex2/internal/utils"
)

func GetSut(t *testing.T, path string, fname string) (*utils.StructMap, *string) {
	data, err := ioutil.ReadFile(path + "/" + fname)

	if err != nil {
		t.FailNow()
	}

	result := string(data)

	pkg, err := LoadPackage(path)

	if err != nil {
		t.FailNow()
	}

	sut := MapPackage(pkg)

	if sut == nil {
		t.FailNow()
	}

	return sut, &result
}
