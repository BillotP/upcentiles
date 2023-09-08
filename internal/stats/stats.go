package stats

import (
	"upcentile/internal/api"
	"upcentile/internal/upfluence"

	"golang.org/x/exp/slices"
)

// Percentiles compute the percentiles  p_50, p_90 and p_99 values for a particular AnalysisDimension of StreamEvent datas.
func Percentiles(dimension api.AnalysisDimension, datas []*upfluence.StreamEvent) [3]uint64 {
	var values = []uint64{}

	for _, el := range datas {
		switch dimension {
		case api.Comments:
			values = append(values, el.Comments())
		case api.Favorites:
			values = append(values, el.Favorites())
		case api.Retweets:
			values = append(values, el.Retweets())

		default:
			values = append(values, el.Likes())
		}
	}
	// If no lib allowed, sort might be implement using a partition sort for small list,
	// and quicksort for longest
	slices.Sort(values)
	var out = [3]uint64{}
	out[0] = values[int(float64(50.0/100)*float64(len(datas)-1))]
	out[1] = values[int(float64(90.0/100)*float64(len(datas)-1))]
	out[2] = values[int(float64(99.0/100)*float64(len(datas)-1))]
	return out
}
