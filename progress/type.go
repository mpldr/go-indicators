package progress

import (
	"strings"
	"sync"
)

type Progress struct {
	Style string
	Width int

	mtx sync.Mutex
}

func (p *Progress) GetBar(parts, total float64) (result string) {
	if p.Width <= 0 {
		p.Width = 10
	}
	maxIndex := len(ProgressStyles[p.Style]) - 1
	if total == 0 {
		result += strings.Repeat(ProgressStyles[p.Style][maxIndex], p.Width)
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
		result += ProgressStyles[p.Style][maxIndex]

		percent -= percperchar
		counter++
	}

	if counter == p.Width {
		return
	}

	charindex := 0
	dingeling := percperchar / float64(maxIndex+1)
	for percent > dingeling {
		percent -= dingeling
		charindex++
	}

	if charindex > 0 {
		result += ProgressStyles[p.Style][charindex]
		counter++
	}

	for counter < p.Width {
		result += ProgressStyles[p.Style][0]
		counter++
	}
	return
}
