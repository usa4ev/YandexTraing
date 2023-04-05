package main

import (
	"fmt"
	"os"
)

const (
	// modes
	freeze = "freeze"
	heat   = "heat"
	auto   = "auto"
	fan    = "fan"
)

func main() {
	f, _ := os.Open("input.txt")
	var troom, tcond int
	var mode string

	fmt.Fscanf(f, "%d %d\n%s", &troom, &tcond, &mode)

	res := executeTemp(troom, mode, tcond)

	f, _ = os.Create("output.txt")

	fmt.Fprint(f, res)
}

func executeTemp(troom int, mode string, tcond int) int {
	res := troom

	switch mode {
	case freeze:
		if tcond < troom {
			res = tcond
		}
	case heat:
		if tcond > troom {
			res = tcond
		}
	case auto:
		res = tcond
	}

	return res
}
