package sets

import (
	"testing"
)

func TestIntersect(t *testing.T) {
	data := []struct {
		set1 []string
		set2 []string
		want []string
	}{
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "d", "e"}, []string{"a", "b"}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"c"}, []string{"c"}},
		{[]string{"b", "c", "a"}, []string{"c", "d", "b", "x", "z", "a", "l"}, []string{"a", "b", "c"}},
	}

	for _, val := range data {
		if result := Intersect(val.set1, val.set2); !SetEqual(result, val.want) {
			t.Errorf("Fail: Intersect(%v, %v) expected %v but got %v", val.set1, val.set2, val.want, result)
		}
	}
}

func TestComplement(t *testing.T) {
	data := []struct {
		set1 []string
		set2 []string
		want []string
	}{
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c"}, []string{"d"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d", "e"}, []string{}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"c"}, []string{"a", "b", "d", "e"}},
		{[]string{"b", "c", "a", "q"}, []string{"c", "d", "b", "x", "z", "a", "l"}, []string{"q"}},
	}

	for _, val := range data {
		if result := Complement(val.set1, val.set2); !SetEqual(result, val.want) {
			t.Errorf("Fail: Complement(%v, %v) expected %v but got %v", val.set1, val.set2, val.want, result)
		}
	}
}

func TestMissing(t *testing.T) {
	data := []struct {
		required []string
		got      []string
		want     []string
	}{
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c"}, []string{"d"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, []string{}},
		{[]string{"a", "b", "c"}, []string{"c"}, []string{"a", "b"}},
		{[]string{}, []string{"a", "b", "c"}, []string{}},
		{[]string{"a", "b"}, []string{}, []string{"a", "b"}},
	}

	for _, val := range data {
		if result := Missing(val.required, val.got); !SetEqual(result, val.want) {
			t.Errorf("Fail: Missing(%v, %v) expected %v but got %v", val.required, val.got, val.want, result)
		}
	}
}

func TestContains(t *testing.T) {
	data := []struct {
		element string
		set     []string
		want    bool
	}{
		{"a", []string{"c", "f", "ab"}, false},
		{"a", []string{"c", "f", "a"}, true},
		{"a", []string{"c", "", "ab"}, false},
		{"ab", []string{"c", "f", ""}, false},
		{"a", []string{"a", "f", "ab"}, true},
		{"", []string{"a", "f", "ab"}, false},
		{"", []string{"a", "f", "ab", ""}, true},
	}

	for _, val := range data {
		if result := Contains(val.element, val.set); result != val.want {
			t.Errorf("Fail: Contains('%s', %+v) should be %+v", val.element, val.set, val.want)
		}
	}
}

func BenchmarkIntersectSmall(b *testing.B) {
	set1 := []string{"a", "b", "c", "d"}
	set2 := []string{"a", "b", "c"}

	for i := 0; i < b.N; i++ {
		Intersect(set1, set2)
	}
}

func BenchmarkIntersectMedium(b *testing.B) {
	set1 := []string{"one", "two", "three", "four"}
	set2 := []string{"four", "two", "one", "three", "five"}

	for i := 0; i < b.N; i++ {
		Intersect(set1, set2)
	}
}

func BenchmarkIntersectLarge(b *testing.B) {
	set1 := []string{"prefix.one", "prefix.two", "prefix.three", "prefix.four", "prefix.five", "prefix.six"}
	set2 := []string{"prefix.three", "prefix.one", "prefix.none", "prefix.none.second", "prefix.six", "prefix.two"}

	for i := 0; i < b.N; i++ {
		Intersect(set1, set2)
	}
}

func BenchmarkComplementSmall(b *testing.B) {
	set1 := []string{"a", "b", "c", "d"}
	set2 := []string{"a", "b", "c"}

	for i := 0; i < b.N; i++ {
		Complement(set1, set2)
	}
}

func BenchmarkComplementLarge(b *testing.B) {
	set1 := []string{"prefix.one", "prefix.two", "prefix.three", "prefix.four", "prefix.five", "prefix.six"}
	set2 := []string{"prefix.three", "prefix.one", "prefix.none", "prefix.none.second", "prefix.six", "prefix.two"}

	for i := 0; i < b.N; i++ {
		Complement(set1, set2)
	}
}

func BenchmarkMissingSmall(b *testing.B) {
	set1 := []string{"a", "b", "c", "d"}
	set2 := []string{"a", "b", "c"}

	for i := 0; i < b.N; i++ {
		Missing(set1, set2)
	}
}

func BenchmarkMissingMedium(b *testing.B) {
	set1 := []string{"one", "two", "three", "four"}
	set2 := []string{"four", "two", "one", "three", "five"}

	for i := 0; i < b.N; i++ {
		Missing(set1, set2)
	}
}

func BenchmarkMissingLarge(b *testing.B) {
	set1 := []string{"prefix.one", "prefix.two", "prefix.three", "prefix.four", "prefix.five", "prefix.six", "prefix.seven", "prefix.eight", "prefix.nine", "prefix.ten", "prefix.eleven", "prefix.twelve", "prefix.thirteen"}
	set2 := []string{"prefix.three", "prefix.one", "prefix.none", "prefix.none.second", "prefix.six", "prefix.two"}

	for i := 0; i < b.N; i++ {
		Missing(set1, set2)
	}
}
