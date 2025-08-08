package gci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
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
	surface.SetPreviewImage(img)
	ci.AddSurface(surface)
	assert.Equal(t, 1, len(ci.Surfaces))
	assert.Equal(t, "https://www.a.com/b.jpg", surface.PreviewImage.RawUrl)
	assert.Equal(t, true, surface.PreviewImage.Redownload)
	surface.PreviewImage.SetRedownload(false)
	assert.Equal(t, false, surface.PreviewImage.Redownload)
	assert.Equal(t, "https://www.a.com/b.jpg", surface.PreviewImage.Url.String)

	surface.PreviewImage.SetError("xxx")
	assert.Equal(t, "xxx", surface.PreviewImage.Error.ValueOrZero())
	assert.Equal(t, false, surface.PreviewImage.Valid)

	surface.PreviewImage.SetError(nil)
	assert.Equal(t, "", surface.PreviewImage.Error.ValueOrZero())
	assert.Equal(t, true, surface.PreviewImage.Valid)

	region := NewRegion("a")
	assert.Equal(t, "a", region.Name.ValueOrZero())
	assert.Equal(t, UnknownType, region.Type)

	text, err := NewText("", "aaa")
	assert.Equal(t, true, err == nil)

	text, err = NewText("", " bbb ")
	assert.Equal(t, nil, err)
	region.AddText(text)
	assert.Equal(t, 1, len(region.Texts))
	assert.Equal(t, TextType, region.Type)

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

func TestGoodsCustomizedInformation_From(t *testing.T) {
	type fields struct {
		RawData  null.String
		Surfaces []Surface
	}
	type args struct {
		previewImage string
		texts        []string
		images       []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"t1 未提供 texts, images 参数值",
			fields{
				RawData: null.StringFrom(`{"preview_image":"",texts:[],images:[]}`),
			},
			args{
				"",
				[]string{},
				[]string{},
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Equal(t, "gci: Either texts or images must be filled in, but not both can be empty", err.Error())
				return true
			},
		},
		{
			"t2 文本没有分隔符，无法区分 label, value",
			fields{
				RawData: null.StringFrom(`{"preview_image":"",texts:["a"],images:[]}`),
			},
			args{
				"",
				[]string{"a"},
				[]string{},
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Equal(t, "gci: invalid text: a", err.Error())
				return true
			},
		},
		{
			"t3 文本没有 label",
			fields{
				RawData: null.StringFrom(`{"preview_image":"",texts:[":b"],images:[]}`),
			},
			args{
				"",
				[]string{":b"},
				[]string{},
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Equal(t, "gci: invalid label", err.Error())
				return true
			},
		},
		{
			"t4 文本有 label，无 value",
			fields{
				RawData: null.StringFrom(`{"preview_image":"",texts:["a:"],images:[]}`),
			},
			args{
				"",
				[]string{"a:"},
				[]string{},
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Equal(t, nil, err)
				g, _ := i[1].(*GoodsCustomizedInformation)
				assert.Equal(t, 1, len(g.Surfaces))
				exceptedText, _ := NewText("a", "")
				actualText := g.Surfaces[0].Regions[0].Texts[0]
				assert.Equal(t, exceptedText, actualText)
				assert.Equal(t, "a", actualText.Label)
				assert.Equal(t, "", actualText.Value)
				return true
			},
		},
		{
			"t5 文本有 label，有 value",
			fields{
				RawData: null.StringFrom(`{"preview_image":"",texts:["a:123"],images:[]}`),
			},
			args{
				"",
				[]string{"a:123"},
				[]string{},
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Equal(t, nil, err)
				g, _ := i[1].(*GoodsCustomizedInformation)
				assert.Equal(t, 1, len(g.Surfaces))
				exceptedText, _ := NewText("a", "123")
				actualText := g.Surfaces[0].Regions[0].Texts[0]
				assert.Equal(t, exceptedText, actualText)
				assert.Equal(t, "a", actualText.Label)
				assert.Equal(t, "123", actualText.Value)
				return true
			},
		},
		{
			"t6 无文本，有图片",
			fields{
				RawData: null.StringFrom(`{"preview_image":"",texts:[],images:["https://www.a.com/b.jpg"]}`),
			},
			args{
				"",
				[]string{},
				[]string{"https://www.a.com/b.jpg"},
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Equal(t, nil, err)
				g, _ := i[1].(*GoodsCustomizedInformation)
				assert.Equal(t, 1, len(g.Surfaces))
				assert.Equal(t, 0, len(g.Surfaces[0].Regions[0].Texts))
				images := g.Surfaces[0].Regions[0].Images
				assert.Equal(t, 1, len(images))
				exceptedImage, _ := NewImage("https://www.a.com/b.jpg", false)
				assert.Equal(t, exceptedImage, images[0])
				assert.Equal(t, "https://www.a.com/b.jpg", exceptedImage.Url.ValueOrZero())
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gci := &GoodsCustomizedInformation{
				RawData:  tt.fields.RawData,
				Surfaces: tt.fields.Surfaces,
			}
			err := gci.Build(Material{
				Name:         "",
				PreviewImage: tt.args.previewImage,
				Texts:        tt.args.texts,
				Images:       tt.args.images,
			})
			tt.wantErr(t, err, fmt.Sprintf("Build(%v, %v, %v)", tt.args.previewImage, tt.args.texts, tt.args.images), gci)
		})
	}
}
