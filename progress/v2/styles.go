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
