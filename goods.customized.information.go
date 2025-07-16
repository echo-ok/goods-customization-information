package gci

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

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

type GoodsCustomizedInformation struct {
	RawData  null.String `json:"raw_data"` // 原始数据
	Surfaces []Surface   `json:"surfaces"` // 面
}

func NewGoodsCustomizedInformation() GoodsCustomizedInformation {
	return GoodsCustomizedInformation{
		Surfaces: make([]Surface, 0),
	}
}

func toString(value any) string {
	switch val := value.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	case nil:
		return ""
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr, reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		if b, err := json.Marshal(v.Interface()); err == nil {
			return string(b)
		} else {
			return ""
		}
	default:
		return fmt.Sprintf("%v", value)
	}
}
func (gci *GoodsCustomizedInformation) SetRawData(data any) *GoodsCustomizedInformation {
	str := toString(data)
	if str == "" {
		gci.RawData = null.String{}
	} else {
		gci.RawData = null.StringFrom(str)
	}
	return gci
}

func (gci *GoodsCustomizedInformation) AddSurface(surface Surface) *GoodsCustomizedInformation {
	gci.Surfaces = append(gci.Surfaces, surface)
	return gci
}

func (gci *GoodsCustomizedInformation) String() string {
	b, err := json.Marshal(gci)
	if err != nil {
		return ""
	}
	return string(b)
}
