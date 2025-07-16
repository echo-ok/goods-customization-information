package gci

import "gopkg.in/guregu/null.v4"

// Region 区域
type Region struct {
	Name   null.String `json:"name"`   // 区域名称
	Type   Type        `json:"type"`   // 类型
	Texts  []Text      `json:"texts"`  // 定制文本
	Images []Image     `json:"images"` // 定制图片
	Ok     bool        `json:"ok"`     // 是否可用
	Error  error       `json:"error"`  // 错误信息
}

func NewRegion(name ...string) Region {
	r := Region{}
	if len(name) != 0 {
		if name[0] != "" {
			r.Name = null.StringFrom(name[0])
		}
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
