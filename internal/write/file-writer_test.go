package write

import (
	"fmt"
	"os"
	"testing"
)

func TestFileWriter(t *testing.T) {
	testString := "Lorem Ipsum"
	path := "../../examples/test.ts"
	err := FileWriter(path, testString)

	if err != nil {
		msg := fmt.Sprintf("Could not write to file path. %v", err)
		fmt.Println(msg)
		t.Error(msg)
	}

	d, err := os.ReadFile(path)

	if err != nil {
		msg := "Could not read test file."
		fmt.Println(msg)
		t.Error(msg)
	}

	contents := string(d)
	if contents != testString {
		fmt.Printf("Contents is not equal to testString:\n%s\n---\n%s\n", contents, testString)
		t.Fail()
	}

	err = os.Remove(path)
	if err != nil {
		msg := "Could not remove test file."
		fmt.Println(msg)
		t.Error(msg)
	}
}
