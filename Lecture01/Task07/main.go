package main

import (
	"fmt"
	"os"
)

func main(){
	var n,k,m,res int

	f,_ := os.Open("input.txt")
	fmt.Fscanf(f, "%d %d %d", &n, &k, &m)

	if n >= k && k >= m{
		for n / k > 0{
			nk := n / k
			n = n % k
			for i := 0; i < nk; i++{
				ik := k
	
				res += k / m
	
				n += ik % m
			 }
		}
	}else{
		res = 0
	}

	f,_ = os.Create("output.txt")
	fmt.Fprintf(f, "%d", res)
}