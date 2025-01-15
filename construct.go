package fyneThemeBuilder

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"sync"
)

// // // // // // // // // // // // // // // // // //

type ThemeInterfaceObj struct {
	parent     fyne.Theme
	paletteMap map[fyne.ThemeColorName]*PaletteObj

	propertyLock sync.RWMutex
	isDarkMode   bool
	primaryColor fyne.ThemeColorName

	HookFont func(fyne.TextStyle) fyne.Resource     // function to be installed by the user to create his own Font processing script
	HookIcon func(fyne.ThemeIconName) fyne.Resource // function to be installed by the user to create his own Icon processing script
	HookSize func(fyne.ThemeSizeName) float32       // function to be installed by the user to create his own Size processing script

}

func construct() *ThemeInterfaceObj {
	obj := new(ThemeInterfaceObj)

	obj.parent = fyne.CurrentApp().Settings().Theme()
	obj.isDarkMode = fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark

	obj.HookFont = func(t fyne.TextStyle) fyne.Resource { return obj.parent.Font(t) }
	obj.HookIcon = func(t fyne.ThemeIconName) fyne.Resource { return obj.parent.Icon(t) }
	obj.HookSize = func(t fyne.ThemeSizeName) float32 { return obj.parent.Size(t) }

	return obj
}

/* ###################### */

// New initialize the shared palette
//
// can be used to set a saved custom theme
func New(palette map[fyne.ThemeColorName]*PaletteObj) *ThemeInterfaceObj {
	obj := construct()
	obj.paletteMap = palette
	for _, p := range palette {
		obj.primaryColor = p.PrimaryColor.Name
	}
	return obj
}

// NewOne initializes the palette with one base color
func NewOne(primaryColor *PalettePrimaryColorObj, elements *PaletteElementsObj) *ThemeInterfaceObj {
	obj := construct()
	obj.paletteMap = make(map[fyne.ThemeColorName]*PaletteObj)

	obj.paletteMap[primaryColor.Name] = &PaletteObj{
		PrimaryColor: primaryColor,
		Elements:     elements,
	}
	obj.primaryColor = primaryColor.Name

	return obj
}

// NewDefault initializes the default palette
//
// values taken from DefaultPalette
func NewDefault() *ThemeInterfaceObj {
	return New(DefaultPalette)
}
