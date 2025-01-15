package fyneThemeBuilder

import (
	"fyne.io/fyne/v2"
	"image/color"
)

// // // // // // // // // // // // // // // // // //

type PalettePowerSideObj struct {
	Light color.NRGBA `json:"light"`
	Dark  color.NRGBA `json:"dark"`
}

func (p *PalettePowerSideObj) Color(isDark bool) color.NRGBA {
	if isDark {
		return p.Dark
	}
	return p.Light
}

type PalettePrimaryColorForegroundObj struct {
	Global  *PalettePowerSideObj `json:"global"`
	Primary *PalettePowerSideObj `json:"primary"`
	Error   *PalettePowerSideObj `json:"error"`
	Success *PalettePowerSideObj `json:"success"`
	Warning *PalettePowerSideObj `json:"warning"`
}

type PaletteElementsObj struct {
	Foreground *PalettePrimaryColorForegroundObj `json:"foreground"`

	Focus             *PalettePowerSideObj `json:"focus"`
	Selection         *PalettePowerSideObj `json:"selection"`
	Background        *PalettePowerSideObj `json:"background"`
	Button            *PalettePowerSideObj `json:"button"`
	Disabled          *PalettePowerSideObj `json:"disabled"`
	DisabledButton    *PalettePowerSideObj `json:"disabled_button"`
	Error             *PalettePowerSideObj `json:"error"`
	HeaderBackground  *PalettePowerSideObj `json:"header_background"`
	Hover             *PalettePowerSideObj `json:"hover"`
	InputBackground   *PalettePowerSideObj `json:"input_background"`
	InputBorder       *PalettePowerSideObj `json:"input_border"`
	MenuBackground    *PalettePowerSideObj `json:"menu_background"`
	OverlayBackground *PalettePowerSideObj `json:"overlay_background"`
	Placeholder       *PalettePowerSideObj `json:"placeholder"`
	Pressed           *PalettePowerSideObj `json:"pressed"`
	ScrollBar         *PalettePowerSideObj `json:"scroll_bar"`
	Separator         *PalettePowerSideObj `json:"separator"`
	Shadow            *PalettePowerSideObj `json:"shadow"`
	Success           *PalettePowerSideObj `json:"success"`
	Warning           *PalettePowerSideObj `json:"warning"`
}

type PalettePrimaryColorObj struct {
	Color *PalettePowerSideObj `json:"color"`
	Name  fyne.ThemeColorName  `json:"name"`
}

type PaletteObj struct {
	PrimaryColor *PalettePrimaryColorObj `json:"primary_color"`
	Elements     *PaletteElementsObj     `json:"elements"`
}
