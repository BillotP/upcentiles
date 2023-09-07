package stats

import (
	"fmt"
	"upcentile/internal"
	"upcentile/internal/upfluence"

	"golang.org/x/exp/slices"
)

/**
Percentiles can be calculated using the formula
n = (P/100) x N,
where P = percentile,
N = number of values in a data set (sorted from smallest to largest),
and n = ordinal rank of a given value.
*/

// type EventSet []*upfluence.StreamEvent

func Percentiles(dimension internal.AnalysisDimension, datas []*upfluence.StreamEvent) [3]uint64 {
	var values = []uint64{}

	for _, el := range datas {
		switch dimension {
		case internal.Comments:
			values = append(values, el.Comments())
		case internal.Favorites:
			values = append(values, el.Favorites())
		case internal.Retweets:
			values = append(values, el.Retweets())

		default:
			values = append(values, el.Likes())
		}
	}
	fmt.Printf("Unsorted :\n%+v\n", values)
	// If no lib allowed, sort might be implement using a partition sort for small list,
	// and quicksort for longest
	slices.Sort(values)
	fmt.Printf("Sorted :\n%+v\n", values)
	var out = [3]uint64{}
	out[0] = values[int(float64(50.0/100)*float64(len(datas)-1))]
	out[1] = values[int(float64(90.0/100)*float64(len(datas)-1))]
	out[2] = values[int(float64(99.0/100)*float64(len(datas)-1))]
	return out
}
