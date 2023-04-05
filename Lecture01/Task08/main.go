package main

import (
	"fmt"
	"os"
)

func main(){
	var a,b,n,m, min, max int

	f,_ := os.Open("input.txt")
	fmt.Fscanf(f,"%d\n%d\n%d\n%d", &a, &b, &n, &m)

	nmin, nmax := ((a+1) * n) - a, ((a+1) * n) + a
	mmin, mmax := ((b+1) * m) - b, ((b+1) * m) + b

	f,_ = os.Create("output.txt")

	if nmax < mmin || mmax < nmin{
		fmt.Fprintf(f, "%d", -1)
		return
	} 

	max = minInt(nmax, mmax)
	min = maxInt(nmin, mmin)

	fmt.Fprintf(f, "%d %d", min, max)
}

func maxInt(a,b int) int{
	if a > b{ return a}

	return b
}

func minInt(a,b int) int{
	if a < b{ return a}

	return b
}