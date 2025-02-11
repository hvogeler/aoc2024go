package onsen

import (
	"bufio"
	"strings"
)

type Pattern string

func (p Pattern) String() string {
	return string(p)
}

type Design struct {
	design     string
	isPossible bool
}

type Onsen struct {
	pattern map[Pattern]bool
	designs []Design
	longestPatternLength int
}

func (o Onsen) ContainsPattern(p Pattern) bool {
	_, exists := o.pattern[p]
	return exists
}

func NewOnsen() Onsen {
	return Onsen{
		pattern: make(map[Pattern]bool, 0),
		designs: []Design{},
	}
}

func OnsenFromStr(s string) Onsen {
	onsen := NewOnsen()
	scanner := bufio.NewScanner(strings.NewReader(s))
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		if row == 0 {
			patterns := strings.Split(line, ", ")
			for _, pattern := range patterns {
				onsen.pattern[Pattern(pattern)] = true
			}
		}
		if row > 1 {
			onsen.designs = append(onsen.designs, Design{
				design:     line,
				isPossible: false,
			})
		}
	}
	return onsen
}
