package progress

import (
	"fmt"
	"math"
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

	a2 := p.Width
	b2, _ := GetPercentage(parts, total)
	c2 := float64(len(ProgressStyles[p.Style]))
	d2 := float64(100 / a2)
	completeSegments := math.Floor(b2 / d2)
	e2 := completeSegments
	f2 := b2 - d2*e2
	g2 := d2 / c2
	segmentIndex := int(math.Floor(f2 / g2))

	if b2 >= 100 {
		result += strings.Repeat(ProgressStyles[p.Style][maxIndex], p.Width)
		return
	}

	fmt.Println(completeSegments)
	result += strings.Repeat(ProgressStyles[p.Style][maxIndex], int(completeSegments))

	if completeSegments < float64(p.Width) {
		result += ProgressStyles[p.Style][segmentIndex]
	}

	if completeSegments+1 < float64(p.Width) {
		result += strings.Repeat(ProgressStyles[p.Style][0], p.Width-int(completeSegments)-1)
	}
	return
}
