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

func TestSetDarkMode(t *testing.T) {
	app.NewWithID(time.Now().String())
	obj := NewDefault()

	result := obj.SetDarkMode(true)
	assert.True(t, result, "The method should return true")
	assert.True(t, obj.isDarkMode, "The mode should be set to dark")

	result = obj.SetDarkMode(false)
	assert.True(t, result, "The method should return true")
	assert.False(t, obj.isDarkMode, "The mode should be set to light")
}

func TestSetPrimaryColor(t *testing.T) {
	app.NewWithID(time.Now().String())
	obj := NewDefault()

	existingColor := fyne.ThemeColorName(theme.ColorBlue)
	obj.paletteMap[existingColor] = &PaletteObj{}

	result := obj.SetPrimaryColor(existingColor)
	assert.True(t, result, "The method should return true")
	assert.Equal(t, existingColor, obj.primaryColor, "The color should be set correctly")

	nonExistingColor := fyne.ThemeColorName("non_existing_color")
	result = obj.SetPrimaryColor(nonExistingColor)
	assert.False(t, result, "The method should return false")
	assert.NotEqual(t, nonExistingColor, obj.primaryColor, "The color should not change")
}

/* ###################### */

func BenchmarkSetDarkMode(b *testing.B) {
	app.NewWithID(time.Now().String())
	obj := NewDefault()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj.SetDarkMode(i%2 == 0)
	}
}

func BenchmarkSetPrimaryColor(b *testing.B) {
	app.NewWithID(time.Now().String())
	obj := NewDefault()
	obj.paletteMap[theme.ColorBlue] = &PaletteObj{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj.SetPrimaryColor(theme.ColorBlue)
	}
}
