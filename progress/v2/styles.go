package progress

import (
	"sync"

	"gitlab.com/poldi1405/go-ansi"
)

// ProgressStyles contains the styles of the progressbars. Custom styles can
// easily be appended and used.
var ProgressStyles = map[string]Style{
	"double":          Style{Bar: []string{"="}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
	"double-":         Style{Bar: []string{"-", "="}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
	"single":          Style{Bar: []string{"-"}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
	"parallelogram":   Style{Bar: []string{"▰"}, Empty: []string{"▱"}},
	"spaced-blocks":   Style{Bar: []string{"▮"}, Empty: []string{"▯"}},
	"braille":         Style{Bar: []string{"⠁", "⠃", "⠇", "⠏", "⠟", "⠿"}, Empty: []string{" "}, FrameStart: "(", FrameEnd: ")"},
	"block":           Style{Bar: []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}, Empty: []string{" "}, FrameStart: "▕", FrameEnd: "▏"},
	"pacman":          Style{Bar: []string{"-"}, Empty: []string{" ", "o"}, Head: []string{"c", ansi.Yellow("C")}, FrameStart: "[", FrameEnd: "]", EmptyStyle: func(o string) string { return ansi.Red(o) }, HeadStyle: func(c string) string { return ansi.Yellow(c) }},
	"wget":            Style{Bar: []string{"="}, Empty: []string{" "}, Head: []string{">"}, FrameStart: "[", FrameEnd: "]"},
	"doublesinglebar": Style{Bar: []string{"=", "-"}, AlternatingBar: true, Empty: []string{" "}, Head: []string{">"}, FrameStart: "[", FrameEnd: "]"},
	"":                Style{Bar: []string{"="}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
}

// Style acts as a container for the progressbar design
type Style struct {
	Bar            []string // The characters used to construct the bar
	AlternatingBar bool     // The bar is not incremental like `block` but like `doublesinglebar`

	Empty []string // The characters used to construct the unfilled part of the progressbar

	Head []string // The character replacing the last one of the bar, if multiple are given they will be alternated

	FrameStart string // A character to print before the bar
	FrameEnd   string // A character to print after the unfilled part

	BarStyle   func(char string) string // A function to apply to Bar characters
	HeadStyle  func(char string) string // A function to apply to Head characters
	EmptyStyle func(char string) string // A function to apply to Empty-ish(?) characters
}

// progressStyleMtx locks the map so there are no concurrent map-accesses
var progressStyleMtx sync.RWMutex

// DefineStyle allows registering a custom progress-style to use. If the name
// is already defined it is overwritten. Directly defining it is possible but
// not recommended.
func DefineStyle(name string, pb Style) {
	progressStyleMtx.Lock()
	defer progressStyleMtx.Unlock()

	ProgressStyles[name] = pb
}
