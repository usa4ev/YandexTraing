package main

import (
	"fmt"
	"os"
)

func main() {
	var a1, b1, a2, b2, r1, r2 int
	f, _ := os.Open("input.txt")
	fmt.Fscanf(f, "%d %d %d %d", &a1, &b1, &a2, &b2)

	if a1 == a2 && a1 == b1 && a1 == b2 {
		// это все квадраты
		r1, r2 = a1, a1
	}


	s1 := a1 * b1
	s2 := a2 * b2

	if s1 > s2{
		switch{
		case min(a1,b1) >= max(a2,b2):
			r1 = max(a1, b1) + min(a2, b2)
			r2 = max(min(a1, b1), max(a2, b2)) 
		case max(a1,b1) > max(a2,b2):
			r1 = min(a1, b1) + min(a2, b2)
			r2 = max(max(a1, b1), max(a2, b2)) 
		default:
			r1 = min(a1, b1) + min(a2, b2)
			r2 = max(max(a1, b1), max(a2, b2))
		}
	}else{
		switch{
		case min(a2,b2) >= max(a1,b1):
			r1 = max(a2, b2) + min(a1, b1)
			r2 = max(min(a2, b2), max(a1, b1)) 
		case max(a2,b2) > max(a1,b1):
			r1 = min(a2, b2) + min(a1, b1)
			r2 = max(max(a2, b2), max(a1, b1)) 
		default:
			r1 = min(a2, b2) + min(a1, b1)
			r2 = max(max(a2, b2), max(a1, b1))
		}
	}

	f, _ = os.Create("output.txt")
	fmt.Fprintf(f, "%d %d", r1, r2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
