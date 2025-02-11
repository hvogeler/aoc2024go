package onsen

import (
	"bufio"
	"fmt"
	"strings"
)

type Pattern string

func (p Pattern) String() string {
	return string(p)
}

type Design struct {
	design        string
	designPattern []Pattern
	isPossible    bool
}

func (d Design) IsPatternPossible() bool {
	return d.isPossible
}

func (d Design) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("\nDesign: %s\n", d.design))
	s.WriteString(fmt.Sprintf("Is Possible to create? %v\n", d.isPossible))
	for i, pat := range d.designPattern {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(pat.String())
	}
	s.WriteString("\n")
	return s.String()
}

type Onsen struct {
	pattern              map[Pattern]bool
	designs              []*Design
	longestPatternLength int
}

func (o Onsen) Designs() []*Design {
	return o.designs
}

func (o *Onsen) CheckDesigns() {
	for _, design := range o.designs {
		o.CheckDesign(design)
	}
}

func (o *Onsen) CheckDesign(d *Design) *Design {
	patLenToCheck := o.LongestPatternLength()
	isPatternPossible := false
	cursor := 0
	for cursor < len(d.design) {
		if cursor + patLenToCheck > len(d.design) {
			patLenToCheck = len(d.design) - cursor
		}
		isPatternPossible = false
		PatLoop:
		for patLen := patLenToCheck; patLen > 0; patLen-- {
			substr := d.design[cursor:(cursor + patLen)]
			if o.ContainsPattern(Pattern(substr)) {
				d.designPattern = append(d.designPattern, Pattern(substr))
				isPatternPossible = true
				cursor += patLen
				break PatLoop
			}
		}
		if !isPatternPossible {
			d.isPossible = false
			return d
		}
	}
	d.isPossible = true
	return d
}

func (o Onsen) CountDesigns() (int, int, int) {
	var possible int
	var impossible int
	for _, design := range o.designs {
		if design.isPossible {
			possible++
		} else {
			impossible++
		}
	}
	return len(o.designs), possible, impossible
}

func (o Onsen) ContainsPattern(p Pattern) bool {
	_, exists := o.pattern[p]
	return exists
}

func (o Onsen) LongestPatternLength() int {
	return o.longestPatternLength
}

func NewOnsen() Onsen {
	return Onsen{
		pattern: make(map[Pattern]bool, 0),
		designs: []*Design{},
	}
}

func OnsenFromStr(s string) Onsen {
	onsen := NewOnsen()
	scanner := bufio.NewScanner(strings.NewReader(s))
	maxPatLen := 0
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		if row == 0 {
			patterns := strings.Split(line, ", ")
			for _, pattern := range patterns {
				onsen.pattern[Pattern(pattern)] = true
				if len(pattern) > maxPatLen {
					maxPatLen = len(pattern)
				}
			}
		}
		if row > 1 {
			onsen.designs = append(onsen.designs, &Design{
				design:     line,
				isPossible: false,
			})
		}
	}
	onsen.longestPatternLength = maxPatLen
	return onsen
}
