package gci

import (
	"gopkg.in/guregu/null.v4"
)

// Region 定制区域
type Region struct {
	Name   null.String `json:"name"`   // 区域名称
	Type   Type        `json:"type"`   // 类型
	Texts  []Text      `json:"texts"`  // 定制文本
	Images []Image     `json:"images"` // 定制图片
	Valid  bool        `json:"valid"`  // 是否有效
	Error  null.String `json:"error"`  // 错误信息
}

func NewRegion(name ...string) Region {
	r := Region{
		Type:   UnknownType,
		Valid:  false,
		Texts:  make([]Text, 0),
		Images: make([]Image, 0),
	}
	if len(name) != 0 && name[0] != "" {
		r.Name = null.StringFrom(name[0])
	}
	return r
}

func (r *Region) typecast() *Region {
	var tn, in = len(r.Texts), len(r.Images)
	if tn == 0 && in == 0 {
		r.Type = UnknownType
	} else if tn != 0 && in != 0 {
		r.Type = TextImageType
	} else if tn == 0 {
		r.Type = ImageType
	} else {
		r.Type = TextType
	}
	r.Valid = !r.Error.Valid
	return r
}

func (r *Region) AddText(text Text) *Region {
	r.Texts = append(r.Texts, text)
	r.typecast()
	return r
}

func (r *Region) AddImage(image Image) *Region {
	r.Images = append(r.Images, image)
	r.typecast()
	return r
}

func (r *Region) SetError(msg any) *Region {
	str := toString(msg)
	if str == "" {
		r.Error = null.NewString("", false)
	} else {
		r.Error = null.StringFrom(str)
	}
	r.Valid = !r.Error.Valid
	return r
}
