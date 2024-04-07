package regexParser

import (
	"github.com/Hazegard/mcsologs/structs/message"
	"regexp"
	"testing"
)

func TestParseDeath(t *testing.T) {
	input := "[17:14:01] [Server thread/INFO]: TestPlayer was slain by Vex"
	ans := ParseDeath(input)

	wants := message.DeathMessage{
		Message:    "TestPlayer was slain by Vex",
		DeadPlayer: "TestPlayer",
		Killer:     "Vex",
	}

	if wants != ans {
		t.Errorf("ParseDeath() = %+v, want %+v", ans, wants)
	}
}

func TestIgnoreVillagerDeath(t *testing.T) {
	input := "[15:58:17] [Server thread/INFO]: Villager class_1646['Villager'/442, l='ServerLevel[world]', x=-182.77, y=67.00, z=110.80] died, message: 'Villager was slain by Zombie'"
	ans := ParseDeath(input)

	var wants message.Message

	if wants != ans {
		t.Errorf("ParseDeath() = %+v, want %+v", ans, wants)
	}
}

func TestNewDeathMessageFromRegex(t *testing.T) {
	input := "[17:14:01] [Server thread/INFO]: TestPlayer left the confines of this world while fighting TestPlayer2"
	ans := newDeathMessageFromRegex(
		mapRegexGroupMatch(
			regexp.MustCompile(`: (?P<message>(?P<dead_player>.*?) left the confines of this world while fighting (?P<killer>.*))`),
			input),
	)

	wants := message.DeathMessage{
		Message:    "TestPlayer left the confines of this world while fighting TestPlayer2",
		DeadPlayer: "TestPlayer",
		Killer:     "TestPlayer2",
	}
	if wants != ans {
		t.Errorf("newDeathMessageFromRegex() = %+v, want %+v", ans, wants)
	}
}
