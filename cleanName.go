package main

import "strings"

// GetTheNames ... asd
// aaa [aaa]
func getTheNames(name string) string {
	i := strings.Index(name, "[")
	if i >= 0 {
		j := strings.Index(name[i:], "]")
		if j >= 0 {
			return name[i+1 : j+i]
		}
	}
	return ""
}
