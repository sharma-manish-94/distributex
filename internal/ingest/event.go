package ingest

import "time"

type Event struct {
	EventID   string         `json:"event_id"`  // unique ID
	TenantID  string         `json:"tenant_id"` // who sent the event
	Timestamp time.Time      `json:"timestamp"` // when did the event happen
	Name      string         `json:"event_name"`
	Payload   map[string]any `json:"payload"`
}
