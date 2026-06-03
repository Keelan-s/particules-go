package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

// Draw affiche l'interface utilisateur et l'ensemble de ses widgets à l'écran.
func (ui *UI) Draw(screen *ebiten.Image) {

	vector.FillRect(
		screen,
		float32(ui.X), float32(ui.Y),
		float32(ui.ColumnWidth*ui.NumColumn),
		float32(ui.LineHeight*(ui.NumLine+1)),
		color.RGBA{25, 25, 25, 255},
		true)

	for _, widget := range ui.Widgets {
		widget.draw(screen, ui)
	}
}

// draw affiche le widget Toggle.
func (t *Toggle) draw(screen *ebiten.Image, ui *UI) {
	col := color.RGBA{80, 80, 80, 255}
	if *t.Value {
		col = color.RGBA{0, 255, 0, 255}
	}

	size := float32(ui.LineHeight) * 0.8

	vector.FillRect(
		screen,
		float32(t.X+ui.Margin+ui.LineHeight*2/10),
		float32(t.Y+ui.LineHeight*2/10),
		size,
		size,
		col,
		true,
	)

	text.Draw(
		screen,
		t.Label,
		basicfont.Face7x13,
		t.X+ui.LineHeight+ui.Margin*2,
		t.Y+ui.LineHeight*8/10,
		color.White,
	)
}

// draw affiche le widget TextField.
func (t *TextField) draw(screen *ebiten.Image, ui *UI) {
	col := color.RGBA{255, 255, 255, 255}
	if t.Active {
		col = color.RGBA{255, 0, 0, 255}
	}

	text.Draw(
		screen,
		fmt.Sprintf("%s: %s", t.Label, t.Buffer),
		basicfont.Face7x13,
		t.X+ui.Margin*2+ui.LineHeight,
		t.Y+ui.LineHeight*8/10,
		col,
	)
}
