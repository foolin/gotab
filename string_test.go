// Copyright ©2022 Foolin. All rights reserved.

package gotab

import "testing"

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
	}
}

func TestStrPadding(t *testing.T) {
	println(StrPadding("Hello", "-", 12) + "|")
	println(StrPadding("Hello, World", "-", 12) + "|")
	println(StrPadding("Hello，世界", "-", 12) + "|")
	println(StrPadding("您好，世界", "-", 12) + "|")
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
