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
	message := fmt.Sprintf("%s has made the advancement: %s", m.Player, m.getAdvancementUrl())
	return MapPlayer(message, m.Player)
}

func (m AdvancementMessage) GetTitle() string {
	return "Advancement"
}

func (m AdvancementMessage) GetWHColor() string {
	return "0x00FFFF"
}

func (m AdvancementMessage) getAdvancementUrl() string {
	return fmt.Sprintf("[%s](https://minecraft.fandom.com/wiki/Advancement#%s)", m.Advancement, strings.ReplaceAll(m.Advancement, " ", "_"))
}
