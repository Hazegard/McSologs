package message

type DeathMessage struct {
	Message    string
	DeadPlayer string
	Killer     string
}

func NewDeathMessage(name string, killer string, message string) Message {
	return DeathMessage{Message: message, DeadPlayer: name, Killer: killer}
}

func (m DeathMessage) IsEmpty() bool {
	return m.Message == "" && m.DeadPlayer == ""
}

func (m DeathMessage) GetMessage() string {
	return MapPlayer(m.Message, m.Killer, m.DeadPlayer)
}

func (m DeathMessage) GetTitle() string {
	return "RIP"
}

func (m DeathMessage) GetWHColor() string {
	return "0xDC143C"
}
