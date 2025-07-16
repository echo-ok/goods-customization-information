package gci

import "gopkg.in/guregu/null.v4"

type Text struct {
	Label string `json:"label"` // 标签
	Value any    `json:"value"` // 值
}

func NewText(label string, value any) Text {
	return Text{label, value}
}

type RegionText struct {
	Region null.String `json:"region"` // 区域
	Text
}
