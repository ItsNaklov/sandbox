package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
		var repeated strings.Builder 
		for i := 0; i < repeatCount; i++ {
			// += called "the Add AND assigment operator"
			repeated.WriteString(character)
		}
		return repeated.String()
}
