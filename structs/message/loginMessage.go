package message

import (
	"fmt"
)

type LoginMessage struct {
	Player string
}

func NewLoginMessage(player string) LoginMessage {
	return LoginMessage{Player: player}
}

func (m LoginMessage) IsEmpty() bool {
	return m.Player == ""
}

func (m LoginMessage) GetMessage() string {
	return fmt.Sprintf("%s joined the game", m.Player)
}
