package message

import (
	"fmt"
)

// LogoutMessage holds the parsed message when a player log out
type LogoutMessage struct {
	Player string
}

// NewLogoutMessage return the struct holding the parsed message
func NewLogoutMessage(player string) LogoutMessage {
	return LogoutMessage{Player: player}
}

// IsEmpty returns whether the message is empty
func (m LogoutMessage) IsEmpty() bool {
	return m.Player == ""
}

// GetMessage returns the message corresponding to a player logging out
func (m LogoutMessage) GetMessage() string {
	return fmt.Sprintf("%s left the game", m.Player)
}

// GetTitle returns the title used by the discord notification corresponding to a player logging out
func (m LogoutMessage) GetTitle() string {
	return "Goodbye"
}

// GetWHColor returns the color used yu the discord notification corresponding to a player logging out
func (m LogoutMessage) GetWHColor() string {
	return "0xFF8C00"
}
