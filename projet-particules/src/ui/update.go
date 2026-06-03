package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update met à jour tous les widgets de l'UI.
//
// Cette méthode itère sur chaque widget et appelle sa fonction `update`,
// permettant la gestion des interactions utilisateur (clics, clavier, etc...).
func (ui *UI) Update() {
	for _, widget := range ui.Widgets {
		widget.update(ui)
	}
}

// update gère l'interaction de l'utilisateur avec un Toggle.
func (t *Toggle) update(ui *UI) {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}

	mx, my := ebiten.CursorPosition()

	size := (ui.LineHeight * 8) / 10
	x := t.X + ui.Margin + (ui.LineHeight*2)/10
	y := t.Y + (ui.LineHeight*2)/10

	if pointInRect(mx, my, x, y, size, size) {
		*t.Value = !*t.Value
	}
}

// handleClick détermine si un TextField doit devenir actif.
func (t *TextField) handleClick(ui *UI) {

	mx, my := ebiten.CursorPosition()
	x, y := t.X+ui.Margin, t.Y+ui.LineHeight/10
	w, h := ui.ColumnWidth-ui.Margin*2, (ui.LineHeight*8)/10

	if pointInRect(mx, my, x, y, w, h) {
		t.Active = true
		t.Buffer = ""
	} else {
		t.Active = false
	}
}

// handleKeyboard gère la saisie clavier pour un TextField actif.
//
// Les caractères saisis sont ajoutés au buffer, Backspace supprime
// le dernier caractère, Enter applique la valeur via `Apply` et
// désactive le champ, Escape désactive le champ sans appliquer.
func (t *TextField) handleKeyboard() {

	inputRunes := ebiten.AppendInputChars([]rune{})
	t.Buffer += string(inputRunes)

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(t.Buffer) > 0 {
		t.Buffer = t.Buffer[:len(t.Buffer)-1]
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		t.Apply(t.Buffer)
		t.Active = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		t.Active = false
	}
}

// update gère l'état d'un TextField pour une frame.
func (t *TextField) update(ui *UI) {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		t.handleClick(ui)
	}

	if !t.Active {
		t.Buffer = t.Sync()
		return
	}

	t.handleKeyboard()

}
