package main

import (
	"project-particles/collision"
	"project-particles/particles"
	"project-particles/ui"
)

// game représente l'état principal du jeu et implémente l'interface Game d'Ebiten.
//
// Ce type central contient :
//   - system : le système de particules gérant les particules actives,
//   - grid   : la grille de collision pour les particules,
//   - ui     : l'interface utilisateur pour modifier dynamiquement les paramètres.
//
// Il dispose d'une méthode Update, d'une méthode Draw et d'une méthode
// Layout. C'est un élément de ce type qui est utilisé pour mettre à jour et
// afficher un système de particules.
type game struct {
	system particles.System
	grid   collision.CollisionGrid
	ui     ui.UI
}
