package main

import "testing"

func TestSquare(t *testing.T) {
	for i := 1; i < 100000000; i++ {
		result := square(i)
		if result/i != i {
			t.Fatalf("expected: %d, result: %d\n", i*i, result)
		}
	}
}

func TestDayOfWeek(t *testing.T) {
	type test struct {
		Input    int
		Expected string
	}
	var testCase = []test{
		{Input: 1, Expected: "shanbe"},
		{Input: 2, Expected: "yeshanbe"},
		{Input: 3, Expected: "doshanbe"},
		{Input: 4, Expected: "seshanbe"},
		{Input: 5, Expected: "charshanbe"},
		{Input: 6, Expected: "panjshanbe"},
		{Input: 7, Expected: "jomeh"},
		{Input: 8, Expected: ""},
		{Input: 0, Expected: ""},
		{Input: -1, Expected: ""},
	}
	for _, c := range testCase {
		result := dayOfWeek(c.Input)
		if result != c.Expected {
			t.Errorf("expected: %s, result: %s\n", c.Expected, result)
		}
	}
}
