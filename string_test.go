// Copyright ©2022 Foolin. All rights reserved.

package gotab

import (
	"testing"
	"unicode/utf8"
)

var cases = []string{
	"Hello",
	"HelloWorld",
	"Hello, World",
	"Hello，World",
	"Hello，世界",
	"你好, 世界",
	"您好，世界",
	"Hello Go World",
	"您好Go World",
	"您好Go语言世界!",
	"您好, Go世界",
	"您好, Go语言世界!",
}

func TestStrWidth(t *testing.T) {
	for _, s := range cases {
		println(s, "W:", StrWidth(s))
		println(s, "L:", len(s))
		println(s, "R:", utf8.RuneCountInString(s))
	}
}

func TestStrPadding(t *testing.T) {
	n := 24
	println(Padding("Hello", "-", 12+n) + "|")
	println(Padding("Hello, World", "-", 12+n) + "|")
	println(Padding("Hello界", "-", 12+n) + "|")
	println(Padding("Hello世界", "-", 11+n) + "|")
	println(Padding("Hello，世界", "-", 10+n) + "|")
	println(Padding("您好，世界", "-", 9+n) + "|")
	println(Padding("您好中世界", "-", 9+n) + "|")
}

func TestStrPadding2(t *testing.T) {
	for i := 14; i < 30; i++ {
		println("\n\n------ Width:", i, " ------")
		for _, s := range cases {
			println(Padding(s, "-", i) + "|")
		}

	}
}
