package gci

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
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
	Surfaces []Surface   `json:"surfaces"` // 定制面
}

func NewGoodsCustomizedInformation() GoodsCustomizedInformation {
	return GoodsCustomizedInformation{
		RawData:  null.NewString("", false),
		Surfaces: make([]Surface, 0),
	}
}

func (gci *GoodsCustomizedInformation) Reset() *GoodsCustomizedInformation {
	gci.RawData = null.NewString("", false)
	gci.Surfaces = make([]Surface, 0)
	return gci
}

// Build 构建定制信息
func (gci *GoodsCustomizedInformation) Build(previewImage string, texts []string, images []string) error {
	if len(texts) == 0 && len(images) == 0 {
		return errors.New("gci: Either texts or images must be filled in, but not both can be empty")
	}
	gci.Reset()
	tb, _ := json.Marshal(texts)
	ib, _ := json.Marshal(images)
	gci.RawData = null.NewString(fmt.Sprintf(`{"preview_image": "%s", texts": "%s", "images": "%s"}`, previewImage, string(tb), string(ib)), true)
	region := NewRegion()
	for _, lineStr := range texts {
		lineStr = strings.TrimSpace(lineStr)
		label, value, ok := strings.Cut(lineStr, ":")
		if !ok {
			return fmt.Errorf("gci: invalid text: %s", lineStr)
		}

		label = strings.TrimSpace(label)
		value = strings.TrimSpace(value)
		if label == "" {
			return errors.New("gci: invalid label")
		}

		text, err := NewText(label, value)
		if err != nil {
			return err
		}

		region.AddText(text)
	}
	for _, img := range images {
		image, err := NewImage(img, false)
		if err != nil {
			return err
		}
		region.AddImage(image)
	}
	surface := NewSurface()
	if previewImage != "" {
		image, err := NewImage(previewImage, false)
		if err != nil {
			return err
		}
		surface.SetPreviewImage(image)
	}
	surface.AddRegion(region)
	gci.AddSurface(surface)
	return nil
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
