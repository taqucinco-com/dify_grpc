package dify

type EventData struct {
	Event  string  `json:"event"`
	TaskID string  `json:"task_id"`
	Answer *string `json:"answer"`
}

func send(ch chan<- string) {
	close(ch)
}
