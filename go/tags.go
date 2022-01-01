package asu

type TagType string

const (
	TAG_B    TagType = "b"
	TAG_POS  TagType = "pos"
	TAG_MOVE TagType = "move"
)

type namedTag interface {
	Name() string
}

type TagTypeInt interface {
	namedTag
	Argument() int
}

func TagExists(text string, tag TagType) bool {
	rule := GetTagRegex(tag)

	if rule == "" {
		return false
	}

	return RegexPass(rule, text)
}

func GetTagRegex(tag TagType) string {
	switch tag {
	case TAG_B:
		return regexTagB
	case TAG_POS:
		return regexPos
	default:
		panic("Invalid tag: " + tag)
	}
}

func formatTagParsingError(tag namedTag, err error) error {
	return NewError("error parsing %s tag: %v", tag.Name(), err)
}
