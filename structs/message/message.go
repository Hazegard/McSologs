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

// mapPLayer returns the input message and replace the player Minecraft username by the
// corresponding discord ID
func mapPlayer(message string, players ...string) string {
	playerMapper := config.GetConfig().Players
	mes := message
	for _, player := range players {
		mes = strings.ReplaceAll(mes, player, playerMapper.Get(player))
		fmt.Println(mes)
	}
	return mes
}
