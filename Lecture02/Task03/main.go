package main

import (
	"fmt"
	"os"
)

func main(){
	var(
		n,x,ans int
		a []int
	)

	fin,_ := os.Open("input.txt")
	defer fin.Close()

	fmt.Fscanln(fin, &n)

	a = make([]int,n)

	for i := 0; i < n; i++{
		var ai int 
		fmt.Fscan(fin, &ai)
		a[i] = ai
	}

	fmt.Fscan(fin, &x)

	ans = alg(a, x)

	fout,_ := os.Create("output.txt")
	defer fout.Close()

	fmt.Fprint(fout, ans)
}

func alg(a []int, x int) int{
	var min,lmin,ans int 
	flag := true

	for _,v := range a{
		lmin = x - v

		if lmin < 0{
			lmin = -lmin
		}

		if flag{
			flag = !flag
			min = lmin
			ans = v
			continue
		}
		
		if lmin < min{
			min = lmin
			ans = v
		}
	}

	return ans
}