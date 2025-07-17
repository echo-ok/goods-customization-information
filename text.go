package gci

import (
	"strings"
)

// Text 定制文本
type Text struct {
	Label string `json:"label"` // 标签
	Value string `json:"value"` // 值
}

func NewText(label, value string) (Text, error) {
	return Text{
		Label: strings.TrimSpace(label),
		Value: strings.TrimSpace(value),
	}, nil
}
