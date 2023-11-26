package contracts

type SMS interface {
	Send(phone string, message Message) error
}

type Driver interface {
	Send(phone string, message Message, config map[string]string) error
}

type Message struct {
	Data    map[string]string
	Content string
}
