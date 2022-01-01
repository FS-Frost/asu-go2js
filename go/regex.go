package asu

import "regexp"

func ApplyRegex(rule string, target string) (result map[string]string) {
	var compRegEx = regexp.MustCompile(rule)
	match := compRegEx.FindStringSubmatch(target)

	result = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			result[name] = match[i]
		}
	}
	return result
}

func RegexPass(rule string, target string) bool {
	results := ApplyRegex(rule, target)
	return len(results) > 0
}
