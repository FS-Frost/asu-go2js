package asu

import (
	"regexp"
)

const (
	preTag  = `\{[^\}]*`
	postTag = `[^\{]*\}`
)

func ReplaceTag(text string, tagType TagType, newTag string) string {
	tagRegex := GetTagRegex(tagType)
	inGroupTagRegex := preTag + tagRegex + postTag
	re := regexp.MustCompile(inGroupTagRegex)
	index := re.FindStringSubmatchIndex(text)
	groups := re.SubexpNames()
	step := 2
	var start int
	var end int

	for i := 0; i < len(index); i += step {
		group := groups[i/step]

		if group != "tag" {
			continue
		}

		start = index[i]
		end = index[i+1]
		break
	}

	newText := text[:start] + newTag + text[end:]
	return newText
}
