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
	mes := message
	for _, player := range players {
		mes = strings.ReplaceAll(mes, player, fmt.Sprintf("<@%s>", playerMapper.Get(player)))
		fmt.Println(mes)
	}
	return mes
}
