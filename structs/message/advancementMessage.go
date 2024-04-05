package message

import (
	"fmt"
	"strings"
)

type AdvancementType int

const (
	GOAL AdvancementType = iota
	ADVANCEMENT
)

// AdvancementMessage holds the parsed message when an advancement is achieved by a player
type AdvancementMessage struct {
	Player      string
	Advancement string
	Type        AdvancementType
}

// NewAdvancementMessage return the struct holding the parsed message
func NewAdvancementMessage(name, advancement string) AdvancementMessage {
	return AdvancementMessage{Advancement: advancement, Player: name, Type: ADVANCEMENT}
}

// NewGoalMessage return the struct holding the parsed message
func NewGoalMessage(name, advancement string) AdvancementMessage {
	return AdvancementMessage{Advancement: advancement, Player: name, Type: GOAL}
}

// IsEmpty returns whether the message is empty
func (m AdvancementMessage) IsEmpty() bool {
	return m.Player == ""
}

// GetMessage returns the message corresponding to a player achieving an advancement
func (m AdvancementMessage) GetMessage() string {
	switch m.Type {
	case ADVANCEMENT:
		message := fmt.Sprintf("%s has made the advancement: %s", m.Player, m.getAdvancementUrl())
		return mapPlayer(message, m.Player)
	case GOAL:
		message := fmt.Sprintf("%s has reached the goal: %s", m.Player, m.getAdvancementUrl())
		return mapPlayer(message, m.Player)
	}
	return ""
}

// GetTitle returns the title used by the discord notification corresponding to a player achieving an advancement
func (m AdvancementMessage) GetTitle() string {
	return "Advancement"
}

// GetWHColor returns the color used by the discord notification corresponding to a player achieving an advancement
func (m AdvancementMessage) GetWHColor() string {
	return "0x00FFFF"
}

// getAdvancementUrl returns the advancement URL in the Minecraft Wiki
func (m AdvancementMessage) getAdvancementUrl() string {
	return fmt.Sprintf(
		"[%s](https://minecraft.fandom.com/wiki/Advancement#%s)",
		m.Advancement,
		strings.ReplaceAll(m.Advancement, " ", "_"),
	)
}
