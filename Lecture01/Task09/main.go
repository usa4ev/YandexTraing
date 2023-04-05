package main

import (
	"fmt"
	"os"
)

func main(){
	var a,b,c,d,e int
	f,_ := os.Open("input.txt")
	fmt.Fscanf(f,"%d\n%d\n%d\n%d\n%d", &a, &b, &c, &d, &e)

	maxB := max(d, e)
	minB := min(d, e)

	ans := ""

	if (maxB >= max(a, b) && minB >= min(a,b)) ||
	(maxB >= max(c, b) && minB >= min(c,b)) ||
	(maxB >= max(a, c) && minB >= min(a,c)){
		ans = "YES"
	}else{
		ans = "NO"
	}
	
	f,_ = os.Create("output.txt")
	fmt.Fprint(f, ans) 
}

func max(a,b int)int{
	if a > b {return a}

	return b
}

func min(a,b int)int{
	if a < b {return a}

	return b
}