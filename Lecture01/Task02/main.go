package main

import (
	"fmt"
	"os"
)

func main(){
	var a, b, c int
	
	ok := true

	f,_ := os.Open("input.txt")

	fmt.Fscanf(f, "%d\n%d\n%d", &a, &b, &c)

	if a + b <= c || b + c <= a || c + a <= b{
		ok = false
	}

	f,_ = os.Create("output.txt")

	if ok{
		fmt.Fprint(f, "YES")
	}else{
		fmt.Fprint(f, "NO")
	}
}