package strings

import "fmt"

func WhereIn(s []string) string {
	newString := ""
	for i, k := range s {
		newString += fmt.Sprintf("'%s'", k)
		if i < len(s)-1 {
			newString += ","
		}
	}
	return newString
}
