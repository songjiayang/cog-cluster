package cog

type Input struct {
	Input               map[string]interface{} `json:"input"`
	Webhook             string                 `json:"webhook,omitempty"`
	WebhookEventsFilter []string               `json:"webhook_events_filter,omitempty"`
	OutputFilePrefix    string                 `json:"output_file_prefix,omitempty"`
}
