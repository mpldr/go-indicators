package spinner

import "sync"

// SpinnerStyles contains the styles of spinners
var SpinnerStyles = map[string][]string{
	"braille":           []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	"braille-thin":      []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	"windows-10":        []string{"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀"},
	"line":              []string{"|", "/", "-", "\\"},
	"dots":              []string{".  ", ".. ", "..."},
	"dots-bounce":       []string{".  ", ".. ", "...", " ..", "  .", " ..", "...", ".. "},
	"pipe":              []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
	"grow-block":        []string{"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"},
	"expand-block":      []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "▊", "▋", "▌", "▍", "▎"},
	"tiny-rotation-box": []string{"▖", "▘", "▝", "▗"},
	"circle":            []string{"◜", "◠", "◝", "◞", "◡", "◟"},
	"bouncing-bar":      []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"},
	"bouncing-ball":     []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"},
	"clock":             []string{"🕛 ", "🕐 ", "🕑 ", "🕒 ", "🕓 ", "🕔 ", "🕕 ", "🕖 ", "🕗 ", "🕘 ", "🕙 ", "🕚 "},
	"":                  []string{"|", "/", "-", "\\"},
}

// spinnerStyleMtx locks the map so there are no concurrent map-accesses
var spinnerStyleMtx sync.RWMutex

// DefineStyle allows registering a custom spinner-style to use. If the name is
// already defined it is overwritten. Directly defining them is possible but
// not recommended.
func DefineStyle(name string, characters []string) {
	spinnerStyleMtx.Lock()
	defer spinnerStyleMtx.Unlock()

	SpinnerStyles[name] = characters
}
