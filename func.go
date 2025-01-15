package fyneThemeBuilder

import "fyne.io/fyne/v2"

// // // // // // // // // // // // // // // // // //

// SetDarkMode setting the color mode
//
// defaults to fyne by default
func (e *ThemeInterfaceObj) SetDarkMode(b bool) bool {
	e.propertyLock.Lock()
	defer e.propertyLock.Unlock()
	e.isDarkMode = b
	return true
}

func (e *ThemeInterfaceObj) GetDarkMode() bool {
	e.propertyLock.RLock()
	defer e.propertyLock.RUnlock()
	return e.isDarkMode
}

// SetPrimaryColor sets the primary color
//
// if there is no such color, then the change will not be applied
func (e *ThemeInterfaceObj) SetPrimaryColor(n fyne.ThemeColorName) bool {
	_, ok := e.paletteMap[n]
	if !ok {
		return false
	}

	e.propertyLock.Lock()
	defer e.propertyLock.Unlock()
	e.primaryColor = n

	return true
}

func (e *ThemeInterfaceObj) GetPrimaryColor() fyne.ThemeColorName {
	e.propertyLock.RLock()
	defer e.propertyLock.RUnlock()
	return e.primaryColor
}
