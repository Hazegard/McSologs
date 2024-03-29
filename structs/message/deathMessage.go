package message

import (
	"fmt"
)

type DeathMessage struct {
	Message    string
	DeadPlayer string
}

func NewDeathMessage(name, message string) Message {
	return DeathMessage{Message: message, DeadPlayer: name}
}

func (m DeathMessage) IsEmpty() bool {
	return m.Message == "" && m.DeadPlayer == ""
}

func (m DeathMessage) GetMessage() string {
	return fmt.Sprintf("GG %s!\n%s", m.DeadPlayer, m.Message)
}
