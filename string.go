// Copyright ©2022 Foolin. All rights reserved.

package gotab

import (
	"fmt"
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

func StrPadding2(s, pad string, width int) string {
	//fmt.Println("W", StrWidth(s), "L", len(s), "R", utf8.RuneCountInString(s))
	//sw := StrWidth(s)
	//rw := utf8.RuneCountInString(s)
	//if sw > rw {
	//	width = width - (sw - rw)
	//}
	//func Align(s, pad string, width int) string {
	format := fmt.Sprint("%-", width, "v")
	//println("format: ", format)
	return fmt.Sprintf(format, s)
}

func StrPadding1(s, pad string, width int) string {
	width += width % 4 //tab
	sw := StrWidth(s)
	gap := width - sw
	tab := "\t"
	//rw := len(s)
	//if sw != rw {
	//	tab = "\t"
	//}
	if gap > 0 {
		return s + strings.Repeat(pad, gap) + tab
	}
	return s + tab
}

func Padding(s, pad string, width int) string {
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
	return width - (width % 4) + 2 //tab fix
}
