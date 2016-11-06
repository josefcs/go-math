package statistics

// interface for running and windowed statistics
type Statistics interface {
	// AddValue adds a number to the running statistics.
	AddValue(x float64)
	// N returns the number of added samples
	N() uint
	// PopulationVariance returns the variance of added values (sum/n).
	PopulationVariance() float64
	// SampleVariance returns the sample variance (sum / (n-1)).
	SampleVariance() float64
	// PopulationStandardDeviation returns the standard deviation calculated with n.
	PopulationStandardDeviation() float64
	// SampleStandardDeviation returns the standard deviation calculated with (n-1).
	SampleStandardDeviation() float64
	// Mean is the average over all added values
	Mean() float64
}
