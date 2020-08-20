package spinner

import (
	"strings"
	"sync"
	"unicode/utf8"
)

// Spinner contains the settings of a spinner
type Spinner struct {
	style        string
	currentIndex int

	spinChars    []string
	spinCharsMtx sync.RWMutex
}

// Next returns the next character for the spinner
func (s *Spinner) Next() string {
	s.loadIfUnloaded()
	s.spinCharsMtx.RLock()
	s.currentIndex++
	if s.currentIndex >= len(s.spinChars) {
		s.currentIndex = 0
	}
	s.spinCharsMtx.RUnlock()
	return s.Current()
}

// Current returns the current character for the spinner
func (s *Spinner) Current() string {
	s.loadIfUnloaded()
	s.spinCharsMtx.RLock()
	defer s.spinCharsMtx.RUnlock()
	// just in case someone changed the style mid-way
	if s.currentIndex >= len(s.spinChars) {
		s.currentIndex = 0
	}

	return SpinnerStyles[s.style][s.currentIndex]
}

// SetStyle loads a style into the spinner
func (s *Spinner) SetStyle(style string) {
	if style == s.style {
		return
	}
	s.spinCharsMtx.Lock()
	defer s.spinCharsMtx.Unlock()

	s.style = style
	spinnerStyleMtx.RLock()
	s.spinChars = SpinnerStyles[s.style]
	spinnerStyleMtx.RUnlock()
}

// Clear returns the amount of characters for the first spinner-state in spaces
// in order to clear the spinner if required.
func (s *Spinner) Clear() string {
	return strings.Repeat(" ", utf8.RuneCountInString(SpinnerStyles[s.style][0]))
}

// loadIfUnloaded makes sure a style is always loaded.
func (s *Spinner) loadIfUnloaded() {
	s.spinCharsMtx.Lock()
	if len(s.spinChars) > 0 {
		s.spinCharsMtx.Unlock()
		return
	}

	spinnerStyleMtx.RLock()
	s.spinChars = SpinnerStyles[""]
	s.spinCharsMtx.Unlock()
	spinnerStyleMtx.RUnlock()
}
