/*
	Based on posts by John D Cook
	 - http://www.johndcook.com/blog/standard_deviation/
*/
package statistics

import "math"

// RunningStat Represents a running statistics object, where the values are calculated with
// every added value. This calculation is numerically relatively stable.
// implements the Statistics interface
type RunningStat struct {
	n    uint
	newM float64
	oldM float64
	newS float64
	oldS float64
}

// NewRunningStat
func NewRunningStat() *RunningStat {
	return &RunningStat{}
}

func (r *RunningStat) AddValue(x float64) {
	r.n++
	// See Knuth TAOCP vol 2, 3rd edition, page 232
	if r.n == 1 {
		r.oldM = x
		r.newM = x
		r.oldS = 0.0
	} else {
		r.newM = r.oldM + (x-r.oldM)/float64(r.n)
		r.newS = r.oldS + (x-r.oldM)*(x-r.newM)

		// set up for next iteration
		r.oldM = r.newM
		r.oldS = r.newS
	}
}

func (r *RunningStat) N() uint {
	return r.n
}

func (r *RunningStat) Mean() float64 {
	return r.newM
}

func (r *RunningStat) PopulationVariance() float64 {
	if r.n > 1 {
		return r.newS / float64(r.n)
	}

	return 0.0
}

func (r *RunningStat) SampleVariance() float64 {
	if r.n > 1 {
		return r.newS / (float64(r.n) - 1)
	}

	return 0.0
}

func (r *RunningStat) PopulationStandardDeviation() float64 {
	return math.Sqrt(r.PopulationVariance())
}

func (r *RunningStat) SampleStandardDeviation() float64 {
	return math.Sqrt(r.SampleVariance())
}
