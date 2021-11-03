package main

import (
	"math"
	"testing"
)

type TestCase struct {
	x	float64
	y	float64
}

func Test(t *testing.T) {
	cases := []TestCase{
		TestCase{x: -10,y: math.NaN()},
		TestCase{x: -7,y: 2},
		TestCase{x: -6.5,y: 2},
		TestCase{x: -4,y: -0.5},
		TestCase{x: -2,y: 0},
		TestCase{x: -1,y: 2 - math.Sqrt(3)},
		TestCase{x: 0,y: 2},
		TestCase{x: 1,y: math.Sqrt(3)},
		TestCase{x: 2,y: 0},
		TestCase{x: 2.5,y: -0.5},
		TestCase{x: 3,y: -1},
		TestCase{x: 4,y: math.NaN()},
	}

	for caseNum, item := range cases {
		y := f(item.x)
		if math.IsNaN(y) && math.IsNaN(item.y) {
			continue // ok
		}

		if y != item.y {
			t.Errorf("Case %v, f(%v) = %v, expected %v",
				caseNum, item.x, y, item.y)
		}
	}
}
