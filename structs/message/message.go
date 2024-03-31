package message

import (
	"fmt"
	"github.com/Hazegard/McSoLogs/config"
	"strings"
)

type Message interface {
	GetMessage() string
	IsEmpty() bool
	GetWHColor() string
	GetTitle() string
}

func MapPlayer(message string, players ...string) string {
	playerMapper := config.GetConfig().Players
	mes := ""
	for _, player := range players {
		mes = strings.ReplaceAll(message, player, fmt.Sprintf("<@%s>", playerMapper.Get(player)))
	}
	return mes
}
