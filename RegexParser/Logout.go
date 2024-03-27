package RegexParser

import (
	"mc_so_logs/structs/Message"
	"regexp"
)

var logoutRegex = regexp.MustCompile(` (?P<player>\S*?) left the game`)

func ParseLogout(line string) Message.Message {
	_, result := mapRegexGroupMatch(logoutRegex, line)
	return newLogoutMessageFromRegex(result)
}

func newLogoutMessageFromRegex(regexResult map[string]string) Message.LogoutMessage {
	return Message.NewLogoutMessage(regexResult["player"])
}
