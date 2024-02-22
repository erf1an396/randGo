package mathematic

import (
	"math"
	"testing"
)

func TestPoewrH(t *testing.T) {
	type test struct {
		Base           float64
		Power          float64
		ExpectedResult float64
	}

	testCases := []test{
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 2, Power: 5, ExpectedResult: 32},
	}
	for _, c := range testCases {
		res := PoewrH(c.Base, c.Power)
		if res != c.ExpectedResult {
			t.Errorf("expected: %f, result: %f\n", c.ExpectedResult, res)
		}
	}
}

func TestMathPow(t *testing.T) {
	type test struct {
		Base           float64
		Power          float64
		ExpectedResult float64
	}

	testCases := []test{
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 2, Power: 5, ExpectedResult: 32},
	}
	for _, c := range testCases {
		res := math.Pow(c.Base, c.Power)
		if res != c.ExpectedResult {
			t.Errorf("expected: %f, result: %f\n", c.ExpectedResult, res)
		}
	}
}

var result float64

func BenchmarkPoewrH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := PoewrH(2, 4)
		result = res
	}
}

func BenchmarkPowerMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := PowerMath(2, 4)
		result = res
	}
}

func BenchmarkPowerSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := PowerSlice(2, 4)
		result = res
	}
}

func BenchmarkPowerSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := PowerSlice2(2, 4)
		result = res
	}
}
