package RegexParser

import (
	"mc_so_logs/structs/Message"
	"regexp"
)

var advancementRegex = regexp.MustCompile(` (?P<player>\S*?) has made the advancement \[(?P<advancement>.*)]`)

func ParseAdvancement(line string) Message.Message {
	_, result := mapRegexGroupMatch(advancementRegex, line)
	return newAdvancementMessageFromRegex(result)
}

func newAdvancementMessageFromRegex(regexResult map[string]string) Message.AdvancementMessage {
	return Message.NewAdvancementMessage(regexResult["player"], regexResult["advancement"])
}
