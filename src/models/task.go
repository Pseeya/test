package models

const (
	StatusDone        = "Done"
	StatuseProcessing = "In process"
	StatusError       = "Error"
)

type Task struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}
