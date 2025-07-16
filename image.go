package gci

import (
	"errors"
	"strings"

	"github.com/hiscaler/filer-go"
	"gopkg.in/guregu/null.v4"
)

// Image 定制图片
type Image struct {
	redownload bool        // 是否重新下载
	Label      null.String `json:"label"`   // 标签
	RawUrl     string      `json:"raw_url"` // 图片原始地址
	Url        null.String `json:"url"`     // 图片地址
	Valid      bool        `json:"valid"`   // 是否有效
	Error      null.String `json:"error"`   // 错误信息
}

func NewImage(url string, redownload bool) (Image, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return Image{}, errors.New("gci: url is empty")
	}
	img := Image{redownload: false, RawUrl: url}
	img.SetRedownload(redownload)
	return img, nil
}

func (img *Image) SetRedownload(b bool) *Image {
	if img == nil {
		return img
	}
	img.redownload = b
	if !b {
		img.Url = null.StringFrom(img.RawUrl)
		img.Valid = true
	}
	return img
}

func (img *Image) Redownload() bool {
	return img.redownload
}

func (img *Image) SetUrl(url string) *Image {
	img.Url = null.StringFrom(url)
	img.SetError(nil)
	return img
}

func (img *Image) SetError(msg any) *Image {
	str := toString(msg)
	if str == "" {
		img.Error = null.NewString("", false)
	} else {
		img.Error = null.StringFrom(str)
	}
	img.Valid = !img.Error.Valid
	return img
}

// SaveTo Save image to local
func (img *Image) SaveTo(filename string) (string, error) {
	if img == nil {
		return "", errors.New("gci: image is nil")
	}
	if !img.Url.Valid {
		return "", errors.New(img.Error.ValueOrZero())
	}

	if !img.Url.Valid {
		return "", errors.New("gci: image.url is empty")
	}

	fer := filer.NewFiler()
	err := fer.Open(img.Url.String)
	if err != nil {
		return "", err
	}
	return fer.SaveTo(filename)
}
