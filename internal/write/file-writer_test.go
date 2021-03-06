package write

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileWriter(t *testing.T) {
	testString := "Lorem Ipsum"
	path := "../../examples"
	fname := "test.ts"
	err := FileWriter(path, fname, &testString)

	if err != nil {
		msg := fmt.Sprintf("Could not write to file path. %v", err)
		fmt.Println(msg)
		t.Errorf(msg)
	}

	fullpath := path + "/" + fname
	d, err := ioutil.ReadFile(fullpath)

	if err != nil {
		msg := "Could not read test file."
		fmt.Println(msg)
		t.Errorf(msg)
	}

	contents := string(d)
	if contents != testString {
		fmt.Printf("Contents is not equal to testString:\n%s\n---\n%s\n", contents, testString)
		t.Fail()
	}

	err = os.Remove(fullpath)
	if err != nil {
		msg := "Could not remove test file."
		fmt.Println(msg)
		t.Errorf(msg)
	}
}
