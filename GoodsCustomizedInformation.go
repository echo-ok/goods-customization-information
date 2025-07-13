package gci

import "gopkg.in/guregu/null.v4"

const (
	UnknownType = "unknown"   // 未知
	TextType    = "text"      // 仅文字
	ImageType   = "image"     // 仅图片
	TextImage   = "textImage" // 文字和图片
)

type Text struct {
	Region null.String `json:"region"` // 区域
	Label  string      `json:"label"`  // 标签
	Value  string      `json:"value"`  // 值
}

type Image struct {
	RawUrl string      `json:"raw_url"`  // 图片原始地址
	Name   null.String `json:"name"`     // 图片名称
	Url    null.String `json:"download"` // 可访问地址
	Error  error       `json:"error"`    // 错误信息
	Ok     bool        `json:"ok"`       // 是否可下载
}

// Download 图片下载
func (img Image) Download(filename string) error {
	if !img.Url.Valid {
		return img.Error
	}
	return nil
}

type Surface struct {
	ID           null.String `json:"id"`            // ID
	Type         string      `json:"type"`          // 类型
	PreviewImage null.String `json:"preview_image"` // 预览图
	Images       []Image     `json:"images"`        // 定制图片
	Texts        []Text      `json:"texts"`         // 定制内容
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
