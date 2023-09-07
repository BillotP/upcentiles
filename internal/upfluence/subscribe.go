package upfluence

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const UpfluenceStreamAddress = "https://stream.upfluence.co/stream"

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
		for scanner.Scan() {
			line := scanner.Text()
			jsonb := strings.Replace(string(line), "data: ", "", 1)
			var event = new(StreamEvent)
			if err = json.Unmarshal([]byte(jsonb), event); err != nil {
				return err
			}
			ch <- event
			break
		}
	}
	return nil
}
