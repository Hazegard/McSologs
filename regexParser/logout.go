package regexParser

import (
	"github.com/Hazegard/McSoLogs/structs/message"
	"regexp"
)

var logoutRegex = regexp.MustCompile(` (?P<player>\S*?) left the game`)

func ParseLogout(line string) message.Message {
	_, result := mapRegexGroupMatch(logoutRegex, line)
	return newLogoutMessageFromRegex(result)
}

func newLogoutMessageFromRegex(regexResult map[string]string) message.LogoutMessage {
	return message.NewLogoutMessage(regexResult["player"])
}
