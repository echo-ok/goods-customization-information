package gci

import (
	"errors"
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
	label = strings.TrimSpace(label)
	emptyLabel := label == ""
	emptyValue := false
	if str, ok := any(value).(string); ok {
		str = strings.TrimSpace(str)
		emptyValue = str == ""
		value = any(str).(T)
	}
	if emptyLabel && emptyValue {
		return Text{}, errors.New("gci: label and value are empty")
	}
	return Text{
		Label: label,
		Value: value,
	}, nil
}
