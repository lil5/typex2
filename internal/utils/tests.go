package utils

import "strings"

func ContainsAll(s string, substrs []string) bool {
	ca := false
	for i := 0; i < len(substrs) && !ca; i++ {
		ca = strings.Contains(s, substrs[i])
	}

	return ca
}
