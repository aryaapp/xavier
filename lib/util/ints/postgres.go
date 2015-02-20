package ints

import "fmt"

func WhereIn(s []int) string {
	newString := ""
	for i, k := range s {
		newString += fmt.Sprintf("'%i'", k)
		if i < len(s)-1 {
			newString += ","
		}
	}
	return newString
}
