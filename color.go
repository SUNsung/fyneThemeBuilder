package fyneThemeBuilder

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// // // // // // // // // // // // // // // // // //

func (e *ThemeInterfaceObj) color(n fyne.ThemeColorName) color.Color {
	e.propertyLock.RLock()
	defer e.propertyLock.RUnlock()

	if e.paletteMap == nil || len(e.paletteMap) == 0 {
		if e.isDarkMode {
			return e.parent.Color(n, theme.VariantDark)
		} else {
			return e.parent.Color(n, theme.VariantLight)
		}
	}

	palette, ok := e.paletteMap[e.primaryColor]
	if !ok {
		fmt.Println(e.primaryColor)
		if e.isDarkMode {
			return e.parent.Color(n, theme.VariantDark)
		} else {
			return e.parent.Color(n, theme.VariantLight)
		}
	}

	switch n {
	case theme.ColorNamePrimary, theme.ColorNameHyperlink:
		return palette.PrimaryColor.Color.Color(e.isDarkMode)

	case theme.ColorNameForeground:
		return palette.Elements.Foreground.Global.Color(e.isDarkMode)
	case theme.ColorNameForegroundOnPrimary:
		return palette.Elements.Foreground.Primary.Color(e.isDarkMode)
	case theme.ColorNameForegroundOnError:
		return palette.Elements.Foreground.Error.Color(e.isDarkMode)
	case theme.ColorNameForegroundOnSuccess:
		return palette.Elements.Foreground.Success.Color(e.isDarkMode)
	case theme.ColorNameForegroundOnWarning:
		return palette.Elements.Foreground.Warning.Color(e.isDarkMode)

	case theme.ColorNameFocus:
		return palette.Elements.Focus.Color(e.isDarkMode)
	case theme.ColorNameSelection:
		return palette.Elements.Selection.Color(e.isDarkMode)
	case theme.ColorNameBackground:
		return palette.Elements.Background.Color(e.isDarkMode)
	case theme.ColorNameButton:
		return palette.Elements.Button.Color(e.isDarkMode)
	case theme.ColorNameDisabled:
		return palette.Elements.Disabled.Color(e.isDarkMode)
	case theme.ColorNameDisabledButton:
		return palette.Elements.DisabledButton.Color(e.isDarkMode)
	case theme.ColorNameError:
		return palette.Elements.Error.Color(e.isDarkMode)
	case theme.ColorNameHover:
		return palette.Elements.Hover.Color(e.isDarkMode)
	case theme.ColorNameHeaderBackground:
		return palette.Elements.HeaderBackground.Color(e.isDarkMode)
	case theme.ColorNameInputBackground:
		return palette.Elements.InputBackground.Color(e.isDarkMode)
	case theme.ColorNameInputBorder:
		return palette.Elements.InputBorder.Color(e.isDarkMode)
	case theme.ColorNameMenuBackground:
		return palette.Elements.MenuBackground.Color(e.isDarkMode)
	case theme.ColorNameOverlayBackground:
		return palette.Elements.OverlayBackground.Color(e.isDarkMode)
	case theme.ColorNamePlaceHolder:
		return palette.Elements.Placeholder.Color(e.isDarkMode)
	case theme.ColorNamePressed:
		return palette.Elements.Pressed.Color(e.isDarkMode)
	case theme.ColorNameScrollBar:
		return palette.Elements.ScrollBar.Color(e.isDarkMode)
	case theme.ColorNameSeparator:
		return palette.Elements.Separator.Color(e.isDarkMode)
	case theme.ColorNameShadow:
		return palette.Elements.Shadow.Color(e.isDarkMode)
	case theme.ColorNameSuccess:
		return palette.Elements.Success.Color(e.isDarkMode)
	case theme.ColorNameWarning:
		return palette.Elements.Warning.Color(e.isDarkMode)
	}

	return color.Transparent
}
