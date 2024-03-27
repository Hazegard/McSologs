package Message

type Message interface {
	GetMessage() string
	IsEmpty() bool
}
