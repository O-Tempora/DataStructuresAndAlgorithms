package main

import "strings"

// TabFormatter replaces spaces with tabs.
type TabFormatter interface {
	SpaceToTab(s string) string
}

// Tabber implements TabFormatter
type Tabber struct {
	trimSpaces bool
}

// SpaceToTab implementation of TabFormatter.
func (n Tabber) SpaceToTab(s string) string {
	if n.trimSpaces {
		s = strings.TrimSpace(s)
	}
	return strings.ReplaceAll(s, " ", "\t")
}
