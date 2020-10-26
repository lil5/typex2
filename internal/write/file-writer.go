package write

import (
	"os"
)

func FileWriter(
	path string, fname string, s *string) error {
	f, err := os.Create(path + "/" + fname)
	if err != nil {
		return err
	}
	_, err = f.WriteString(*s)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
