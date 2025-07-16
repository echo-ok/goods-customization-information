package gci

// Text 定制文本
type Text struct {
	Label string `json:"label"` // 标签
	Value any    `json:"value"` // 值
}

func NewText(label string, value any) Text {
	return Text{label, value}
}
