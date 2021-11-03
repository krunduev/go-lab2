package main

import (
	"fmt"
)

func get(name string) float64 {
	fmt.Printf("Enter %v: \n", name)
	var res float64
	_, err := fmt.Scanln(&res)
	if err != nil {
		panic(fmt.Sprintf("Error in reading %v.\n", name))
	}
	return res
}

func check(x float64, y float64, r float64) bool {
	if x*x + y*y > r*r {
		return false
	}
	if x >=0 && y >= 0 && x <= y {
		return true
	}
	if x <= 0 && y <= 0 && x >= y {
		return true
	}
	return false
}

func main() {
	r := get("R")
	x := get("X")
	y := get("Y")

	result := check(x, y, r)
	if result {
		fmt.Printf("point (%v, %v) is in area\n", x, y)
	} else {
		fmt.Printf("point (%v, %v) is not in area\n", x, y)
	}
}
