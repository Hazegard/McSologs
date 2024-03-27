package RegexParser

import (
	"mc_so_logs/structs/Message"
	"regexp"
)

var loginRegex = regexp.MustCompile(` (?P<player>\S*?) joined the game`)

func newLoginMessageFromRegex(regexResult map[string]string) Message.LoginMessage {
	return Message.NewLoginMessage(regexResult["player"])
}

func ParseLogin(line string) Message.Message {
	_, result := mapRegexGroupMatch(loginRegex, line)
	return newLoginMessageFromRegex(result)
}
