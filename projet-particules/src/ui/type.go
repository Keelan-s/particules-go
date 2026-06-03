package ui

import "github.com/hajimehoshi/ebiten/v2"

// UI représente un conteneur d'interface utilisateur.
//
// Il définit le layout global (position, dimensions, colonnes, ...)
// et contient l'ensemble des widgets à rendre et mettre à jour.
type UI struct {
	Visible     bool
	X, Y        int
	ColumnWidth int
	LineHeight  int
	Margin      int
	NumColumn   int
	NumLine     int
	Widgets     []Widget
}

// Widget définit les fonctions minimales d'un élément d'interface utilisateur
//
// Chaque widget est responsable de sa mise à jour logique (update)
// et de son rendu graphique (draw).
type Widget interface {
	update(*UI)
	draw(*ebiten.Image, *UI)
}

// TextField représente un champ de saisie texte.
type TextField struct {
	Label  string
	X, Y   int
	Buffer string
	Active bool
	Sync   func() string
	Apply  func(string)
}

// Toggle représente un bouton on/off lié à une valeur booléenne.
type Toggle struct {
	Label string
	X, Y  int
	Value *bool
}

// LayoutCursor gère le positionnement séquentiel des widgets
// dans un conteneur UI en respectant le layout en lignes et colonnes.
type LayoutCursor struct {
	X, Y   int
	Parent *UI
}
