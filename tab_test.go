// Copyright ©2022 Foolin. All rights reserved.

package gotab

import (
	"os"
	"testing"
)

const summaryEn = "Gotab is a golang library for console/writer output formatting and alignment text."
const summaryCn = "Gotab是用于控制台/编写器输出格式和对齐文本的golang库"

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
	v1 := "Cobra supports local flags which will only run when this command"
	writer.AddLine("A", v1)
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
	writer.AddLine("NotReady2", "ABC", v1)
	writer.AddLine("NotReady2", v1, "EDF")
	writer.Print()
}

func TestPrintTable3(t *testing.T) {
	writer := NewWriter()
	writer.Config.Split = "|"
	writer.AddLine("A", summaryEn)
	writer.AddLine("Hello, 世界", summaryEn)
	writer.AddLine("您好, 世界", summaryEn)
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
