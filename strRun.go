package main

import (
	"fmt"

	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {

		fmt.Println("x%", s[i])

	}
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d", runeValue, idx)

	}

	for i, w := 0, 0; i < len(s); i += w {

	}
}
