package request

type QueueRequest struct {
	URL  string      `json:"url" validate:"required,url"`
	Data interface{} `json:"data" validate:"required"`
}
