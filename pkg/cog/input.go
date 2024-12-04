package cog

import "encoding/json"

type Input struct {
	Input               map[string]interface{} `json:"input"`
	Webhook             string                 `json:"webhook,omitempty"`
	WebhookEventsFilter []string               `json:"webhook_events_filter,omitempty"`
	OutputFilePrefix    string                 `json:"output_file_prefix,omitempty"`
}

func (input *Input) Marshal() []byte {
	data, _ := json.Marshal(input)
	return data
}
