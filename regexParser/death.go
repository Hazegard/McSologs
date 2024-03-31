package regexParser

import (
	"github.com/Hazegard/McSoLogs/structs/message"
	"regexp"
	"strings"
)

var deathRegexes []*regexp.Regexp

func init() {
	for _, text := range deathMessages {
		deathRegexes = append(deathRegexes, regexp.MustCompile(`: (?P<message>`+text+`)`))
		deathRegexes = append(deathRegexes, regexp.MustCompile(`: '(?P<message>`+text+`)'`))
	}
}

func newDeathMessageFromRegex(regexResult map[string]string) message.Message {
	return message.NewDeathMessage(regexResult["dead_player"], regexResult["killer"], regexResult["message"])
}

func ParseDeath(line string) message.Message {
	var mes message.Message
	if strings.Contains(line, "Villager class") {
		return mes
	}
	for _, regexe := range deathRegexes {
		isFound, result := mapRegexGroupMatch(regexe, line)
		if isFound {
			mes = newDeathMessageFromRegex(result)
			return mes
		}
	}
	return mes
}
