package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"regexp"
)

// advancementRegex is the regex used to parse advancement achievement in the log file
var advancementRegex = regexp.MustCompile(` ?(?P<player>\S*?) has made the advancement \[(?P<advancement>.*)]`)
var goalRegex = regexp.MustCompile(` ?(?P<player>\S*?) has reached the goal \[(?P<goal>.*)]`)
var challengeRegex = regexp.MustCompile(` ?(?P<player>\S*?) has completed the challenge \[(?P<challenge>.*)]`)

// ParseAdvancement parse the line using the advancementRegex
// and returns the Message struct corresponding
func ParseAdvancement(line string) message.Message {
	result := mapRegexGroupMatch(advancementRegex, line)
	if result.IsSuccess() {
		return newAdvancementMessageFromRegex(result)
	}
	result = mapRegexGroupMatch(goalRegex, line)
	if result.IsSuccess() {
		return newGoalMessageFromRegex(result)
	}

	result = mapRegexGroupMatch(challengeRegex, line)
	if result.IsSuccess() {
		return newChallengeMessageFromRegex(result)
	}
	return message.AdvancementMessage{}
}

// newAdvancementMessageFromRegex returns the Message struct corresponding to an achievement from the parsed regex groups
func newAdvancementMessageFromRegex(regexResult RegexResult) message.AdvancementMessage {
	return message.NewAdvancementMessage(regexResult["player"], regexResult["advancement"])
}

// newGoalMessageFromRegex returns the Message struct corresponding to an achievement from the parsed regex groups
func newGoalMessageFromRegex(regexResult RegexResult) message.AdvancementMessage {
	return message.NewGoalMessage(regexResult["player"], regexResult["goal"])
}

// newChallengeMessageFromRegex returns the Message struct corresponding to an achievement from the parsed regex groups
func newChallengeMessageFromRegex(regexResult RegexResult) message.AdvancementMessage {
	return message.NewChallengeMessage(regexResult["player"], regexResult["challenge"])
}
