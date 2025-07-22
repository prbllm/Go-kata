package actionpurpose

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	separator = ":"
)

type EventType struct {
	Author   string
	Aim      string
	Verb     string
	Negative bool
	Action   string
}

type ActionPurposeParser struct {
	resultAction map[string]int
	lastAction   string
	reg          *regexp.Regexp
}

func NewActionPurposeParser() *ActionPurposeParser {
	return &ActionPurposeParser{
		resultAction: make(map[string]int),
		reg:          regexp.MustCompile(`^\s*([A-Za-z]+)\s+(is|am)\s+(?:(not)\s+)?([A-Za-z]+)!`),
	}
}

func (a *ActionPurposeParser) ParseLine(line string) error {
	parts := strings.SplitN(line, separator, 2)

	if len(parts) != 2 {
		return fmt.Errorf("message %v without separator %v", line, separator)
	}
	var event EventType
	event.Author = strings.TrimSpace(parts[0])

	regResults := a.reg.FindStringSubmatch(parts[1])
	if regResults == nil {
		return fmt.Errorf("unable to parse seccond part of message %v", parts[1])
	}

	event.Aim, event.Verb, a.lastAction = regResults[1], regResults[2], regResults[4]

	if len(regResults[3]) == 0 {
		event.Negative = false
	} else if regResults[3] == "not" {
		event.Negative = true
	} else {
		return fmt.Errorf("unable to check negative or not: %v", regResults[3])
	}
	if event.Aim == "I" && event.Verb == "am" {
		event.Aim = event.Author
	}

	var result int
	if event.Author == event.Aim {
		if event.Negative {
			result = -1
		} else {
			result = 2
		}
	} else {
		if event.Negative {
			result = -1
		} else {
			result = 1
		}
	}

	if _, ok := a.resultAction[event.Author]; !ok {
		a.resultAction[event.Author] = 0
	}
	if _, ok := a.resultAction[event.Aim]; !ok {
		a.resultAction[event.Aim] = 0
	}
	a.resultAction[event.Aim] += result
	return nil
}

func (a *ActionPurposeParser) GetResult() []string {
	var (
		max   int
		keys  []string
		first = true
	)

	for k, v := range a.resultAction {
		switch {
		case first || v > max:
			max, keys, first = v, []string{k}, false
		case v == max:
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	return keys
}

func (a *ActionPurposeParser) CleanData() {
	a.resultAction = make(map[string]int)
	a.lastAction = ""
}
