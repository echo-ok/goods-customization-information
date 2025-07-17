package gci

import (
	"strings"
)

// ValueType 值类型
type ValueType interface {
	~string | ~int | ~int64 | ~float64 | ~[]string | ~[]int | ~[]int64 | ~[]float64
}

// Text 定制文本
type Text struct {
	Label string `json:"label"` // 标签
	Value any    `json:"value"` // 值
}

func NewText[T ValueType](label string, value T) (Text, error) {
	return Text{
		Label: strings.TrimSpace(label),
		Value: value,
	}, nil
}
