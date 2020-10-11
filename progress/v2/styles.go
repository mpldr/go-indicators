package progress

import (
	"sync"

	"gitlab.com/poldi1405/go-ansi"
)

// ProgressStyles contains the styles of the progressbars. Custom styles can
// easily be appended and used.
var ProgressStyles = map[string]ProgressStyle{
	"double":        ProgressStyle{Bar: []string{"="}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
	"double-":       ProgressStyle{Bar: []string{"-", "="}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
	"single":        ProgressStyle{Bar: []string{"-"}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
	"parallelogram": ProgressStyle{Bar: []string{"▰"}, Empty: []string{"▱"}},
	"spaced-blocks": ProgressStyle{Bar: []string{"▮"}, Empty: []string{"▯"}},
	"block":         ProgressStyle{Bar: []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}, Empty: []string{" "}, FrameStart: "▕", FrameEnd: "▏"},
	"pacman":        ProgressStyle{Bar: []string{"-"}, Empty: []string{" ", "o"}, Head: []string{"c", ansi.Yellow("C")}, FrameStart: "[", FrameEnd: "]", EmptyStyle: func(o string) string { return ansi.Red(o) }, HeadStyle: func(c string) string { return ansi.Yellow(c) }}
	"":              ProgressStyle{Bar: []string{"="}, Empty: []string{" "}, FrameStart: "[", FrameEnd: "]"},
}

type ProgressStyle struct {
	Bar            []string
	AlternatingBar bool

	Empty []string

	Head []string

	FrameStart string
	FrameEnd   string

	BarStyle   func(char string) string
	HeadStyle  func(char string) string
	EmptyStyle func(char string) string
}

// progressStyleMtx locks the map so there are no concurrent map-accesses
var progressStyleMtx sync.RWMutex

// DefineStyle allows registering a custom progress-style to use. If the name
// is already defined it is overwritten. Directly defining it is possible but
// not recommended.
func DefineStyle(name string, characters []string) {
	progressStyleMtx.Lock()
	defer progressStyleMtx.Unlock()

	ProgressStyles[name] = characters
}
