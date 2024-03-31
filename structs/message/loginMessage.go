package message

import (
	"fmt"
)

// LoginMessage holds the parsed message when a player log in
type LoginMessage struct {
	Player string
}

// NewLoginMessage return the struct holding the parsed message
func NewLoginMessage(player string) LoginMessage {
	return LoginMessage{Player: player}
}

// IsEmpty returns whether the message is empty
func (m LoginMessage) IsEmpty() bool {
	return m.Player == ""
}

// GetMessage returns the message corresponding to a player logging in
func (m LoginMessage) GetMessage() string {
	return fmt.Sprintf("%s joined the game", m.Player)
}

// GetTitle returns the title used by the discord notification corresponding to a player logging in
func (m LoginMessage) GetTitle() string {
	return "Welcome"
}

// GetWHColor returns the color used by the discord notification corresponding to a player logging in
func (m LoginMessage) GetWHColor() string {
	return "0x7CFC00"
}
