package progress

import (
	"errors"
	"math"
	"strings"
	"sync"
)

// Progress contains the settings of a Progressbar.
type Progress struct {
	style string
	Width int

	prgChars    []string
	prgCharsMtx sync.RWMutex
}

// SetStyle sets the style that should be used with the progressbar
func (p *Progress) SetStyle(style string) error {
	if style == p.style {
		return nil
	}
	progressStyleMtx.RLock()
	defer progressStyleMtx.RUnlock()

	prgc, exists := ProgressStyles[style]
	if !exists {
		return errors.New("undefined style")
	}

	p.prgCharsMtx.Lock()
	p.prgChars = prgc
	p.prgCharsMtx.Unlock()
	p.style = style
	return nil
}

// GetBar returns the string that represents the progressbar in the set style
// for the given progress
func (p *Progress) GetBar(parts, total float64) (result string) {
	if p.Width <= 0 {
		p.Width = 10
	}
	p.prgCharsMtx.RLock()
	defer p.prgCharsMtx.RUnlock()
	maxIndex := len(p.prgChars) - 1
	if total == 0 {
		result += strings.Repeat(p.prgChars[maxIndex], p.Width)
		return
	}

	percperchar := float64(100) / float64(p.Width)
	percent, _ := GetPercentage(parts, total)
	counter := 0

	if percent > 100 {
		percent = float64(100)
	} else if percent < 0 {
		percent = float64(0)
	}

	for percent > percperchar {
		result += p.prgChars[maxIndex]

		percent -= percperchar
		counter++
	}

	if counter == p.Width {
		return
	}

	charindex := 0
	dingeling := percperchar / float64(maxIndex+1)
	temp := math.Floor(percent / dingeling)
	for percent > dingeling {
		percent -= dingeling
		charindex++
	}
	charindex = int(temp)

	if charindex > 0 {
		if charindex > maxIndex {
			charindex = maxIndex
		}
		result += p.prgChars[charindex]
		counter++
	}

	for counter < p.Width {
		result += p.prgChars[0]
		counter++
	}
	return
}
