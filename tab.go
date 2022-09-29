// Copyright ©2022 Foolin. All rights reserved.

package gotab

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type TabWriter struct {
	io.Writer
	Config     TabConfig
	groupData  map[int]*groupData
	groupIndex int
}

func NewWriter() *TabWriter {
	return &TabWriter{
		Writer:     os.Stdout,
		Config:     DefaultConfig,
		groupData:  make(map[int]*groupData, 1),
		groupIndex: 0,
	}
}

func NewWriterWithConfig(config TabConfig) *TabWriter {
	return &TabWriter{
		Writer:     os.Stdout,
		Config:     config,
		groupData:  make(map[int]*groupData, 1),
		groupIndex: 0,
	}
}

func (r *TabWriter) WithWriter(write io.Writer) *TabWriter {
	r.Writer = write
	return r
}

func (r *TabWriter) lastData() *groupData {
	data, ok := r.groupData[r.groupIndex]
	if !ok {
		data = &groupData{}
		r.groupData[r.groupIndex] = data
	}
	return data
}

func (r *TabWriter) AddLine(v ...any) *TabWriter {
	r.lastData().AddLine(v...)
	return r
}

func (r *TabWriter) AddLines(lines ...[]any) *TabWriter {
	r.lastData().AddLines(lines...)
	return r
}

func (r *TabWriter) AddEmptyLine() *TabWriter {
	return r.AddLine("")
}

func (r *TabWriter) AddDivider() *TabWriter {
	return r.AddDividerWithValue(DefaultDivider)
}

func (r *TabWriter) AddDividerWithValue(div Divider) *TabWriter {
	r.groupIndex++
	r.lastData().Div = div
	return r
}

func (r *TabWriter) Print() {
	writer := r.Writer
	config := r.Config
	groupIndex := r.groupIndex
	groupData := r.groupData

	for g := 0; g <= groupIndex; g++ {
		group := groupData[g]
		data := group.Data
		widths := make(map[int]int)
		//maxWidth := 0
		for i := range data {
			size := len(data[i])
			//total := 0
			for j := 0; j < (size - 1); j++ {
				value := fmt.Sprint(data[i][j])
				width := StrWidth(value)
				if width > widths[j] {
					widths[j] = width
				}
				//total += width
			}
			//if total > maxWidth {
			//	maxWidth = total
			//}
		}

		//Print
		if g > 0 {
			if len(group.Div) > 0 {
				fmt.Fprint(writer, group.Div, config.NewLine)
			}
		}

		for i := range data {
			line := data[i]
			size := len(line)
			if size == 0 {
				fmt.Fprint(writer, config.NewLine)
				continue
			}

			for n := range line {
				value := fmt.Sprint(line[n])
				if n == 0 && isDivider(value) {

				}
				//Is End Or Sample Value
				width, ok := widths[n]
				maxWidth := getArrayValue(config.MaxWiths, n, 0)
				minWidth := getArrayValue(config.MinWiths, n, 0)
				if minWidth > 0 && width < minWidth {
					width = minWidth

				}
				if maxWidth > 0 && width > maxWidth {
					value = StrTruncate(value, maxWidth, config.Ellipsis)
					width = maxWidth
				}
				padValue := Padding(value, config.Pad, width)
				if ok && n < (size-1) {
					fmt.Fprintf(writer, " %s %s", padValue, config.Split)
				} else {
					fmt.Fprintf(writer, " %s", padValue)
				}
			}

			fmt.Fprint(writer, config.NewLine)
		}
	}

}

func getArrayValue[T any](arr []T, idx int, defaultValue T) T {
	if idx < len(arr) {
		return arr[idx]
	}
	return defaultValue
}

func (r *TabWriter) PrintWriter(writer io.Writer) {
	r.WithWriter(writer).Print()
}

func (r *TabWriter) String() string {
	builder := &strings.Builder{}
	tempWriter := r.Writer
	r.WithWriter(builder).Print()
	r.Writer = tempWriter
	return builder.String()
}

func Print(writer io.Writer, lines ...[]any) {
	NewWriter().WithWriter(writer).AddLines(lines...).Print()
}

func isDivider(v any) bool {
	_, ok := v.(Divider)
	return ok
}
