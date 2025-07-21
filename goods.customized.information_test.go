package gci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetRawData(t *testing.T) {
	ci := NewGoodsCustomizedInformation()
	type data struct {
		Value  any
		String string
	}
	rawData := []data{
		{Value: `{"a": 1}`, String: `{"a": 1}`},
		{Value: 1, String: "1"},
		{Value: struct {
			Name string
			Age  int
		}{Name: "Jake", Age: 18}, String: `{"Name":"Jake","Age":18}`},
	}
	for _, datum := range rawData {
		ci.SetRawData(datum.Value)
		assert.Equal(t, datum.String, ci.RawData.ValueOrZero())
	}
}

func Test_Surface(t *testing.T) {
	ci := NewGoodsCustomizedInformation()
	ci.SetRawData(`[{"name": "Joke", "age": 12}]`)
	surface := NewSurface()
	assert.Equal(t, true, surface.PreviewImage == nil, "surface preview image default equal nil")
	img, err := NewImage("https://www.a.com/b.jpg", true)
	assert.Equal(t, nil, err)
	surface.SetPreviewImage(&img)
	ci.AddSurface(surface)
	assert.Equal(t, 1, len(ci.Surfaces))
	assert.Equal(t, "https://www.a.com/b.jpg", surface.PreviewImage.RawUrl)
	assert.Equal(t, true, surface.PreviewImage.Redownload)
	img.SetRedownload(false)
	assert.Equal(t, false, surface.PreviewImage.Redownload)
	assert.Equal(t, "https://www.a.com/b.jpg", surface.PreviewImage.Url.String)

	img.SetError("xxx")
	assert.Equal(t, "xxx", surface.PreviewImage.Error.ValueOrZero())
	assert.Equal(t, false, surface.PreviewImage.Valid)

	img.SetError(nil)
	assert.Equal(t, "", surface.PreviewImage.Error.ValueOrZero())
	assert.Equal(t, true, surface.PreviewImage.Valid)

	region := NewRegion("a")
	assert.Equal(t, "a", region.Name.ValueOrZero())

	text, err := NewText("", "aaa")
	assert.Equal(t, true, err == nil)

	text, err = NewText("", " bbb ")
	assert.Equal(t, nil, err)
	region.AddText(text)
	assert.Equal(t, 1, len(region.Texts))

	// 空值处理
	text, err = NewText("", "     ")
	assert.Equal(t, true, err != nil)
	if err == nil {
		region.AddText(text)
		assert.Equal(t, "", text.Value)
		assert.Equal(t, 2, len(region.Texts))
	}

	text, err = NewText("", []int{1, 2, 3})
	assert.Equal(t, nil, err)
	region.AddText(text)
	assert.Equal(t, 2, len(region.Texts))

	surface.AddRegion(region)
	assert.Equal(t, 1, len(surface.Regions))
	ci.AddSurface(surface)
	assert.Equal(t, 2, len(ci.Surfaces))
	assert.Equal(t, 2, len(ci.Surfaces[1].Regions[0].Texts))
	assert.Equal(t, "bbb", ci.Surfaces[1].Regions[0].Texts[0].Value)
	assert.Equal(t, []int{1, 2, 3}, ci.Surfaces[1].Regions[0].Texts[1].Value)

	t.Logf("Goods Customized Information JSON = %s", ci.String())
}
