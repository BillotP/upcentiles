package handler

import (
	"log"
	"net/http"
	"time"
	"upcentile/internal/api"
	"upcentile/internal/stats"
	"upcentile/internal/upfluence"

	"github.com/labstack/echo/v4"
)

// GetAnalysisHandler is a JSON api handler for GET /analysis route with *duration* and *dimension* query parameters.
func GetAnalysisHandler(c echo.Context) error {
	var params = new(api.AnalysisParam)
	if err := c.Bind(params); err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	if err := params.Validate(); err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	dd, _ := time.ParseDuration(params.Duration)
	cc := make(chan *upfluence.StreamEvent)
	go func() error {
		if err := upfluence.Subscribe(dd, cc); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		return nil
	}()

	var datas = []*upfluence.StreamEvent{}

	var out = api.AnalysisResponse{}

	for ev := range cc {
		datas = append(datas, ev)
	}
	out.TotalPosts = uint64(len(datas) - 1)
	out.MinTimestamp = datas[0].Timestamp()
	out.MaxTimestamp = datas[len(datas)-1].Timestamp()
	out.Fill(params.Dimension, stats.Percentiles(params.Dimension, datas))
	return c.JSON(http.StatusOK, out)
}
