// Copyright ©2022 Foolin. All rights reserved.

package gotab

type TabConfig struct {
	Pad      string
	Split    string
	NewLine  string
	Ellipsis string
	MinWiths []int
	MaxWiths []int
}

var DefaultConfig = TabConfig{
	Pad:      " ",
	Split:    "",
	NewLine:  "\n",
	Ellipsis: "...",
	MinWiths: []int{},
	MaxWiths: []int{},
}
