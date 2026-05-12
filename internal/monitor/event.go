package monitor

import "time"

type AMFEvent struct {
	Timestamp          time.Time `json:"timestamp"`
	Layer              string    `json:"layer"`
	EventType          string    `json:"event_type"`
	UeKey              string    `json:"ue_key,omitempty"`
	Supi               string    `json:"supi,omitempty"`
	Guti               string    `json:"guti,omitempty"`
	RanUeNgapId        int64     `json:"ran_ue_ngap_id,omitempty"`
	AmfUeNgapId        int64     `json:"amf_ue_ngap_id,omitempty"`
	RawNgapUeId        string    `json:"raw_ngap_ue_id,omitempty"`
	PduSessionId       int32     `json:"pdu_session_id,omitempty"`
	Procedure          string    `json:"procedure,omitempty"`
	ContextAction      string    `json:"context_action,omitempty"`
	ContextAvailable   bool      `json:"context_available,omitempty"`
	RetentionCandidate bool      `json:"retention_candidate,omitempty"`
	ReleaseReason      string    `json:"release_reason,omitempty"`
	SourceFile         string    `json:"source_file,omitempty"`
	Note               string    `json:"note,omitempty"`
}
