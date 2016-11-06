package statistics

import (
	"testing"

	"github.com/josefcs/go-math/equality"
)

type testsetdata struct {
	values         []float64
	mean           float64
	popStdDev      float64
	popVariance    float64
	sampleStdDev   float64
	sampleVariance float64
}

var tests = []testsetdata{
	{values: []float64{1, 2},
		mean:           1.5,
		popStdDev:      0.5,
		popVariance:    0.25,
		sampleStdDev:   0.7071067811865476,
		sampleVariance: 0.5},
	{values: []float64{936, 1628, 1330, 827, 1609, 1190, 877, 901, 1636, 815, 574, 882, 1158, 1740, 630, 1221, 560, 1039, 1460, 696},
		mean:           1085.45,
		popStdDev:      368.4256336087379,
		popVariance:    135737.44749999998,
		sampleStdDev:   377.9967244358217,
		sampleVariance: 142881.52368421052},
}

func TestRunning(t *testing.T) {
	for index, set := range tests {
		var stat Statistics = &RunningStat{}

		for _, val := range set.values {
			stat.AddValue(val)
		}

		if int(stat.N()) != len(set.values) {
			t.Errorf("Fail: number of points, set %d: got %d, want approx %d", index, int(stat.N()), len(set.values))
			t.FailNow()

		}

		if !equality.RelativeEqual(stat.Mean(), set.mean, equality.Epsilon) {
			t.Errorf("Fail: arithmetic mean, set %d: got %g, want approx %g", index, stat.Mean(), set.mean)
		}

		if !equality.RelativeEqual(stat.PopulationStandardDeviation(), set.popStdDev, equality.Epsilon) {
			t.Errorf("Fail: pop std dev, set %d: got %g, want approx %g", index, stat.PopulationStandardDeviation(), set.popStdDev)
		}

		// TODO: also test the other functions
	}
}
