package main

import (
	"fmt"
	"strings"
)

func main() {
	//输入: "the sky is blue"
	//输出: "blue is sky the"

	fmt.Println(reverseWords(" 1"))
}

func reverseWords(s string) string {
	sb := strings.Builder{}

	for {
		a, b := findab(s)
		if a == 0 && b == 0 {
			if sb.Len() != 0 {
				return sb.String()[0 : sb.Len()-1]
			} else {
				return ""
			}
		}
		sb.WriteString(s[a:b])
		sb.WriteString(" ")

		s = s[:a]
	}
}

func findab(s string) (int, int) {
	b := len(s)
	a := b
	abc := false

	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 {
			if s[i] != ' ' {
				abc = true
			}
			if abc {
				if s[i] != ' ' {
					return 0, b
				} else {
					return a, b
				}
			}
			return 0, 0
		}

		if s[i] == ' ' {
			if abc {
				return a, b
			}
			b = i
			a = b
		} else {
			a = i
			abc = true
		}
	}

	return 0, 0
}
