package typescript

import (
	"testing"

	"github.com/lil5/typex2/internal/read"
	"github.com/lil5/typex2/internal/utils"
)

func TestGetTypeContent(t *testing.T) {
	sut, result := read.GetSut(t, "../../../examples", "index.ts")
	indent := 0

	for name, tt := range *sut {
		ttContent := getTypeContent(tt, &indent)

		hasAll := utils.ContainsAll(*result, []string{ttContent, name})

		if !hasAll {
			t.Error()
		}
	}
}
