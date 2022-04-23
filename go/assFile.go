package asu

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func OpenAssFile(path string) (*AssFile, error) {
// 	assFile := NewAssFile()

// 	file, err := os.Open(path)

// 	if err != nil {
// 		return nil, NewError("error opening AssFile: %v", err)
// 	}

// 	defer file.Close()

// 	scriptInfoActivo := false
// 	aegisubProjectGarbageActivo := false
// 	stylesActivo := false
// 	graphicsActivo := false
// 	fontsActivo := false
// 	eventsActivo := false
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		l := scanner.Text()

// 		switch l {
// 		case "[Script Info]":
// 			scriptInfoActivo = true
// 			continue
// 		case "[Aegisub Project Garbage]":
// 			scriptInfoActivo = false
// 			aegisubProjectGarbageActivo = true
// 			continue
// 		case "[V4+ Styles]":
// 			aegisubProjectGarbageActivo = false
// 			stylesActivo = true
// 			continue
// 		case "[Graphics]":
// 			stylesActivo = false
// 			graphicsActivo = true
// 			continue
// 		case "[Fonts]":
// 			fontsActivo = false
// 			eventsActivo = true
// 			continue
// 		default:
// 			break
// 		}

// 		processLine(assFile, l)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return nil, NewError("error parsing AssFile: %v", err)
// 	}

// 	return assFile, nil
// }

// func processLine(assFile *AssFile, l string) {

// }

// type AssFile struct {
// 	ScriptInfo            ScriptInfo
// 	AegisubProjectGarbage AegisubProjectGarbage
// 	Fonts                 []AttachedFont
// 	Graphics              []AttachedGraphic
// 	Styles                []Style
// 	Events                []Line
// 	Directory             string
// 	Name                  string
// 	Extension             string
// }

// func NewAssFile() *AssFile {
// 	f := &AssFile{}
// 	return f
// }

// type WrapStyle int

// const (
// 	SmartWrappingTopLineWidder    WrapStyle = 0
// 	EndLineWordWrapping           WrapStyle = 1
// 	NoWordWrapping                WrapStyle = 2
// 	SmartWrappingBottomLineWidder WrapStyle = 3
// )

// type YCbCrMatrix int

// const (
// 	None    = 0
// 	TV_601  = 1
// 	PC_601  = 2
// 	TV_709  = 3
// 	PC_709  = 4
// 	TV_FCC  = 5
// 	PC_FCC  = 6
// 	TV_240M = 7
// 	PC_240M = 8
// )

// type ScriptInfo struct {
// 	Comments              []string
// 	Title                 string
// 	ScriptType            string
// 	WrapStyle             WrapStyle
// 	PlayResX              int
// 	PlayResY              int
// 	ScaledBorderAndShadow bool
// 	YCbCrMatrix           YCbCrMatrix
// 	Others                []string
// }

// type AegisubProjectGarbage struct {
// 	AutomationScripts string
// 	ExportFilters     string
// 	ExportEncoding    string
// 	LastStyleStorage  string
// 	AudioFile         string
// 	VideoFile         string
// 	TimecodesFile     string
// 	KeyframesFile     string
// 	VideoZoomPercent  float64
// 	ScrollPosition    int
// 	ActiveLine        int
// 	VideoPosition     int
// 	VideoARMode       int
// 	VideoARValue      float64
// 	Others            []string
// }

// type AttachedFileType string

// const (
// 	FONT    AttachedFileType = "fontname"
// 	GRAPHIC AttachedFileType = "filename"
// )

// type attachedFile struct {
// 	Name string
// 	Data []string
// }

// func attachedFileToString(fileType AttachedFileType, file attachedFile) string {
// 	s := fmt.Sprintf("%s: %s", fileType, file.Name)

// 	for _, d := range file.Data {
// 		s += "\n" + d
// 	}

// 	return s
// }

// type AttachedFont struct {
// 	attachedFile
// }

// func (f *AttachedFont) AttachedFileType() AttachedFileType {
// 	return FONT
// }

// func (f *AttachedFont) ToString() string {
// 	return attachedFileToString(f.AttachedFileType(), f.attachedFile)
// }

// type AttachedGraphic struct {
// 	attachedFile
// }

// func (f *AttachedGraphic) AttachedFileType() AttachedFileType {
// 	return GRAPHIC
// }

// func (f *AttachedGraphic) ToString() string {
// 	return attachedFileToString(f.AttachedFileType(), f.attachedFile)
// }

// // TODO: implementar tag
// type TagTypeColor string

// type Style struct {
// 	Name           string
// 	FontName       string
// 	FontSize       float64
// 	PrimaryColor   TagTypeColor
// 	SecondaryColor TagTypeColor
// 	OutlineColor   TagTypeColor
// 	BackColor      TagTypeColor
// 	PrimaryAlpha   int
// 	SecondaryAlpha int
// 	OutlineAlpha   int
// 	BackAlpha      int
// 	Bold           bool
// 	Italic         bool
// 	Underline      bool
// 	StrikeOut      bool
// 	ScaleX         float64
// 	ScaleY         float64
// 	Spacing        float64
// 	Angle          float64
// 	BorderStyle    bool
// 	Outline        float64
// 	Shadow         float64
// 	Alignment      Alignment
// 	MarginLeft     int
// 	MarginRight    int
// 	MarginVertical int
// 	Encoding       Encoding
// }

// type Alignment int

// const (
// 	DownLeft     = 1
// 	DownCenter   = 2
// 	DownRight    = 3
// 	CenterLeft   = 4
// 	CenterCenter = 5
// 	CenterRight  = 6
// 	UpLeft       = 7
// 	UpCenter     = 8
// 	UpRight      = 9
// )

// type Encoding int

// // TODO: usar nombres en inglés
// func (e Encoding) ToString() string {
// 	switch e {
// 	case ENCODE_ANSI:
// 		return "ANSI"
// 	case ENCODE_DEFAULT:
// 		return "Predeterminado"
// 	case ENCODE_SYMBOL:
// 		return "Símbolo"
// 	case ENCODE_MAC:
// 		return "Mac"
// 	case ENCODE_SHIFT_JIS:
// 		return "Shift_JIS"
// 	case ENCODE_HANGEUL:
// 		return "Hangeul"
// 	case ENCODE_JOHAB:
// 		return "Johab"
// 	case ENCODE_GB2312:
// 		return "GB2312"
// 	case ENCODE_CHINESE_BIG5:
// 		return "Chino BIG5"
// 	case ENCODE_GREEK:
// 		return "Griego"
// 	case ENCODE_TURKISH:
// 		return "Turco"
// 	case ENCODE_VIETNAMESE:
// 		return "Vietnamita"
// 	case ENCODE_HEBREW:
// 		return "Hebre"
// 	case ENCODE_ARAB:
// 		return "Árabe"
// 	case ENCODE_BALTIC:
// 		return "Báltico"
// 	case ENCODE_RUSSIAN:
// 		return "Ruso"
// 	case ENCODE_THAI:
// 		return "Tailandés"
// 	case ENCODE_EASTERN_EUROPE:
// 		return "Europa del Este"
// 	case ENCODE_OEM:
// 		return "OEM"
// 	default:
// 		return ""
// 	}
// }

// const (
// 	ENCODE_ANSI           Encoding = 0
// 	ENCODE_DEFAULT        Encoding = 1
// 	ENCODE_SYMBOL         Encoding = 2
// 	ENCODE_MAC            Encoding = 77
// 	ENCODE_SHIFT_JIS      Encoding = 128
// 	ENCODE_HANGEUL        Encoding = 129
// 	ENCODE_JOHAB          Encoding = 130
// 	ENCODE_GB2312         Encoding = 134
// 	ENCODE_CHINESE_BIG5   Encoding = 136
// 	ENCODE_GREEK          Encoding = 161
// 	ENCODE_TURKISH        Encoding = 162
// 	ENCODE_VIETNAMESE     Encoding = 163
// 	ENCODE_HEBREW         Encoding = 177
// 	ENCODE_ARAB           Encoding = 178
// 	ENCODE_BALTIC         Encoding = 186
// 	ENCODE_RUSSIAN        Encoding = 204
// 	ENCODE_THAI           Encoding = 222
// 	ENCODE_EASTERN_EUROPE Encoding = 238
// 	ENCODE_OEM            Encoding = 255
// )
