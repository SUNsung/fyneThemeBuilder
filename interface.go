package fyneThemeBuilder

import (
	"fyne.io/fyne/v2"
	"image/color"
)

// // // // // // // // // // // // // // // // // //

func (e *ThemeInterfaceObj) Font(s fyne.TextStyle) fyne.Resource {
	return e.HookFont(s)
}

func (e *ThemeInterfaceObj) Icon(n fyne.ThemeIconName) fyne.Resource {
	return e.HookIcon(n)
}

func (e *ThemeInterfaceObj) Size(n fyne.ThemeSizeName) float32 {
	return e.HookSize(n)
}

func (e *ThemeInterfaceObj) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return e.color(n)
}
