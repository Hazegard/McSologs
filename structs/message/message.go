package message

type Message interface {
	GetMessage() string
	IsEmpty() bool
}
