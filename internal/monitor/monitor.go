package monitor

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	enabled = true
	file    *os.File
)

func SetEnabled(v bool) {
	mu.Lock()
	defer mu.Unlock()
	enabled = v
}

func IsEnabled() bool {
	mu.Lock()
	defer mu.Unlock()
	return enabled
}

func RecordEvent(e AMFEvent) error {
	mu.Lock()
	defer mu.Unlock()

	if !enabled {
		return nil
	}

	if e.Timestamp.IsZero() {
		e.Timestamp = time.Now().UTC()
	}

	if file == nil {
		var err error
		file, err = os.OpenFile("amf_event_trace.jsonl", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			return err
		}
	}

	encoder := json.NewEncoder(file)
	return encoder.Encode(e)
}
