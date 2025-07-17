package gci

import "gopkg.in/guregu/null.v4"

// Surface 定制面
type Surface struct {
	Name         null.String `json:"name"`          // 名称
	PreviewImage *Image      `json:"preview_image"` // 预览图
	Regions      []Region    `json:"regions"`       // 定制区域
}

func NewSurface(name ...string) Surface {
	sf := Surface{
		PreviewImage: nil,
		Regions:      make([]Region, 0),
	}
	if len(name) != 0 {
		if name[0] != "" {
			sf.Name = null.StringFrom(name[0])
		}
	}
	return sf
}

func (sf *Surface) SetPreviewImage(image *Image) *Surface {
	sf.PreviewImage = image
	return sf
}
func (sf *Surface) AddRegion(region Region) *Surface {
	sf.Regions = append(sf.Regions, region)
	return sf
}
