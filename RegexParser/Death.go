package RegexParser

import (
	"mc_so_logs/structs/Message"
	"regexp"
)

var deathRegexes []*regexp.Regexp

func init() {
	for _, text := range deathMessages {
		deathRegexes = append(deathRegexes, regexp.MustCompile(`: (?P<message>`+text+`)`))
		deathRegexes = append(deathRegexes, regexp.MustCompile(`: '(?P<message>`+text+`)'`))
	}
}

func newDeathMessageFromRegex(regexResult map[string]string) Message.Message {
	return Message.NewDeathMessage(regexResult["dead_player"], regexResult["message"])
}

func ParseDeath(line string) Message.Message {
	var mes Message.Message
	for _, regexe := range deathRegexes {
		isFound, result := mapRegexGroupMatch(regexe, line)
		if isFound {
			mes = newDeathMessageFromRegex(result)
			return mes
		}
	}
	return mes
}
