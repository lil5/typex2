package dart

func buildClass(name string, dep string) (string, string) {
	s1 := "class " + name + " "

	if dep != "" {
		s1 += "extends " + dep
		s1 += " "
	}

	s1 += "{\n"
	s2 := "}\n\n"

	return s1, s2
}

func buildTypeAlias(name string) (string, string) {
	s1 := "typedef " + name + " = "

	s2 := "\n\n"
	return s1, s2
}
