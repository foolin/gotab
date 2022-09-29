// Copyright ©2022 Foolin. All rights reserved.

package gotab

import (
	"os"
	"testing"
)

const summaryEn = "Gotab is a go library for align text"
const summaryCn = "Gotab是控制台对齐的go库"

func TestPrintTable(t *testing.T) {
	writer := NewWriter()
	writer.Config.Split = "|"
	writer.AddLine("Hello, World", summaryEn)
	writer.AddLine("Hello，世A界B", summaryEn)
	writer.AddLine("您好,NI世界ABC", summaryCn)
	writer.AddDivider()
	data := [][]any{
		{"node2.example.com", "Ready", "1.12", "1.12"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
	}
	writer.AddLines(data...)
	writer.AddDividerWithValue("-----------------")
	writer.AddLine("Ready", "ABC", summaryEn)
	writer.AddLine("NotReady", summaryCn, "EDF")
	writer.Print()
}

func TestPrintTable2(t *testing.T) {
	writer := NewWriter()
	writer.Config.Split = "|"
	writer.AddDivider()
	data := [][]any{
		{"node2.example.com", "Ready", "1.12", "1.12"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
	}
	writer.AddLines(data...)
	writer.AddDividerWithValue("-----------------")
	writer.AddLine("NotReady2", "ABC", summaryEn)
	writer.AddLine("NotReady2", summaryCn, "EDF")
	writer.Print()
}

func TestPrintTable3(t *testing.T) {
	writer := NewWriter()
	writer.Config.Split = "|"
	writer.AddLine("A", summaryEn)
	writer.AddLine("Hello, 世界", summaryEn)
	writer.AddLine("您好, 世界", summaryEn)
	//writer.AddDivider()
	data := [][]any{
		{"node2.example.com", "Ready", "1.12", "1.12"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
	}
	writer.AddLines(data...)
	writer.AddDividerWithValue("-----------------")
	writer.AddLine("您好，世界", "ABC", summaryEn)
	writer.AddLine("Hello, 世界", "ABC", summaryEn)
	writer.AddLine("A", summaryEn, "Hello World")
	result := writer.String()

	println("|---------------------------|")
	println(result)
	println("|---------------------------|")
}

func TestPrintTable4(t *testing.T) {
	config := DefaultConfig
	config.Split = "|"
	config.MinWiths = []int{100}
	//config.MinWith = 20
	//config.MaxWith = 30
	writer := NewWriterWithConfig(config)
	writer.AddLine("A", summaryEn)
	writer.AddLine("Hello, 世界", summaryEn)
	writer.AddLine("您好, 世界", summaryEn)
	writer.AddDividerWithValue("-----------------")
	data := [][]any{
		{"node2.example.com", "Ready", "1.12", "1.12"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
	}
	writer.AddLines(data...)
	writer.AddDividerWithValue("-----------------")
	writer.AddLine("Hello, 世界", "ABC", summaryEn)
	writer.AddLine("A", summaryEn, "Hello World")
	result := writer.String()

	println("|---------------------------|")
	println(result)
	println("|---------------------------|")
}

func TestPrintTable5(t *testing.T) {
	data := [][]any{
		{"node1.example.com", summaryEn, "A"},
		{"----"},
		{"node2.example.com", "Ready", "1.12", "1.12"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
		{"node3.example.com", "Ready", "compute", "1.13"},
		{"node4.example.com", "NotReady", "compute", "1.11"},
	}
	Print(os.Stdout, data...)
}

func TestIsDivider(t *testing.T) {
	var str string = "================="
	println("Div: ", isDivider(str))
	println("Val: ", str)

	var div Divider = "================="
	println("Div: ", isDivider(div))
	println("Val: ", div)
}
