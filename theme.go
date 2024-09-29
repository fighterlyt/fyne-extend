package fyne_extend

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	extendTheme "fyne.io/x/fyne/theme"
)

var (
	//go:embed MiSans-Normal.ttf
	miSans            []byte
	ResourceMiSansTTF = &fyne.StaticResource{
		StaticName:    "MiSans-Normal.ttf",
		StaticContent: miSans,
	}
)

func GetFont() []byte {
	return miSans
}

const (
	defaultFactor = 1
)

type Theme struct {
	extendTheme.Adwaita
	factor float32
}

func NewTheme(factor float32) *Theme {
	if factor <= 0 {
		factor = defaultFactor
	}

	return &Theme{
		factor: factor,
	}
}

func (m Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}

		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m Theme) Font(style fyne.TextStyle) fyne.Resource {
	return ResourceMiSansTTF
}

func (m Theme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 14 * m.factor
	}
	return theme.DefaultTheme().Size(name) * m.factor
}

func (m *Theme) SetFactor(factor float32) {
	m.factor = factor
}
