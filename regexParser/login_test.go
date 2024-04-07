package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"testing"
)

func TestParseLogin(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer joined the game"
	ans := ParseLogin(input)

	wants := message.LoginMessage{Player: "TestPlayer"}

	if wants != ans {
		t.Errorf("ParseLogin() = %+v, want %+v", ans, wants)
	}
}

func TestNewLoginMessageFromRegex(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer joined the game"
	ans := newLoginMessageFromRegex(mapRegexGroupMatch(loginRegex, input))

	wants := message.LoginMessage{Player: "TestPlayer"}
	if wants != ans {
		t.Errorf("newLoginMessageFromRegex() = %+v, want %+v", ans, wants)
	}
}

func TestParseLoginInvalid(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer left the game"
	ans := ParseLogin(input)

	var wants message.LoginMessage
	if wants != ans {
		t.Errorf("ParseLogin() = %+v, want %+v", ans, wants)
	}

}

func TestParseLoginInvalid2(t *testing.T) {
	input := "[16:39:15] [Server thread/WARN]: Mismatch in destroy block pos: class_2338{x=71, y=105, z=501} class_2338{x=70, y=104, z=503}"
	ans := ParseLogin(input)

	var wants message.LoginMessage
	if wants != ans {
		t.Errorf("ParseLogin() = %+v, want %+v", ans, wants)
	}

}
