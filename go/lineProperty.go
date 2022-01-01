package asu

type LineProperty string

const (
	LINE_TYPE            LineProperty = "type"
	LINE_LAYER           LineProperty = "layer"
	LINE_START           LineProperty = "start"
	LINE_END             LineProperty = "end"
	LINE_STYLE           LineProperty = "style"
	LINE_ACTOR           LineProperty = "actor"
	LINE_MARGIN_LEFT     LineProperty = "marginL"
	LINE_MARGIN_RIGHT    LineProperty = "marginR"
	LINE_MARGIN_VERTICAL LineProperty = "marginV"
	LINE_EFFECT          LineProperty = "effect"
	LINE_CONTENT         LineProperty = "content"
)

const regexLineProp = `(?P<type>(?:Dialogue|Comment)): (?P<layer>\d),(?P<start>\d*:\d{2}:\d{2}\.\d{2}),(?P<end>\d:\d{2}:\d{2}\.\d{2}),(?P<style>[^,]*),(?P<actor>[^,]*),(?P<marginL>\d*),(?P<marginR>\d*),(?P<marginV>\d*),(?P<effect>[^,]*),(?P<content>.*)`
