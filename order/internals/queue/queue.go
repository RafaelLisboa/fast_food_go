package queue


type QueuePublisher interface {
	SendMessage([]byte) error
}



