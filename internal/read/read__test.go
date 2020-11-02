package read

import (
	"testing"

	"github.com/lil5/typex2/internal/utils"
)

func TestRead(t *testing.T) {
	sut, result := GetSut(t, "../../examples", "index.ts")

	if sut == nil {
		t.FailNow()
	}

	for name := range *sut {
		if !utils.ContainsAll(*result, []string{name}) {
			t.Errorf("Does not contain %v", name)
		}
	}

}
