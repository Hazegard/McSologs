package regexParser

import (
	"github.com/Hazegard/McSoLogs/structs/message"
	"regexp"
)

var advancementRegex = regexp.MustCompile(` (?P<player>\S*?) has made the advancement \[(?P<advancement>.*)]`)

func ParseAdvancement(line string) message.Message {
	_, result := mapRegexGroupMatch(advancementRegex, line)
	return newAdvancementMessageFromRegex(result)
}

func newAdvancementMessageFromRegex(regexResult map[string]string) message.AdvancementMessage {
	return message.NewAdvancementMessage(regexResult["player"], regexResult["advancement"])
}
