package regexParser

import "regexp"

// mapRegexGroupMatch parse the line using the provided regex
// and return a RegexResult structs that holds a map
// containing the regex group name as key and the matched group as value
func mapRegexGroupMatch(regex *regexp.Regexp, line string) RegexResult {
	result := RegexResult{}

	match := regex.FindStringSubmatch(line)

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" && len(match) > 0 {
			result[name] = match[i]
		}
	}
	return result
}

// RegexResult holds the map containing the matched groups in the regex
type RegexResult map[string]string

// IsSuccess return whether the regex matched the line
func (r *RegexResult) IsSuccess() bool {
	return len(*r) != 0
}
