package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"regexp"
)

// loginRegex is the regex used to parse player log in in the log file
var loginRegex = regexp.MustCompile(` (?P<player>\S*?) joined the game`)

// newLoginMessageFromRegex returns the Message struct corresponding to an achievement
// from the parsed regex groups
func newLoginMessageFromRegex(regexResult RegexResult) message.LoginMessage {
	return message.NewLoginMessage(regexResult["player"])
}

// ParseLogin parse the line using the loginRegex
// and returns the Message struct corresponding
func ParseLogin(line string) message.Message {
	result := mapRegexGroupMatch(loginRegex, line)
	return newLoginMessageFromRegex(result)
}
