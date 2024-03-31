package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"regexp"
	"strings"
)

// advancementRegex is a slice holding the  regexes used to parse the deaths in the log file
var deathRegexes []*regexp.Regexp

// init parses the slice of death regexes and populates the deathRegexes slice
func init() {
	for _, text := range deathMessages {
		deathRegexes = append(deathRegexes, regexp.MustCompile(`: (?P<message>`+text+`)`))
		deathRegexes = append(deathRegexes, regexp.MustCompile(`: '(?P<message>`+text+`)'`))
	}
}

// newDeathMessageFromRegex returns the Message struct corresponding to death from the parsed regex groups
func newDeathMessageFromRegex(regexResult RegexResult) message.Message {
	return message.NewDeathMessage(regexResult["dead_player"], regexResult["killer"], regexResult["message"])
}

// ParseDeath iterates over the deathRegexes slice to find a corresponding regex
// and return the corresponding Message struct
func ParseDeath(line string) message.Message {
	var mes message.Message
	if strings.Contains(line, "Villager class") {
		return mes
	}
	for _, regexe := range deathRegexes {
		result := mapRegexGroupMatch(regexe, line)
		if result.IsSuccess() {
			mes = newDeathMessageFromRegex(result)
			return mes
		}
	}
	return mes
}
