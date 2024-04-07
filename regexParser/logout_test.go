package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"testing"
)

func TestParseLogout(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer left the game"
	ans := ParseLogout(input)

	wants := message.LogoutMessage{Player: "TestPlayer"}

	if wants != ans {
		t.Errorf("ParseLogout() = %+v, want %+v", ans, wants)
	}
}

func TestNewLogoutMessageFromRegex(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer left the game"
	ans := newLogoutMessageFromRegex(mapRegexGroupMatch(logoutRegex, input))

	wants := message.LogoutMessage{Player: "TestPlayer"}
	if wants != ans {
		t.Errorf("newLogoutMessageFromRegex() = %+v, want %+v", ans, wants)
	}
}

func TestParseLogoutInvalid(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer joined the game"
	ans := ParseLogout(input)

	var wants message.LogoutMessage
	if wants != ans {
		t.Errorf("ParseLogout() = %+v, want %+v", ans, wants)
	}

}

func TestParseLogoutInvalid2(t *testing.T) {
	input := "[16:39:15] [Server thread/WARN]: Mismatch in destroy block pos: class_2338{x=71, y=105, z=501} class_2338{x=70, y=104, z=503}"
	ans := ParseLogout(input)

	var wants message.LogoutMessage
	if wants != ans {
		t.Errorf("ParseLogout() = %+v, want %+v", ans, wants)
	}

}
