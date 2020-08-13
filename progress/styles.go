package progress

var ProgressStyles = map[string][]string{
	"double":        []string{" ", "="},
	"double-":       []string{" ", "=", "-"},
	"single":        []string{" ", "-"},
	"trapez":        []string{"▱", "▰"},
	"spaced-blocks": []string{"▯", "▮"},
	"block":         []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"},
	"":              []string{" ", "="},
}
