package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	var y float64
	if x < -7.0 || x > 3.0 {
		fmt.Println("Function is undefined")
		return math.NaN()
	} else if x <= -6.0 {
		y = 2
	} else if x <= -2.0 {
		// line throw two point
		// y = (x + a) / b
		// y(-2) = 0  ==>  (-2 + a)/b = 0 ==> a = 2 ==> f = (x + 2) / b
		// y(-6) = -1 ==>  (-6 + 2)/b = -1 => -4/b = -1 ==> b = 4 ==> y = (x+2)/4
		y = (x + 2) / 4
	} else if x <= -0.0 {
		// equation of circle (x - x0)^2 + (y -y0)^2 = R^2
		// where x0, y0 - center of circle
		// in our case x0 = -2, y0 = 2, R = 2
		// (x+2)^2 + (y-2)^2 = 4 ==>
		// (y-2)^2 = 4 - (x+2)^2
		// y = 2 + sqrt(4 - (x+2)^2) OR y = 2 - sqrt(4 - (x+2)^2)
		// in our case we have y = 2 - sqrt(4 - (x+2)^2)
		// because we have lower part of circle
		y =  2 - math.Sqrt(4 - math.Pow(x+2, 2))
	} else if x <= 2.0 {
		// equation of circle (x - x0)^2 + (y -y0)^2 = R^2
		// where x0, y0 - center of circle
		// in our case x0 = 0, y0 = 0, R = 2
		// x^2 + y^2 = 4 ==>
		// y^2 = 4 - x^2
		// y = sqrt(4 - x^2) OR y = - sqrt(4 - x^2)
		// in our case we have y = sqrt(4 - x^2)
		// because we have upper part of circle
		y = math.Sqrt(4 - x*x)
	} else if x <= 3.0 {
		// line throw two point
		// y = (x + a) / b
		// y(2) = 0  ==>  (2 + a)/b = 0 ==> a = -2 ==> f = (x - 2) / b
		// y(3) = -1 ==>  (3 - 2)/b = -1 => 1/b = -1 ==> b = -1 ==> y = 2 - x
		y = 2 - x
	}
	return  y
}

func main() {
	fmt.Println("Enter x: ")
	var x float64
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Error in reading x.")
		return
	}
	y := f(x)
	fmt.Printf("y(%v) =  %v\n", x, y)
}
