package progress

import (
	"container/ring"
	"strings"
)

func barBuilder(percentage float64, width int, s Style) string {
	if s.BarStyle == nil {
		s.BarStyle = stringReturner
	}

	if s.HeadStyle == nil {
		s.HeadStyle = stringReturner
	}

	if s.EmptyStyle == nil {
		s.EmptyStyle = stringReturner
	}

	var bar string
	charactercount := 0
	if s.AlternatingBar {
		bar = getAlternatingBar(s, charactercount)
	} else {
		bar = strings.Repeat(s.Bar[len(s.Bar)-1], charactercount)
	}
}

func getAlternatingBar(s Style, length int) string {
	chars := len(s.Bar)

	charRing := ring.New(chars)
	for _, c := range s.Bar {
		charRing.Value = s.BarStyle(c)
		r = r.Next()
	}

	var result string
	for i := 0; i < length; i++ {
		result += charRing.Value
		charRing = r.Next()
	}

	return result
}

// use this if no StyleFunction is set
func stringReturner(s string) string {
	return s
}
