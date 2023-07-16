package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	string1 := "a4bc2d5e"
	value, _ := unpackString(string1)
	println(value)
}

// Unpack выполняет распаковку строки, содержащую повторяющиеся символы.
func unpackString(s string) (r string, err error) {
	if _, err := strconv.Atoi(s); err == nil {
		return r, errors.New("некорректная строка")
	}

	var prev rune
	var escaped bool
	var b strings.Builder
	for _, char := range s {
		if unicode.IsDigit(char) && !escaped {
			m := int(char - '0')
			r := strings.Repeat(string(prev), m-1)
			b.WriteString(r)
		} else {
			escaped = string(char) == "\\" && string(prev) != "\\"
			if !escaped {
				b.WriteRune(char)
			}
			prev = char
		}
	}

	return b.String(), err
}
