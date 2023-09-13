package stats

import "fmt"

// Percentiles compute the percentiles  p_50, p_90 and p_99 values for a values slice.
func Percentiles(verbose bool, values []int) [3]int {
	Quicksort(values)

	if verbose {
		fmt.Printf("[INFO] Sorted values : %+v\n", values)
	}

	var out = [3]int{}
	out[0] = values[int(float64(50.0/100)*float64(len(values)-1))]
	out[1] = values[int(float64(90.0/100)*float64(len(values)-1))]
	out[2] = values[int(float64(99.0/100)*float64(len(values)-1))]

	if verbose {
		fmt.Printf("[INFO] Found percentiles %+v\n", out)
	}

	return out
}
