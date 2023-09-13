package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"upcentile/internal/api"
	"upcentile/internal/stats"
	"upcentile/internal/upfluence"

	"github.com/labstack/echo/v4"
)

// GetAnalysisHandler is a JSON api handler for GET /analysis route with *duration* and *dimension* query parameters.
func GetAnalysisHandler(verbose bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params = new(api.AnalysisParam)
		if err := c.Bind(params); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		if err := params.Validate(); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// No need to check error as duration has already been check by Validate()
		dd, _ := time.ParseDuration(params.Duration)
		cc := make(chan *upfluence.AnalysisValue)

		if verbose {
			log.Printf("[INFO] Will compute analysis for %s duration on %s dimension\n", params.Duration, params.Dimension)
		}
		go func() error {
			if err := upfluence.SubscribeLight(verbose, dd, params.Dimension, cc); err != nil {
				log.Printf("[ERROR] %s", err.Error())
				// Should return a more succint error to consumer to avoid leaking program  internal errors
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			return nil
		}()

		var timestamps = []int{}

		var values = []int{}

		for ev := range cc {
			timestamps = append(timestamps, ev.Timestamp)
			values = append(values, ev.DimensionValue)
		}

		if verbose {
			fmt.Printf("[INFO] Got %d events with values %+v\n", len(values), values)
		}

		if len(timestamps) == 0 {
			return c.NoContent(http.StatusNoContent)
		}

		var out = api.AnalysisResponse{
			TotalPosts:   len(timestamps) - 1,
			MinTimestamp: timestamps[0],
			MaxTimestamp: timestamps[len(timestamps)-1],
		}

		percentiles := stats.Percentiles(verbose, values)
		out.FillPercentiles(params.Dimension, percentiles)

		return c.JSON(http.StatusOK, out)
	}
}
