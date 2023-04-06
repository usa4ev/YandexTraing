package main

import (
	"fmt"
	"os"
)

func main() {
	fin, _ := os.Open("input.txt")
	defer fin.Close()

	ch := make(chan int, 1)

	var n int
	fmt.Fscan(fin, &n)

	go func() {
		var i int
		for {
			if _, err := fmt.Fscan(fin, &i); err != nil {
				close(ch)
				break
			}

			ch <- i
		}
	}()

	ans := alg(n, ch)

	fout, _ := os.Create("output.txt")
	defer fout.Close()
	fmt.Fprint(fout, ans)
}

func alg(n int, ch chan int) int {
	a := make([]int, n)

	// 1) read ch into int slice
	// 2) find 3rd max
	var max, i int
	for v := range ch {
		a[i] = v

		if v > max {
			max = v
		}

		i++
	}

	// 4) build map(a)
	m := make(map[int]int)
	maxmet := false
	for i, v := range a {
		if !maxmet {
			if v >= max {
				maxmet = true
			}
			continue
		}

		if v%10 == 5 {
			m[i]++
		}

		if i+1 < n && v > a[i+1] {
			m[i]++
		}
	}

	res := -1
	for k, v := range m {
		if v == 2 && a[k] > res {
			res = a[k]
		}
	}

	if res == -1 {
		// no matching result found
		return 0
	}

	ans := 1
	for _, v := range a {
		if v > res {
			ans++
		}
	}

	return ans
}
