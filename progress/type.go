package progress

import (
	"math"
	"strings"
	"sync"
)

type Progress struct {
	style string
	width int

	mtx sync.Mutex
}

func (p *Progress) GetBar(parts, total float64) (result string) {
	maxIndex := len(ProgressStyles[p.style]) - 1
	if total == 0 {
		result += strings.Repeat(ProgressStyles[p.style][maxIndex], p.width)
		return
	}

	// don't question it, I used excel to come up with this
	percent, _ := GetPercentage(parts, total)
	percentPerCharacter := math.Floor(float64(100) / float64(p.width))
	completeSegments := math.Floor(percent / percentPerCharacter)
	percentRemaining := percent - completeSegments*percentPerCharacter
	percentPerSegment := percentPerCharacter / float64(maxIndex+1)
	// TODO: make this not negative multimillion, maybe we should question it...
	character := int(math.Round(percentRemaining / percentPerSegment))

	result += strings.Repeat(ProgressStyles[p.style][maxIndex], int(completeSegments))

	result += ProgressStyles[p.style][character]
	result += strings.Repeat(ProgressStyles[p.style][0], p.width-int(completeSegments)-1)
	return
}
