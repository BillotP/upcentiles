package upfluence

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// UpfluenceStreamAddress is declared as a variable in order to use a mock server for testing.
var UpfluenceStreamAddress = "https://stream.upfluence.co/stream"

// Subscribe to UpfluenceStreamAddress SSE stream for a duration and send message to ch chan.
func Subscribe(duration time.Duration, ch chan *StreamEvent) error {
	req, err := http.NewRequest("GET", UpfluenceStreamAddress, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer close(ch)
	deadline := time.Now().Add(duration)
	for time.Now().Before(deadline) {
		scanner := bufio.NewScanner(resp.Body)
		// Read til '\n' (newline) control character
		for scanner.Scan() {
			line := scanner.Text()
			// Not sure about that dirty hack to get a valid JSON string
			jsonb := strings.Replace(string(line), "data: ", "", 1)
			var event = new(StreamEvent)
			if err = json.Unmarshal([]byte(jsonb), event); err != nil {
				// Also not sure about this kind of error handling
				return err
			}
			ch <- event
			break
		}
	}
	return nil
}
