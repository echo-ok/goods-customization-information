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
	surface := NewSurface()
	assert.Equal(t, true, surface.PreviewImage == nil, "surface preview image default equal nil")
	img := NewImage("https://www.a.com/b.jpg", true)
	img.OriginalUrl = "https://www.a.com/b.jpg"
	surface.PreviewImage = &img
	ci.AddSurface(surface)
	assert.Equal(t, 1, len(ci.Surfaces))
	assert.Equal(t, "https://www.a.com/b.jpg", surface.PreviewImage.OriginalUrl)
	assert.Equal(t, true, surface.PreviewImage.Redownload())
}
