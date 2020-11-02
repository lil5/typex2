package typescript

import (
	"strings"
	"testing"

	"github.com/lil5/typex2/internal/utils"
)

func TestBuidInterface(t *testing.T) {
	name := "Lorem"

	bi1, bi2 := buildInterface(name, &[]string{})

	bi1MustContain := []string{name, "interface", "{"}
	if !utils.ContainsAll(bi1, bi1MustContain) {
		t.Errorf("bi1 Does not contain %v", bi1MustContain)
	}
	if !strings.Contains(bi2, "}") {
		t.Errorf("bi2 Does not contain %s", "}")
	}
}

func TestBuidInterfaceWithExtends(t *testing.T) {
	name := "Lorem"
	extends := []string{"Ipsum", "Dolor"}

	sut1, _ := buildInterface(name, &extends)

	mustContain := append(extends, "extends")

	if !utils.ContainsAll(sut1, mustContain) {
		t.Errorf("Does not contain extended values")
	}
}

func TestBuildTypeAlias(t *testing.T) {
	name := "Lorem"

	bta1, _ := buildTypeAlias(name)

	bta1MustContain := []string{name, "type", " = "}
	if !utils.ContainsAll(bta1, bta1MustContain) {
		t.Errorf("bta1 Does not contain %v", bta1MustContain)
	}
}
