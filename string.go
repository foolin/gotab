// Copyright ©2022 Foolin. All rights reserved.

package gotab

import (
	"github.com/mattn/go-runewidth"
	"regexp"
	"strings"
)

// Ignore color code
var ansi = regexp.MustCompile("\033\\[(?:[0-9]{1,3}(?:;[0-9]{1,3})*)?[m|K]")

func StrWidth(str string) int {
	return runewidth.StringWidth(ansi.ReplaceAllLiteralString(str, ""))
}

// StrTruncate return string truncated with w cells
func StrTruncate(s string, w int, tail string) string {
	return runewidth.Truncate(s, w, tail)
}

func Padding(s, pad string, width int) string {
	sw := StrWidth(s)
	gap := width - sw
	if gap > 0 {
		return s + strings.Repeat(pad, gap)
	}
	return s
}

func PaddingTab(s, pad string, width int) string {
	width = fixTabWidth(width)
	sw := StrWidth(s)
	tab := "\t"
	gap := width - sw
	if gap > 0 {
		return s + strings.Repeat(pad, gap) + tab
	}
	return s + tab
}

func fixTabWidth(width int) int {
	return width - (width % 4) + 2 //fix tab width
}
