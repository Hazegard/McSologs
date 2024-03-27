package RegexParser

import "regexp"

func mapRegexGroupMatch(regex *regexp.Regexp, line string) (bool, map[string]string) {
	result := make(map[string]string)

	match := regex.FindStringSubmatch(line)
	if len(match) == 0 {
		return false, result
	}
	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" && len(match) > 0 {
			result[name] = match[i]
		}
	}
	return true, result
}
