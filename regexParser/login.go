package regexParser

import (
	"github.com/Hazegard/McSoLogs/structs/message"
	"regexp"
)

var loginRegex = regexp.MustCompile(` (?P<player>\S*?) joined the game`)

func newLoginMessageFromRegex(regexResult map[string]string) message.LoginMessage {
	return message.NewLoginMessage(regexResult["player"])
}

func ParseLogin(line string) message.Message {
	_, result := mapRegexGroupMatch(loginRegex, line)
	return newLoginMessageFromRegex(result)
}
