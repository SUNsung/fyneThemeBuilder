![Fork GitHub Release](https://img.shields.io/github/v/release/SUNsung/fyneThemeBuilder)
[![Go Report Card](https://goreportcard.com/badge/github.com/SUNsung/fyneThemeBuilder)](https://goreportcard.com/report/github.com/SUNsung/fyneThemeBuilder)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/SUNsung/fyneThemeBuilder?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/SUNsung/fyneThemeBuilder?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/SUNsung/fyneThemeBuilder)

# fyne.io Theme Builder

This library is written to provide a user-friendly interface to fune.io application designs

The module works with palette collections. One palette contains colors for day and night mode.

There is a default palette collection `DefaultPalette`. It corresponds to the current palette of fyne.io

## Constructors

All constructors return a structure corresponding to the `fyne.Theme` interface.

#### `New(map[fyne.ThemeColorName]*PaletteObj) *ThemeInterfaceObj`
Initializing a collection of palettes from a structure.

#### `NewOne(primaryColor *PalettePrimaryColorObj, elements *PaletteElementsObj) *ThemeInterfaceObj`
Initializes a single palette.

#### `NewDefault() *ThemeInterfaceObj`
Initializes default palettes

## Methods

#### `SetDarkMode(b bool) bool`
Setting the color mode

#### `GetDarkMode() bool`
Getting the color mode state

#### `SetPrimaryColor(n fyne.ThemeColorName) bool`
Set the main theme color.
- this method switches the topic by title
- if the passed name is not in the palettes it will return false

#### `GetPrimaryColor() fyne.ThemeColorName`
Get the name of the main theme color

#### `GetPalette() *PaletteObj`
Get current palette

#### `GetMap() map[fyne.ThemeColorName]*PaletteObj`
Get all initialized palettes

---
## Examples

#### Red standard light theme
```go
myApp := app.New()
myWindow := myApp.NewWindow("window name")

th := fyneThemeBuilder.NewDefault()
myApp.Settings().SetTheme(th)

th.SetDarkMode(false)
th.SetPrimaryColor(theme.ColorRed)

myWindow.SetContent(container.NewVBox(widget.NewLabel("Hello!")))
myWindow.Resize(fyne.NewSize(400, 200))
myWindow.ShowAndRun()
```

#### Text size control
```go
myApp := app.New()
myWindow := myApp.NewWindow("window name")

scale := float32(1)

th := fyneThemeBuilder.NewDefault()
myApp.Settings().SetTheme(th)

th.HookSize = func(name fyne.ThemeSizeName) float32 {
    val := func() float32 {
        switch name {
        case theme.SizeNameSeparatorThickness:
            return 1
        case theme.SizeNameInlineIcon:
            return 20
        case theme.SizeNameInnerPadding:
            return 8
        case theme.SizeNameLineSpacing:
            return 4
        case theme.SizeNamePadding:
            return 4
        case theme.SizeNameScrollBar:
            return 12
        case theme.SizeNameScrollBarSmall:
            return 3
        case theme.SizeNameText:
            return 14
        case theme.SizeNameHeadingText:
            return 24
        case theme.SizeNameSubHeadingText:
            return 18
        case theme.SizeNameCaptionText:
            return 11
        case theme.SizeNameInputBorder:
            return 1
        case theme.SizeNameInputRadius:
            return 5
        case theme.SizeNameSelectionRadius:
            return 3
        case theme.SizeNameScrollBarRadius:
            return 3
        default:
            return 0
        }
    }

    if scale < 0 {
        return val()
    }

    return val() * scale
}

go func() {
    time.Sleep(time.Second*5)
    scale = 2
	myWindow.Content().Refresh()
}()

myWindow.SetContent(container.NewVBox(widget.NewLabel("Hello!")))
myWindow.Resize(fyne.NewSize(400, 200))
myWindow.ShowAndRun()
```
- a hook that substitutes the Size() method in the theme has been initialized
- 5 seconds after the window is launched, the text size will be doubled.


#### Loading and saving the theme
```go
type SaveObj struct {
    IsDarkMode   bool
    PrimaryColor fyne.ThemeColorName
    PaletteMap   map[fyne.ThemeColorName]*fyneThemeBuilder.PaletteObj
}

data := new(SaveObj)
filename := "filename.gob"

data.IsDarkMode = true
data.PrimaryColor = theme.ColorBlue
data.PaletteMap = fyneThemeBuilder.DefaultPalette

file, err := os.Open(filename)
if err == nil {
    gob.NewDecoder(file).Decode(data)
}

myApp := app.New()
myWindow := myApp.NewWindow("window name")

th := fyneThemeBuilder.New(data.PaletteMap)
myApp.Settings().SetTheme(th)
th.SetPrimaryColor(data.PrimaryColor)
th.SetDarkMode(data.IsDarkMode)

defer func() {
    data.PaletteMap = th.GetMap()
    data.IsDarkMode = th.GetDarkMode()
    data.PrimaryColor = th.GetPrimaryColor()
    
    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    
    if err := encoder.Encode(data); err != nil {
        return
    }

    file, err = os.Create(filename)
    if err != nil {
        return
    }

    file.Write(buffer.Bytes())
    file.Close()
}()

myWindow.SetContent(container.NewVBox(widget.NewLabel("Hello!")))
myWindow.Resize(fyne.NewSize(400, 200))
myWindow.ShowAndRun()
```
- the topic is saved to a file when closing
- when opening the file, if the topic has been read, the read topic is applied, otherwise the default topic is used.

#### a light window that after 5 seconds will change the background color to blue
```go
myApp := app.New()
myWindow := myApp.NewWindow("window name")

th := fyneThemeBuilder.NewDefault()
myApp.Settings().SetTheme(th)
th.SetDarkMode(false)
th.SetPrimaryColor("red")

go func() {
time.Sleep(time.Second * 1)
defTh := th.GetPalette()
defTh.Elements.Background.Light = color.NRGBA{R: 0x29, G: 0x6f, B: 0xf6, A: 0xff}
myWindow.Content().Refresh()
}()

myWindow.SetContent(container.NewVBox(widget.NewLabel("Hello!")))
myWindow.Resize(fyne.NewSize(400, 200))
myWindow.ShowAndRun()
```