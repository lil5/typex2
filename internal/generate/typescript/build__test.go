package typescript

import (
	"strings"
	"testing"
)

func TestBuidInterface(t *testing.T) {
	name := "Lorem"

	bi1, bi2 := buildInterface(name, &[]string{})

	bi1MustContain := []string{name, "interface", "{"}
	if !containsAll(bi1, bi1MustContain) {
		t.Errorf("bi1 Does not contain %v", bi1MustContain)
	}
	if !strings.Contains(bi2, "}") {
		t.Errorf("bi2 Does not contain %s", "}")
	}
}

func containsAll(s string, substrs []string) bool {
	ca := false
	for i := 0; i < len(substrs) && !ca; i++ {
		ca = strings.Contains(s, substrs[i])
	}

	return ca
}

func TestBuildTypeAlias(t *testing.T) {
	name := "Lorem"

	bta1, _ := buildTypeAlias(name)

	bta1MustContain := []string{name, "type", " = "}
	if !containsAll(bta1, bta1MustContain) {
		t.Errorf("bta1 Does not contain %v", bta1MustContain)
	}
}
