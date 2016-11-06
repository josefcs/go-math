package equality

import (
	"math"
	"strconv"
	"testing"
)

type nearlyEqualTestSet struct {
	left   float64
	right  float64
	expect bool
}

func TestBig(t *testing.T) {
	testSets := []nearlyEqualTestSet{
		// The precision depends also on the epsilon, so changes in the epsilon
		// might break the tests
		{10000, 10001, false},
		{100000, 100001, true},
		{100000, 10002, false},
		{100001, 100001, true},
		{1000000, 1000001, true},
		{-10000, -10001, false},
		{-100000, -100001, true},
		{-100000, -10002, false},
		{-100001, -100001, true},
		{-1000000, -1000001, true},
		// Around 1.0
		{1.0001, 1.0000, false},
		{1.000001, 1.000000, true},
		{1.0000001, 1.0000000, true},
		{-1.0001, -1.0000, false},
		{-1.000001, -1.000000, true},
		{-1.0000001, -1.0000000, true},
		// Very small numbers
		{0.000000001000000, 0.000000001000001, true},
		{0.000000001000001, 0.000000001000002, true},
		{0.0000000010001, 0.0000000010002, false},
		{0.0, 0.0, true},
		{0.0, -0.0, true},
		{0.0000001, 0.0, false},
		{-0.0000001, 0.0, false},
		// Around 0
		{0.000000001, -0.000000001, false},
		{0.00000000001, -0.00000000001, false},
		// NaN
		{math.NaN(), math.NaN(), false},
		{math.NaN(), math.Inf(1), false},
		{math.NaN(), math.Inf(-1), false},
		// Infinity
		{math.Inf(1), math.Inf(1), true},
		{math.Inf(-1), math.Inf(-1), true},
		{math.Inf(-1), math.Inf(1), false},
	}

	for index, set := range testSets {
		if RelativeEqual(set.left, set.right, Epsilon) != set.expect {
			t.Errorf("Fail: test-normal(%d) NearlyEqual( %v , %v ) should be %v", index, toFullString(set.left), toFullString(set.right), set.expect)
		}

		if RelativeEqual(set.right, set.left, Epsilon) != set.expect {
			t.Errorf("Fail: test-reversed(%d) NearlyEqual( %v , %v ) should be %v", index, toFullString(set.right), toFullString(set.left), set.expect)
		}
	}
}

// convenience function that formats a float
func toFullString(f float64) string {
	return strconv.FormatFloat(f, 'g', 6, 64)
}
