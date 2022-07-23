package asu

type LineType string

const (
	LINE_COMMENT  LineType = "Comment"
	LINE_DIALOGUE LineType = "Dialogue"
)

func (s LineType) IsValid() bool {
	return s == LINE_COMMENT || s == LINE_DIALOGUE
}
