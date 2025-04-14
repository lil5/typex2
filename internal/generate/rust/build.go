package rust

func buildClass(ac *ArgCounter, name string, dep string) (string, string) {
	s1 := "struct " + name + ac.typeParams() + " "

	// if dep != "" {
	// 	s1 += "extends " + dep
	// 	s1 += " "
	// }

	s1 += "{\n"
	s2 := "}\n\n"

	return s1, s2
}

func buildTypeAlias(ac *ArgCounter, name string) (string, string) {
	s1 := "type " + name + ac.typeParams() + " = "

	s2 := ";\n\n"
	return s1, s2
}
