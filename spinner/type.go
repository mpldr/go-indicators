package spinner

import (
	"strings"
	"sync"
	"unicode/utf8"
)

type Spinner struct {
	style        string
	currentIndex int

	mtx sync.Mutex
}

func (s *Spinner) Next() string {
	s.mtx.Lock()
	s.currentIndex++
	if s.currentIndex >= len(SpinnerStyles[s.style]) {
		s.currentIndex = 0
	}
	s.mtx.Unlock()
	return s.Current()
}

func (s *Spinner) Current() string {
	s.mtx.Lock()
	if s.currentIndex >= len(SpinnerStyles[s.style]) {
		s.currentIndex = 0
	}
	s.mtx.Unlock()

	return SpinnerStyles[s.style][s.currentIndex]
}

func (s *Spinner) SetStyle(style string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.style = style
}

func (s *Spinner) Clear() string {
	return strings.Repeat(" ", utf8.RuneCountInString(SpinnerStyles[s.style][0]))
}
