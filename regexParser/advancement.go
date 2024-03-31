package regexParser

import (
	"github.com/Hazegard/McSoLogs/structs/message"
	"regexp"
)

// advancementRegex is the regex used to parse advancement achievement in the log file
var advancementRegex = regexp.MustCompile(` (?P<player>\S*?) has made the advancement \[(?P<advancement>.*)]`)

// ParseAdvancement parse the line using the advancementRegex
// and returns the Message struct corresponding
func ParseAdvancement(line string) message.Message {
	result := mapRegexGroupMatch(advancementRegex, line)
	return newAdvancementMessageFromRegex(result)
}

// newAdvancementMessageFromRegex returns the Message struct corresponding to an achievement from the parsed regex groups
func newAdvancementMessageFromRegex(regexResult RegexResult) message.AdvancementMessage {
	return message.NewAdvancementMessage(regexResult["player"], regexResult["advancement"])
}
