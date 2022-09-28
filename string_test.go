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
	n := 3
	println(StrPadding("Hello", "-", 12+n) + "|")
	println(StrPadding("Hello, World", "-", 12+n) + "|")
	println(StrPadding("Hello，世界", "-", 10+n) + "|")
	println(StrPadding("您好，世界", "-", 9+n) + "|")
	println(StrPadding("您好中世界", "-", 9+n) + "|")
}

func TestStrPadding2(t *testing.T) {
	for _, s := range cases {
		println(s, "W:", StrWidth(s))
		println(s, "L:", len(s))
	}
	for i := 23; i < 30; i++ {
		println("\n\n------ Width:", i, " ------")
		for _, s := range cases {
			println(StrPadding(s, "-", i) + "|")
		}

	}
}
