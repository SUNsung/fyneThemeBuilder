package fyneThemeBuilder

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// // // // // // // // // // // // // // // // // //

var (
	DefaultPalette = make(map[fyne.ThemeColorName]*PaletteObj)

	colorArr = []string{
		theme.ColorBlue,
		theme.ColorBrown,
		theme.ColorGray,
		theme.ColorGreen,
		theme.ColorOrange,
		theme.ColorPurple,
		theme.ColorRed,
		theme.ColorYellow,
	}

	//

	colorLightPrimaryBlue   = color.NRGBA{R: 0x29, G: 0x6f, B: 0xf6, A: 0xff}
	colorLightPrimaryBrown  = color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0xff}
	colorLightPrimaryGray   = color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0xff}
	colorLightPrimaryGreen  = color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0xff}
	colorLightPrimaryOrange = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	colorLightPrimaryPurple = color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff}
	colorLightPrimaryRed    = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	colorLightPrimaryYellow = color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0xff}

	colorLightOnPrimaryBlue   = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightOnPrimaryBrown  = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightOnPrimaryGray   = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorLightOnPrimaryGreen  = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorLightOnPrimaryOrange = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorLightOnPrimaryPurple = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightOnPrimaryRed    = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightOnPrimaryYellow = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}

	colorDarkOnPrimaryBlue   = color.NRGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
	colorDarkOnPrimaryBrown  = color.NRGBA{R: 0x5c, G: 0x40, B: 0x33, A: 0xff}
	colorDarkOnPrimaryGray   = color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xff}
	colorDarkOnPrimaryGreen  = color.NRGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xff}
	colorDarkOnPrimaryOrange = color.NRGBA{R: 0xff, G: 0x8c, B: 0x00, A: 0xff}
	colorDarkOnPrimaryPurple = color.NRGBA{R: 0x80, G: 0x00, B: 0x80, A: 0xff}
	colorDarkOnPrimaryRed    = color.NRGBA{R: 0x8b, G: 0x00, B: 0x00, A: 0xff}
	colorDarkOnPrimaryYellow = color.NRGBA{R: 0xff, G: 0xd7, B: 0x00, A: 0xff}

	//

	colorDarkBackground          = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkButton              = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkDisabled            = color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	colorDarkDisabledButton      = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkError               = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	colorDarkForeground          = color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	colorDarkForegroundOnError   = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkForegroundOnSuccess = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkForegroundOnWarning = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkHeaderBackground    = color.NRGBA{R: 0x1b, G: 0x1b, B: 0x1b, A: 0xff}
	colorDarkHover               = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x0f}
	colorDarkInputBackground     = color.NRGBA{R: 0x20, G: 0x20, B: 0x23, A: 0xff}
	colorDarkInputBorder         = color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	colorDarkMenuBackground      = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkOverlayBackground   = color.NRGBA{R: 0x18, G: 0x1d, B: 0x25, A: 0xff}
	colorDarkPlaceholder         = color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	colorDarkPressed             = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	colorDarkScrollBar           = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x99}
	colorDarkSeparator           = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	colorDarkShadow              = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x66}
	colorDarkSuccess             = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	colorDarkWarning             = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}

	colorLightBackground          = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightButton              = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightDisabled            = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightDisabledButton      = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightError               = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	colorLightForeground          = color.NRGBA{R: 0x56, G: 0x56, B: 0x56, A: 0xff}
	colorLightForegroundOnError   = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightForegroundOnSuccess = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightForegroundOnWarning = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightHeaderBackground    = color.NRGBA{R: 0xf9, G: 0xf9, B: 0xf9, A: 0xff}
	colorLightHover               = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x0f}
	colorLightInputBackground     = color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	colorLightInputBorder         = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightMenuBackground      = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightOverlayBackground   = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightPlaceholder         = color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	colorLightPressed             = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x19}
	colorLightScrollBar           = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x99}
	colorLightSeparator           = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightShadow              = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x33}
	colorLightSuccess             = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	colorLightWarning             = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}

	//

	colorLightFocusBlue   = color.NRGBA{R: 0x00, G: 0x6c, B: 0xff, A: 0x2a}
	colorLightFocusBrown  = color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0x7f}
	colorLightFocusGray   = color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0x7f}
	colorLightFocusGreen  = color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x7f}
	colorLightFocusOrange = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0x7f}
	colorLightFocusPurple = color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0x7f}
	colorLightFocusRed    = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0x7f}
	colorLightFocusYellow = color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0x7f}

	colorDarkFocusBlue   = color.NRGBA{R: 0x00, G: 0x47, B: 0xac, A: 0x2a}
	colorDarkFocusBrown  = color.NRGBA{R: 0x4b, G: 0x2c, B: 0x20, A: 0x7f}
	colorDarkFocusGray   = color.NRGBA{R: 0x62, G: 0x62, B: 0x62, A: 0x7f}
	colorDarkFocusGreen  = color.NRGBA{R: 0x5a, G: 0x91, B: 0x2f, A: 0x7f}
	colorDarkFocusOrange = color.NRGBA{R: 0xc7, G: 0x60, B: 0x00, A: 0x7f}
	colorDarkFocusPurple = color.NRGBA{R: 0x72, G: 0x1b, B: 0x82, A: 0x7f}
	colorDarkFocusRed    = color.NRGBA{R: 0xab, G: 0x00, B: 0x1f, A: 0x7f}
	colorDarkFocusYellow = color.NRGBA{R: 0xb2, G: 0xa6, B: 0x29, A: 0x7f}

	//

	colorLightSelectionBlue   = color.NRGBA{R: 0x00, G: 0x6c, B: 0xff, A: 0x40}
	colorLightSelectionBrown  = color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0x3f}
	colorLightSelectionGray   = color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0x3f}
	colorLightSelectionGreen  = color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x3f}
	colorLightSelectionOrange = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0x3f}
	colorLightSelectionPurple = color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0x3f}
	colorLightSelectionRed    = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0x3f}
	colorLightSelectionYellow = color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0x3f}

	colorDarkSelectionBlue   = color.NRGBA{R: 0x00, G: 0x47, B: 0xac, A: 0x40}
	colorDarkSelectionBrown  = color.NRGBA{R: 0x4b, G: 0x2c, B: 0x20, A: 0x3f}
	colorDarkSelectionGray   = color.NRGBA{R: 0x62, G: 0x62, B: 0x62, A: 0x3f}
	colorDarkSelectionGreen  = color.NRGBA{R: 0x5a, G: 0x91, B: 0x2f, A: 0x3f}
	colorDarkSelectionOrange = color.NRGBA{R: 0xc7, G: 0x60, B: 0x00, A: 0x3f}
	colorDarkSelectionPurple = color.NRGBA{R: 0x72, G: 0x1b, B: 0x82, A: 0x3f}
	colorDarkSelectionRed    = color.NRGBA{R: 0xab, G: 0x00, B: 0x1f, A: 0x3f}
	colorDarkSelectionYellow = color.NRGBA{R: 0xb2, G: 0xa6, B: 0x29, A: 0x3f}
)

/* ###################### */

func init() {
	add := func(name fyne.ThemeColorName) {
		buf := PaletteElementsObj{
			Foreground: &PalettePrimaryColorForegroundObj{
				Global: &PalettePowerSideObj{
					colorLightForeground,
					colorDarkForeground,
				},
				Primary: &PalettePowerSideObj{
					colorLightOnPrimaryRed,
					colorDarkOnPrimaryRed,
				},
				Error: &PalettePowerSideObj{
					colorLightForegroundOnError,
					colorDarkForegroundOnError,
				},
				Success: &PalettePowerSideObj{
					colorLightForegroundOnSuccess,
					colorDarkForegroundOnSuccess,
				},
				Warning: &PalettePowerSideObj{
					colorLightForegroundOnWarning,
					colorDarkForegroundOnWarning,
				},
			},

			Focus: &PalettePowerSideObj{
				colorLightFocusRed,
				colorDarkFocusRed,
			},
			Selection: &PalettePowerSideObj{
				colorLightSelectionRed,
				colorDarkSelectionRed,
			},
			Background: &PalettePowerSideObj{
				colorLightBackground,
				colorDarkBackground,
			},
			Button: &PalettePowerSideObj{
				colorLightButton,
				colorDarkButton,
			},
			Disabled: &PalettePowerSideObj{
				colorLightDisabled,
				colorDarkDisabled,
			},
			DisabledButton: &PalettePowerSideObj{
				colorLightDisabledButton,
				colorDarkDisabledButton,
			},
			Error: &PalettePowerSideObj{
				colorLightError,
				colorDarkError,
			},
			HeaderBackground: &PalettePowerSideObj{
				colorLightHeaderBackground,
				colorDarkHeaderBackground,
			},
			Hover: &PalettePowerSideObj{
				colorLightHover,
				colorDarkHover,
			},
			InputBackground: &PalettePowerSideObj{
				colorLightInputBackground,
				colorDarkInputBackground,
			},
			InputBorder: &PalettePowerSideObj{
				colorLightInputBorder,
				colorDarkInputBorder,
			},
			MenuBackground: &PalettePowerSideObj{
				colorLightMenuBackground,
				colorDarkMenuBackground,
			},
			OverlayBackground: &PalettePowerSideObj{
				colorLightOverlayBackground,
				colorDarkOverlayBackground,
			},
			Placeholder: &PalettePowerSideObj{
				colorLightPlaceholder,
				colorDarkPlaceholder,
			},
			Pressed: &PalettePowerSideObj{
				colorLightPressed,
				colorDarkPressed,
			},
			ScrollBar: &PalettePowerSideObj{
				colorLightScrollBar,
				colorDarkScrollBar,
			},
			Separator: &PalettePowerSideObj{
				colorLightSeparator,
				colorDarkSeparator,
			},
			Shadow: &PalettePowerSideObj{
				colorLightShadow,
				colorDarkShadow,
			},
			Success: &PalettePowerSideObj{
				colorLightSuccess,
				colorDarkSuccess,
			},
			Warning: &PalettePowerSideObj{
				colorLightWarning,
				colorDarkWarning,
			},
		}

		col := &PalettePrimaryColorObj{Color: &PalettePowerSideObj{}}
		col.Name = name

		switch name {
		case theme.ColorBlue:
			col.Color.Dark = colorLightPrimaryBlue
			col.Color.Light = colorLightPrimaryBlue
			buf.Foreground.Primary.Dark = colorDarkOnPrimaryBlue
			buf.Foreground.Primary.Light = colorLightOnPrimaryBlue
			buf.Focus.Dark = colorDarkFocusBlue
			buf.Focus.Light = colorLightFocusBlue
			buf.Selection.Dark = colorDarkSelectionBlue
			buf.Selection.Light = colorLightSelectionBlue

		case theme.ColorBrown:
			col.Color.Dark = colorLightPrimaryBrown
			col.Color.Light = colorLightPrimaryBrown
			buf.Foreground.Primary.Dark = colorLightOnPrimaryBrown
			buf.Foreground.Primary.Light = colorDarkOnPrimaryBrown
			buf.Focus.Dark = colorDarkFocusBrown
			buf.Focus.Light = colorLightFocusBrown
			buf.Selection.Light = colorLightSelectionBrown
			buf.Selection.Dark = colorDarkSelectionBrown

		case theme.ColorGray:
			col.Color.Dark = colorLightPrimaryGray
			col.Color.Light = colorLightPrimaryGray
			buf.Foreground.Primary.Dark = colorLightOnPrimaryGray
			buf.Foreground.Primary.Light = colorDarkOnPrimaryGray
			buf.Focus.Dark = colorDarkFocusGray
			buf.Focus.Light = colorLightFocusGray
			buf.Selection.Light = colorLightSelectionGray
			buf.Selection.Dark = colorDarkSelectionGray

		case theme.ColorGreen:
			col.Color.Dark = colorLightPrimaryGreen
			col.Color.Light = colorLightPrimaryGreen
			buf.Foreground.Primary.Dark = colorLightOnPrimaryGreen
			buf.Foreground.Primary.Light = colorDarkOnPrimaryGreen
			buf.Focus.Dark = colorDarkFocusGreen
			buf.Focus.Light = colorLightFocusGreen
			buf.Selection.Light = colorLightSelectionGreen
			buf.Selection.Dark = colorDarkSelectionGreen

		case theme.ColorOrange:
			col.Color.Dark = colorLightPrimaryOrange
			col.Color.Light = colorLightPrimaryOrange
			buf.Foreground.Primary.Dark = colorLightOnPrimaryOrange
			buf.Foreground.Primary.Light = colorDarkOnPrimaryOrange
			buf.Focus.Dark = colorDarkFocusOrange
			buf.Focus.Light = colorLightFocusOrange
			buf.Selection.Light = colorLightSelectionOrange
			buf.Selection.Dark = colorDarkSelectionOrange

		case theme.ColorPurple:
			col.Color.Dark = colorLightPrimaryPurple
			col.Color.Light = colorLightPrimaryPurple
			buf.Foreground.Primary.Dark = colorLightOnPrimaryPurple
			buf.Foreground.Primary.Light = colorDarkOnPrimaryPurple
			buf.Focus.Dark = colorDarkFocusPurple
			buf.Focus.Light = colorLightFocusPurple
			buf.Selection.Light = colorLightSelectionPurple
			buf.Selection.Dark = colorDarkSelectionPurple

		case theme.ColorRed:
			col.Color.Dark = colorLightPrimaryRed
			col.Color.Light = colorLightPrimaryRed
			buf.Foreground.Primary.Dark = colorLightOnPrimaryRed
			buf.Foreground.Primary.Light = colorDarkOnPrimaryRed
			buf.Focus.Dark = colorDarkFocusRed
			buf.Focus.Light = colorLightFocusRed
			buf.Selection.Light = colorLightSelectionRed
			buf.Selection.Dark = colorDarkSelectionRed

		case theme.ColorYellow:
			col.Color.Dark = colorLightPrimaryYellow
			col.Color.Light = colorLightPrimaryYellow
			buf.Foreground.Primary.Dark = colorLightOnPrimaryYellow
			buf.Foreground.Primary.Light = colorDarkOnPrimaryYellow
			buf.Focus.Dark = colorDarkFocusYellow
			buf.Focus.Light = colorLightFocusYellow
			buf.Selection.Light = colorLightSelectionYellow
			buf.Selection.Dark = colorDarkSelectionYellow
		}

		DefaultPalette[name] = &PaletteObj{
			PrimaryColor: col,
			Elements:     &buf,
		}
	}

	for _, c := range colorArr {
		add(fyne.ThemeColorName(c))
	}
}
