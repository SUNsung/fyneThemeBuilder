package fyneThemeBuilder

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// // // // // // // // // // // // // // // // // //

func TestNew(t *testing.T) {
	app.NewWithID(time.Now().String())
	palette := map[fyne.ThemeColorName]*PaletteObj{
		"color1": {
			PrimaryColor: &PalettePrimaryColorObj{Name: "color1"},
			Elements:     &PaletteElementsObj{},
		},
	}

	obj := New(palette)
	assert.NotNil(t, obj, "The object must not be nil")
	assert.Equal(t, palette, obj.paletteMap, "The palette must match the one provided")
}

func TestNewOne(t *testing.T) {
	app.NewWithID(time.Now().String())
	primaryColor := &PalettePrimaryColorObj{Name: "primary1"}
	elements := &PaletteElementsObj{}

	obj := NewOne(primaryColor, elements)
	assert.NotNil(t, obj, "The object must not be nil")
	assert.Equal(t, primaryColor.Name, obj.primaryColor, "The primary color must match")
	assert.NotNil(t, obj.paletteMap[primaryColor.Name], "The palette for the primary color must be created")
}

func TestNewDefault(t *testing.T) {
	app.NewWithID(time.Now().String())
	obj := NewDefault()

	assert.NotNil(t, obj, "The object must not be nil")
	assert.NotEmpty(t, obj.paletteMap, "The palette must be populated")
}

func TestHooks(t *testing.T) {
	app.NewWithID(time.Now().String())
	obj := NewDefault()

	font := obj.HookFont(fyne.TextStyle{})
	assert.NotNil(t, font, "The font must not be nil")

	icon := obj.HookIcon(theme.IconNameCancel)
	assert.NotNil(t, icon, "The icon must not be nil")

	size := obj.HookSize(theme.SizeNameInputBorder)
	assert.True(t, size > 0, "The size must be a positive number")
}

/* ###################### */

func BenchmarkNew(b *testing.B) {
	app.NewWithID(time.Now().String())
	palette := map[fyne.ThemeColorName]*PaletteObj{
		"color1": {
			PrimaryColor: &PalettePrimaryColorObj{Name: "color1"},
			Elements:     &PaletteElementsObj{},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		New(palette)
	}
}

func BenchmarkNewOne(b *testing.B) {
	app.NewWithID(time.Now().String())
	primaryColor := &PalettePrimaryColorObj{Name: "primary1"}
	elements := &PaletteElementsObj{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewOne(primaryColor, elements)
	}
}

func BenchmarkNewDefault(b *testing.B) {
	app.NewWithID(time.Now().String())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewDefault()
	}
}
