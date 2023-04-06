package main

import (
	"fmt"
	"os"
)

func main() {
	fin, _ := os.Open("input.txt")
	defer fin.Close()

	ch := make(chan int, 1)

	go func() {
		var i int
		for {
			if _,err := fmt.Fscan(fin, &i); err != nil {
				close(ch)
				break
			}

			ch <- i
		}
	}()

	ans := alg(ch)

	fout, _ := os.Create("output.txt")
	defer fout.Close()
	fmt.Fprint(fout, ans)
}

func alg(ch chan int) int {
	var l,i,r,c,ans int
	
	for r = range ch {
		c++

		// read first two elements
		if c < 3{
			l = i
			i = r
			continue
		}

		if i > l && i > r{
			ans ++
		}

		l = i
		i = r
	}
	
	return ans
}
