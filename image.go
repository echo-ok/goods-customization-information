package gci

import (
	"errors"
	"strings"

	"gopkg.in/guregu/null.v4"
)

// Image 定制图片
type Image struct {
	Label      null.String `json:"label"`      // 标签
	RawUrl     string      `json:"raw_url"`    // 图片原始地址
	Url        null.String `json:"url"`        // 图片地址
	Redownload bool        `json:"redownload"` // 是否需要下载
	Valid      bool        `json:"valid"`      // 是否有效
	Error      null.String `json:"error"`      // 错误信息
}

func NewImage(url string, redownload bool) (Image, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return Image{}, errors.New("gci: url is empty")
	}
	img := Image{Redownload: false, RawUrl: url}
	img.SetRedownload(redownload)
	return img, nil
}

func (img *Image) SetRedownload(b bool) *Image {
	if img == nil {
		return img
	}
	img.Redownload = b
	if !b {
		img.Url = null.StringFrom(img.RawUrl)
		img.Valid = true
	}
	return img
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
