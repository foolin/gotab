// Copyright ©2022 Foolin. All rights reserved.

package gotab

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"strings"
)

func StrWidth(str string) int {
	return runewidth.StringWidth(str)
	//return runewidth.StringWidth(ansi.ReplaceAllLiteralString(str, ""))
}

// StrTruncate return string truncated with w cells
func StrTruncate(s string, w int, tail string) string {
	return runewidth.Truncate(s, w, tail)
}

func StrPadding(s, pad string, width int) string {
	//func Align(s, pad string, width int) string {
	format := fmt.Sprint("%-", width, "v")
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
