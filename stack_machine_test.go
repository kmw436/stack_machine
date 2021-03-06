package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	sm := NewStackMachine()
	testCases := map[string]int{
		"13+62*7+*": 76,
		"11++":      -1,
	}
	for tc, expected := range testCases {
		sm.numStack = []int{}
		actual := sm.calculate(tc)
		if actual != expected {
			t.Errorf("%s: expected %d, got %d\n", tc, expected, actual)
		}
	}
}
