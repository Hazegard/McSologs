package regexParser

import (
	"testing"
)

func TestMapRegexGroupMatchLogout(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer left the game"
	ans := mapRegexGroupMatch(logoutRegex, input)

	wants := map[string]string{
		"player": "TestPlayer",
	}

	if len(wants) != len(ans) {
		t.Errorf("TestMapRegexGroupMatchLogout failed, expect one element, got %d (%v)", len(wants), ans)
	}

	if _, ok := ans["player"]; !ok {
		t.Errorf("TestMapRegexGroupMatchLogout failed, key \"player\" not found, got %v", ans)

	}

	if ans["player"] != wants["player"] {
		t.Errorf("TestMapRegexGroupMatchLogout failed, expect: %v, got: %v", wants, ans)
	}
}

func TestMapRegexGroupMatchLogint(t *testing.T) {
	input := "[20:22:17] [Server thread/INFO]: TestPlayer joined the game"
	ans := mapRegexGroupMatch(loginRegex, input)

	wants := map[string]string{
		"player": "TestPlayer",
	}

	if len(wants) != len(ans) {
		t.Errorf("TestMapRegexGroupMatchLogout failed, expect one element, got %d (%v)", len(wants), ans)
	}

	if _, ok := ans["player"]; !ok {
		t.Errorf("TestMapRegexGroupMatchLogout failed, key \"player\" not found, got %v", ans)

	}

	if ans["player"] != wants["player"] {
		t.Errorf("TestMapRegexGroupMatchLogout failed, expect: %v, got: %v", wants, ans)
	}
}

func TestMapRegexGroupMatchAdvancement(t *testing.T) {
	input := "[20:48:28] [Server thread/INFO]: TestPlayer has completed the challenge [Great View From Up Here]"
	ans := mapRegexGroupMatch(challengeRegex, input)

	wants := map[string]string{
		"challenge": "Great View From Up Here",
		"player":    "TestPlayer",
	}

	if len(wants) != len(ans) {
		t.Errorf("TestMapRegexGroupMatchLogout failed, expect one element, got %d (%v)", len(wants), ans)
	}

	if _, ok := ans["player"]; !ok {
		t.Errorf("TestMapRegexGroupMatchLogout failed, key \"player\" not found, got %v", ans)

	}

	if ans["player"] != wants["player"] {
		t.Errorf("TestMapRegexGroupMatchLogout failed, expect: %v, got: %v", wants, ans)
	}
}
