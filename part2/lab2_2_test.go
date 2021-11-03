package main

import (
	"encoding/json"
	"testing"
)

type TestCase struct {
	X float64
	Y float64
	R float64
	Res bool
}

func Test(t *testing.T) {
	cases := []TestCase{
		TestCase{X: 0, Y: 0, R: 10, Res: true},
		TestCase{X: 10, Y: 10, R: 10, Res: false},
		TestCase{X: 2, Y: 3, R: 10, Res: true},
		TestCase{X: -2, Y: -3, R: 10, Res: true},
		TestCase{X: 2, Y: 1, R: 10, Res: false},
		TestCase{X: -2, Y: -1, R: 10, Res: false},
		TestCase{X: 5, Y: -5, R: 10, Res: false},
		TestCase{X: -5, Y: 5, R: 10, Res: false},
	}

	for caseNum, item := range cases {
		res := check(item.X, item.Y, item.R)
		if res != item.Res {
			out, _ := json.Marshal(item)
			t.Errorf("Case %v, expected %v, get %v for %v", caseNum, item.Res, res, string(out))
		}
	}
}
