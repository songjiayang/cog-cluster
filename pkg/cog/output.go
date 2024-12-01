package cog

type Output struct {
	TaskID      string      `json:"id"`
	Output      interface{} `json:"output"`
	Status      string      `json:"status"`
	StartedAt   string      `json:"started_at"`
	CompletedAt string      `json:"completed_at"`
}

func (o *Output) IsSuccess() bool {
	return o.Status == "succeeded"
}
