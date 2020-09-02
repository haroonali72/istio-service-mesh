package types

type WorkQueueConsumer struct {
	ApplicationId string `json:"app_id"`
	Token         string `json:"token"`
	Action        string `json:"action"`
}
type WorkQueuePublisher struct {
	ApplicationId string `json:"app_id"`
	Token         string `json:"token"`
	Action        string `json:"action"`
	Status        bool   `json:"status"`
	Message       string `json:"message"`
}
