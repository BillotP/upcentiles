package upfluence

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"upcentile/internal/api"
)

// UpfluenceStreamAddress is declared as a variable in order to use a mock server for testing.
var UpfluenceStreamAddress = "https://stream.upfluence.co/stream"

// SubscribeFull to UpfluenceStreamAddress SSE stream for a duration and send deserialized events to ch chan.
func SubscribeFull(duration time.Duration, ch chan *StreamEvent) error {
	req, err := http.NewRequest("GET", UpfluenceStreamAddress, nil)
	if err != nil {
		return fmt.Errorf("creating request to %s failed : %w", UpfluenceStreamAddress, err)
	}
	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("executing request to %s failed : %w", UpfluenceStreamAddress, err)
	}
	defer close(ch)
	deadline := time.Now().Add(duration)
	for time.Now().Before(deadline) {
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			jsonb := strings.Replace(string(line), "data: ", "", 1)
			var event = new(StreamEvent)
			if err = json.Unmarshal([]byte(jsonb), event); err != nil {
				return fmt.Errorf("failed to unmarshall upfluence sse data: %w", err)
			}
			ch <- event
			break
		}
	}
	return nil
}

// GetAnalysisValues return timestamp and dimension values from a StreamEvent.
func GetAnalysisValues(el *StreamEventV2, dimension api.AnalysisDimension) *AnalysisValue {
	switch dimension {
	case api.Likes:
		if el.Likes != nil {
			return &AnalysisValue{el.Timestamp, *el.Likes}
		}
	case api.Comments:
		if el.Comments != nil {
			return &AnalysisValue{el.Timestamp, *el.Comments}
		}
	case api.Favorites:
		if el.Favorites != nil {
			return &AnalysisValue{el.Timestamp, *el.Favorites}
		}
	case api.Retweets:
		if el.Retweets != nil {
			return &AnalysisValue{el.Timestamp, *el.Retweets}
		}
	}

	return nil
}

// SubscribeLight to UpfluenceStreamAddress SSE stream for a duration and send only required analysis values to ch chan.
func SubscribeLight(verbose bool, duration time.Duration, dimension api.AnalysisDimension, ch chan *AnalysisValue) error {
	defer close(ch)

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", UpfluenceStreamAddress, nil)

	if err != nil {
		return fmt.Errorf("creating request to %s failed : %w", UpfluenceStreamAddress, err)
	}

	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("executing request to %s failed : %w", UpfluenceStreamAddress, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response got status %s expected %d", resp.Status, http.StatusOK)
	}

	deadline := time.Now().Add(duration)

	for time.Now().Before(deadline) {
		scanner := bufio.NewScanner(resp.Body)
		if scanner.Scan() {
			line := scanner.Text()
			jsonb := strings.Replace(line, "data: ", "", 1)

			var event = new(map[string]*StreamEventV2)

			if err = json.Unmarshal([]byte(jsonb), event); err != nil {
				log.Printf("[WARNING] Failed to deserialize event [%s] : %s", jsonb, err.Error())
				continue
			}

			for key, val := range *event {
				if verbose {
					log.Printf("[INFO] Received event %s\n", key)
				}
				if values := GetAnalysisValues(val, dimension); values != nil {
					ch <- values
				}
			}
		}
	}

	return nil
}
