package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")

	i := 0
	k := 0
	flag := true
	ans := "YES"

	for {
		if _, err := fmt.Fscan(f, &i); err != nil {
			break
		}

		if flag {
			k = i
			flag = !flag
			continue
		}

		if k >= i {
			ans = "NO"
			break
		}

		k = i

	}

	f, _ = os.Create("output.txt")
	fmt.Fprint(f, ans)
}
