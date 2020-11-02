package utils

import "testing"

func TestContainsAll(t *testing.T) {
	if !ContainsAll("Lorem ipsum", []string{"Lorem", "ipum"}) {
		t.Error()
	}

	if ContainsAll("Nothing in here", []string{"xyz"}) {
		t.Error()
	}
}
