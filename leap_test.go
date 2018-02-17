package leap

import (
	"testing"
)

// Check a few "golden" values to see that implementations across languages are equivalent
func TestLinearCongruentialGeneratorCompatibility(t *testing.T) {
	golden100 := []int32{0, 55, 62, 8, 45, 59, 86, 97, 82, 59, 73, 37, 17, 56, 86, 21, 90, 37, 38, 83}
	for i, v := range golden100 {
		if result := Hash(uint64(i), 100); result != v {
			t.Errorf("Hash(%v, 100)=%v, expected %v\n", i, result, v)
		}
	}

	var testCases = []struct {
		key      uint64
		buckets  int
		expected int32
	}{
		{10863919174838991, 11, 6},
		{2016238256797177309, 11, 3},
		{1673758223894951030, 11, 5},
		{2, 100001, 80343},
		{2201, 100001, 22152},
		{2202, 100001, 15018},
	}

	for _, testCase := range testCases {
		if g := Hash(testCase.key, testCase.buckets); g != testCase.expected {
			t.Errorf("Hash(%v, %v)=%v, expected %v\n", testCase.key, testCase.buckets, g, testCase.expected)

		}
	}
}
