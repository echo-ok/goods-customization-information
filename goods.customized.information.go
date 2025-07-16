package gci

import (
	"errors"

	"github.com/hiscaler/filer-go"
	"gopkg.in/guregu/null.v4"
)

type Type string

const (
	UnknownType   Type = "unknown"   // 未知
	TextType      Type = "text"      // 仅文字
	ImageType     Type = "image"     // 仅图片
	TextImageType Type = "textImage" // 文字和图片
)

func (t Type) IsValid() bool {
	return t == TextType || t == ImageType || t == TextImageType || t == UnknownType
}

type Text struct {
	Region null.String `json:"region"` // 区域
	Label  string      `json:"label"`  // 标签
	Value  any         `json:"value"`  // 值
}

type Image struct {
	Region null.String `json:"region"`  // 区域
	RawUrl string      `json:"raw_url"` // 图片原始地址
	FInfo  FInfo       `json:"finfo"`   // 文件信息
}

// SaveTo 图片保存
func (img Image) SaveTo(filename string) (string, error) {
	if !img.FInfo.Url.Valid {
		return "", img.FInfo.Error
	}

	if !img.FInfo.Url.Valid && len(img.FInfo.Bytes) == 0 && !img.FInfo.Base64.Valid {
		return "", errors.New("gci: image.url|bytes|base64 value is empty")
	}

	var file any
	fer := filer.NewFiler()
	if len(img.FInfo.Bytes) != 0 {
		file = img.FInfo.Bytes
	} else if img.FInfo.Base64.Valid {
		file = img.FInfo.Base64.String
	} else if img.FInfo.Url.Valid {
		file = img.FInfo.Url.String
	}
	err := fer.Open(file)
	if err != nil {
		return "", err
	}
	return fer.SaveTo(filename)
}

// Surface 面
type Surface struct {
	Name         null.String `json:"name"`          // 名称
	PreviewImage null.String `json:"preview_image"` // 预览图
	Regions      []Region    `json:"regions"`       // 区域内容
}

func (sf *Surface) AddRegion(region Region) *Surface {
	sf.Regions = append(sf.Regions, region)
	return sf
}

// Region 区域
type Region struct {
	Name         null.String `json:"name"`          // 区域名称
	Type         Type        `json:"type"`          // 类型
	PreviewImage null.String `json:"preview_image"` // 预览图
	Texts        []Text      `json:"texts"`         // 定制文本
	Images       []Image     `json:"images"`        // 定制图片
	Ok           bool        `json:"ok"`            // 是否可用
	Error        error       `json:"error"`         // 错误信息
}

func (sf *Region) typecast() *Region {
	var tn, in = len(sf.Texts), len(sf.Images)
	if tn == 0 && in == 0 {
		sf.Type = UnknownType
	} else if tn != 0 && in != 0 {
		sf.Type = TextImageType
	} else if tn == 0 {
		sf.Type = ImageType
	} else {
		sf.Type = TextType
	}
	return sf
}

func (sf *Region) AddText(text Text) *Region {
	sf.Texts = append(sf.Texts, text)
	sf.typecast()
	return sf
}

func (sf *Region) AddImage(image Image) *Region {
	sf.Images = append(sf.Images, image)
	sf.typecast()
	return sf
}

type GoodsCustomizedInformation struct {
	RawData  null.String `json:"raw_data"` // 原始数据
	Surfaces []Surface   `json:"surfaces"` // 面
}

func NewGoodsCustomizedInformation() *GoodsCustomizedInformation {
	return &GoodsCustomizedInformation{
		Surfaces: make([]Surface, 0),
	}
}

func (gci *GoodsCustomizedInformation) AddRawData(data string) *GoodsCustomizedInformation {
	if data == "" {
		gci.RawData = null.String{}
	} else {
		gci.RawData = null.StringFrom(data)
	}
	return gci
}

func (gci *GoodsCustomizedInformation) AddSurface(surface Surface) *GoodsCustomizedInformation {
	gci.Surfaces = append(gci.Surfaces, surface)
	return gci
}
