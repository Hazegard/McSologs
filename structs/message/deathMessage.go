package message

// DeathMessage holds the parsed message when a death occurs
type DeathMessage struct {
	Message    string
	DeadPlayer string
	Killer     string
}

// NewDeathMessage return the struct holding the parsed message
func NewDeathMessage(name string, killer string, message string) Message {
	return DeathMessage{Message: message, DeadPlayer: name, Killer: killer}
}

// IsEmpty returns whether the message is empty
func (m DeathMessage) IsEmpty() bool {
	return m.Message == "" && m.DeadPlayer == ""
}

// GetMessage returns the message corresponding to a death
func (m DeathMessage) GetMessage() string {
	return mapPlayer(m.Message, m.Killer, m.DeadPlayer)
}

// GetTitle returns the title used by the discord notification corresponding to a player logging out
func (m DeathMessage) GetTitle() string {
	return "RIP"
}

// GetWHColor returns the color used bu the discord notification corresponding to a player logging out
func (m DeathMessage) GetWHColor() string {
	return "0xDC143C"
}
