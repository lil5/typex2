package typescript

import "strings"

// run on type struct
func buildInterface(name string, deps *[]string) (string, string) {
	s := "interface " + name + " "

	if deps != nil && len(*deps) > 0 {
		s += "extends "
		s += strings.Join(*deps, ", ")
		s += " "
	}

	s += "{\n"
	s2 := "}\n\n"

	return s, s2
}

// TODO: add pointer check and run basic type or interface

func buildTypeAlias(name string) (string, string) {
	s1 := "type " + name + " = "

	s2 := "\n\n"
	return s1, s2
}
