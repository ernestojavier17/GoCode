package string

import _ "fmt"

func Reverse(s string) string {

	a := []byte(s) // Use []rune(s) for appropriate use distinct Unicode code points instead of bytes

	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		a[i], a[j] = a[j], a[i]
	}

	return string(a)
}
