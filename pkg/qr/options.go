package qr

type LevelBy string

const (
	L = "L"
	M = "M"
	Q = "Q"
	H = "H"
)

type Options struct {
	URL           string  `json:"url,omitempty"`
	Size          int     `json:"size,omitempty"`
	Level         LevelBy `json:"level,omitempty"`
	FgColor       string  `json:"fgColor,omitempty"`
	BgColor       string  `json:"bgColor,omitempty"`
	IncludeMargin bool    `json:"includeMargin,omitempty"`
}
