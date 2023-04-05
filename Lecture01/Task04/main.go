package main

import (
	"fmt"
	"math"
	"os"
)

func main(){
	var a,b,c int
	var res any

	f,_ := os.Open("input.txt")
	fmt.Fscanf(f, "%d\n%d\n%d", &a, &b, &c)

	if a == 0{
		if b < 0 || b != c * c{
			res = "NO SOLUTION"
		}else{
			res = "MANY SOLUTIONS"
		}
		
	}else{
		x := (c*c - b) / a
		if int(math.Sqrt(float64(a * x + b))) != c{
			res = "NO SOLUTION"
		}else{
			res = x
		}
	}

	f, _ = os.Create("output.txt")
	fmt.Fprint(f, res)
}