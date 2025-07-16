package gci

import "gopkg.in/guregu/null.v4"

// FInfo 文件信息
type FInfo struct {
	Url      null.String `json:"url"`       // 图片 URL
	Name     null.String `json:"name"`      // 名称（带扩展名）
	Title    null.String `json:"title"`     // 标题（不带扩展名）
	MimeType null.String `json:"mime_type"` // MIME 类型
	Ext      null.String `json:"ext"`       // 扩展名
	Size     null.Int    `json:"size"`      // 大小
	Bytes    []byte      `json:"bytes"`     // 字节
	Base64   null.String `json:"base64"`    // Base64 内容
	Valid    bool        `json:"valid"`     // 是否有效
	Error    error       `json:"error"`     // 错误信息
}

// SetError 设置错误
func (fi *FInfo) SetError(err error) *FInfo {
	fi.Error = err
	fi.Valid = false
	return fi
}

// SetValid 设置文件是有效的
func (fi *FInfo) SetValid() *FInfo {
	fi.Error = nil
	fi.Valid = true
	return fi
}
