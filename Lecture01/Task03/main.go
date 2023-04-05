package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	var in string

	book := [3]string{}

	f, _ := os.Open("input.txt")
	fmt.Fscanf(f, "%s\n%s\n%s\n%s", &in, &book[0], &book[1], &book[2])

	in = normalize(in)

	f, _ = os.Create("output.txt")

	for i, v := range book {
		if i > 0 {
			fmt.Fprintf(f, "\n")
		}

		if in == normalize(v) {
			fmt.Fprintf(f, "YES")
		} else {
			fmt.Fprintf(f, "NO")
		}
	}
}

func normalize(s string) string {
	res := make([]rune, 0)
	for _, r := range s {
		if unicode.IsDigit(r) {
			res = append(res, r)
		}
	}

	switch len(res) {
	case 11: // country and city code
		res = res[1:]
	case 7: // no code provided, add default code
		res = append([]rune{'4', '9', '5'}, res...)
	}

	return string(res)
}
