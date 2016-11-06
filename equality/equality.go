package equality

import (
	"math"
)

// Epsilon is an arbitrary number, change it to you needs
const Epsilon = 0.00001

// RelativeEqual Compares two float64 values if they are within a relative error.
// The smaller the values, the smaller the accepted value is.
//
// epsilon is the fraction of the allowed relative error
func RelativeEqual(a float64, b float64, epsilon float64) bool {
	diff := math.Abs(a - b)

	if math.IsNaN(a) || math.IsNaN(b) {
		return false // if one of them is NaN, they can't be equal
	} else if a == b {
		return true // handles true equality and infinity
	} else if diff < math.SmallestNonzeroFloat64 {
		// diff can only be positive (also in case a-b is -Inf), so if it is
		// smaller than the smallest nonzero float, then it must be 0.
		// This means that both are equal
		return true
	} else {
		return diff/math.Min(math.Abs(a)+math.Abs(b), math.MaxFloat64) < epsilon
	}
}
