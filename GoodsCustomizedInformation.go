package gci

import (
	"errors"

	"github.com/hiscaler/filer-go"
	"gopkg.in/guregu/null.v4"
)

const (
	UnknownType = "unknown"   // 未知
	TextType    = "text"      // 仅文字
	ImageType   = "image"     // 仅图片
	TextImage   = "textImage" // 文字和图片
)

type Text struct {
	Region null.String `json:"region"` // 区域
	Label  string      `json:"label"`  // 标签
	Value  any         `json:"value"`  // 值
}

type Image struct {
	RawUrl string      `json:"raw_url"`  // 图片原始地址
	Name   null.String `json:"name"`     // 图片名称
	Ext    null.String `json:"ext"`      // 图片扩展名
	Title  null.String `json:"title"`    // 图片标题
	Url    null.String `json:"download"` // 可访问地址
	Bytes  []byte      `json:"bytes"`    // 图片字节
	Base64 null.String `json:"base64"`   // 图片 Base64
	Ok     bool        `json:"ok"`       // 是否可用
	Error  error       `json:"error"`    // 错误信息
}

// SaveTo 图片保存
func (img Image) SaveTo(filename string) (string, error) {
	if !img.Url.Valid {
		return "", img.Error
	}

	if !img.Url.Valid && len(img.Bytes) == 0 && !img.Base64.Valid {
		return "", errors.New("gci: image.url|bytes|base64 value is empty")
	}

	var file any
	fer := filer.NewFiler()
	if len(img.Bytes) != 0 {
		file = img.Bytes
	} else if img.Base64.Valid {
		file = img.Base64.String
	} else if img.Url.Valid {
		file = img.Url.String
	}
	err := fer.Open(file)
	if err != nil {
		return "", err
	}
	return fer.SaveTo(filename)
}

type Surface struct {
	ID           null.String `json:"id"`            // ID
	Type         string      `json:"type"`          // 类型
	PreviewImage null.String `json:"preview_image"` // 预览图
	Texts        []Text      `json:"texts"`         // 定制文本
	Images       []Image     `json:"images"`        // 定制图片
	Ok           bool        `json:"ok"`            // 是否可用
	Error        error       `json:"error"`         // 错误信息
}

func (sf *Surface) typecast() *Surface {
	var tn, in = len(sf.Texts), len(sf.Images)
	if tn == 0 && in == 0 {
		sf.Type = UnknownType
	} else if tn != 0 && in != 0 {
		sf.Type = TextImage
	} else if tn == 0 {
		sf.Type = ImageType
	} else {
		sf.Type = TextType
	}
	return sf
}

func (sf *Surface) AddText(text Text) *Surface {
	sf.Texts = append(sf.Texts, text)
	sf.typecast()
	return sf
}

func (sf *Surface) AddImage(image Image) *Surface {
	sf.Images = append(sf.Images, image)
	sf.typecast()
	return sf
}

type GoodsCustomizedInformation struct {
	RawData  null.String `json:"raw_data"` // 原始数据
	Surfaces []Surface   `json:"surfaces"` // 面
}
