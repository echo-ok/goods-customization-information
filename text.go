package gci

import (
	"errors"
	"strings"
)

// Text 定制文本
type Text struct {
	Label string `json:"label"` // 标签
	Value string `json:"value"` // 值
}

func NewText(label, value string) (Text, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Text{}, errors.New("gci: text value is empty")
	}
	return Text{strings.TrimSpace(label), value}, nil
}
