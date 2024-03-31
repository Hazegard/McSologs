package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"regexp"
)

// loginRegex is the regex used to parse player log out in the log file
var logoutRegex = regexp.MustCompile(` (?P<player>\S*?) left the game`)

// ParseLogout parse the line using the logoutRegex
// and returns the Message struct corresponding
func ParseLogout(line string) message.Message {
	result := mapRegexGroupMatch(logoutRegex, line)
	return newLogoutMessageFromRegex(result)
}

// newLogoutMessageFromRegex returns the Message struct corresponding to an achievement
// from the parsed regex groups
func newLogoutMessageFromRegex(regexResult RegexResult) message.LogoutMessage {
	return message.NewLogoutMessage(regexResult["player"])
}
