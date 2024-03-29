package message

import (
	"fmt"
	"strings"
)

type AdvancementMessage struct {
	Player      string
	Advancement string
}

func NewAdvancementMessage(name, advancement string) AdvancementMessage {
	return AdvancementMessage{Advancement: advancement, Player: name}
}

func (m AdvancementMessage) IsEmpty() bool {
	return m.Player == ""
}

func (m AdvancementMessage) GetMessage() string {
	return fmt.Sprintf("%s has made the advancement: %s", m.Player, strings.ReplaceAll(m.Advancement, " ", "_"))
}
