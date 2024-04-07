package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"testing"
)

func TestParseGoal(t *testing.T) {
	input := "[20:48:28] [Server thread/INFO]: TestPlayer has reached the goal [Sky's the Limit]"
	ans := ParseAdvancement(input)

	wants := message.AdvancementMessage{
		Player:      "TestPlayer",
		Advancement: "Sky's the Limit",
		Type:        message.GOAL,
	}

	if ans != wants {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}
}

func TestNewGoalMessageFromRegex(t *testing.T) {
	input := "[20:48:28] [Server thread/INFO]: TestPlayer has reached the goal [Sky's the Limit]"
	ans := newGoalMessageFromRegex(mapRegexGroupMatch(goalRegex, input))

	wants := message.AdvancementMessage{
		Player:      "TestPlayer",
		Advancement: "Sky's the Limit",
		Type:        message.GOAL,
	}

	if ans != wants {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}
}

func TestParseAdvancement(t *testing.T) {
	input := "[20:48:28] [Server thread/INFO]: TestPlayer has made the advancement [The End?]"
	ans := ParseAdvancement(input)

	wants := message.AdvancementMessage{
		Player:      "TestPlayer",
		Advancement: "The End?",
		Type:        message.ADVANCEMENT,
	}

	if ans != wants {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}
}

func TestNewAdvancementMessageFromRegex(t *testing.T) {
	input := "[20:48:28] [Server thread/INFO]: TestPlayer has made the advancement [The End?]"
	ans := newAdvancementMessageFromRegex(mapRegexGroupMatch(advancementRegex, input))

	wants := message.AdvancementMessage{
		Player:      "TestPlayer",
		Advancement: "The End?",
		Type:        message.ADVANCEMENT,
	}

	if ans != wants {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}
}

func TestParseChallenge(t *testing.T) {
	input := "[20:48:28] [Server thread/INFO]: TestPlayer has completed the challenge [Great View From Up Here]"
	ans := ParseAdvancement(input)

	wants := message.AdvancementMessage{
		Player:      "TestPlayer",
		Advancement: "Great View From Up Here",
		Type:        message.CHALLENGE,
	}

	if ans != wants {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}
}

func TestNewChallengeMessageFromRegex(t *testing.T) {
	input := "[20:48:28] [Server thread/INFO]: TestPlayer has completed the challenge [Great View From Up Here]"
	ans := newChallengeMessageFromRegex(mapRegexGroupMatch(challengeRegex, input))

	wants := message.AdvancementMessage{
		Player:      "TestPlayer",
		Advancement: "Great View From Up Here",
		Type:        message.CHALLENGE,
	}

	if ans != wants {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}
}

func TestParseAdvancementInvalid(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer left the game"
	ans := ParseAdvancement(input)

	var wants message.AdvancementMessage
	if wants != ans {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}

}

func TestParseAdvancementInvalid2(t *testing.T) {
	input := "[16:39:15] [Server thread/WARN]: Mismatch in destroy block pos: class_2338{x=71, y=105, z=501} class_2338{x=70, y=104, z=503}"
	ans := ParseAdvancement(input)

	var wants message.AdvancementMessage
	if wants != ans {
		t.Errorf("ParseAdvancement() = %+v, want %+v", ans, wants)
	}

}
