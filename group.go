// Copyright ©2022 Foolin. All rights reserved.

package gotab

type Divider string

const DefaultDivider = ""

type groupData struct {
	Data [][]any
	Div  Divider
}

func (g *groupData) AddLine(v ...any) {
	g.Data = append(g.Data, v)
}

func (g *groupData) AddLines(lines ...[]any) {
	for _, line := range lines {
		g.Data = append(g.Data, line)
	}
}
