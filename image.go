package gci

import (
	"errors"

	"github.com/hiscaler/filer-go"
	"gopkg.in/guregu/null.v4"
)

// Image 定制图片
type Image struct {
	redownload  bool   // 是否重新下载
	OriginalUrl string `json:"original_url"` // 图片原始地址
	FInfo       FInfo  `json:"finfo"`        // 文件信息
}

func NewImage(url string, redownload bool) Image {
	img := Image{redownload: false, OriginalUrl: url, FInfo: FInfo{}}
	img.SetRedownload(redownload)
	return img
}

func (img *Image) SetRedownload(b bool) *Image {
	if img == nil {
		return img
	}
	img.redownload = b
	if b {
		img.FInfo.Url = null.StringFrom(img.OriginalUrl)
		img.FInfo.Valid = true
	}
	return img
}

func (img *Image) Redownload() bool {
	return img.redownload
}

// SaveTo Save image to local
func (img *Image) SaveTo(filename string) (string, error) {
	if img == nil {
		return "", errors.New("gci: image is nil")
	}
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
