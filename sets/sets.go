package sets

import (
	"sort"
)

// Intersect returns an array of strings that are contained in both slices
func Intersect(set1 []string, set2 []string) []string {
	helperSet := make(map[string]bool)
	var andSet []string

	for _, entry := range set1 {
		helperSet[entry] = false
	}

	for _, candidate := range set2 {
		if _, ok := helperSet[candidate]; ok {
			andSet = append(andSet, candidate)
		}
	}

	return andSet
}

// Complement returns an array, where set2 is "subtracted" from set 1
func Complement(set1 []string, set2 []string) []string {
	helperSet := make(map[string]bool)

	for _, entry := range set2 {
		helperSet[entry] = true
	}

	var sub []string

	for _, entry := range set1 {
		if _, contained := helperSet[entry]; !contained {
			sub = append(sub, entry)
		}
	}

	return sub
}

// Missing checks if all items from required are in got.
// Returns the missing items from got if there are any
func Missing(required []string, got []string) []string {

	having := Intersect(required, got)

	// Subtract the items we fulfill from all required ones.
	// The result is empty if all requirements are met.
	return Complement(required, having)
}

// StrictlyEqual checks if two arrays are equal (also their order)
func StrictlyEqual(a, b []string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// SetEqual checks if one array is the same as the other - apart from the order.
func SetEqual(a, b []string) bool {
	setA := make([]string, len(a))
	copy(setA, a)
	sort.Sort(sort.StringSlice(setA))

	setB := make([]string, len(b))
	copy(setB, b)
	sort.Sort(sort.StringSlice(setB))

	return StrictlyEqual(setA, setB)
}
