package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var DisabledTx = color.NRGBA{R: 125, G: 135, B: 135, A: 255}

type baseTheme struct{}

var _ fyne.Theme = (*baseTheme)(nil)

func (bt baseTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameDisabled:
		return DisabledTx
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (bt baseTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (bt baseTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (bt baseTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
