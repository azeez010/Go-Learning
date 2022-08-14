package main

import (
	"testing"
)

type Testcase struct {
	Tests []int
	Value int
}

func TestCalc(t *testing.T) {
	tests := []Testcase{Testcase{[]int{2, 3, 4, 11}, 20}, Testcase{[]int{2, 3, 4, 11, 5}, 15}}

	for _, test := range tests {
		if Calc(test.Tests...) != test.Value {
			t.Error("Test failed!")
		}
	}

}
