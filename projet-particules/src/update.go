package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update met à jour l'état du jeu pour une frame.
//
// Cette méthode est appelée automatiquement par Ebiten environ 60 fois
// par seconde et gère :
//  1. Les interactions clavier.
//  2. La mise à jour du système de particules.
//  3. La mise à jour de la grille de collision et la résolution des collisions
//  4. La mise à jour de l'interface utilisateur
func (g *game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.ui.Visible = !g.ui.Visible
	}

	g.system.Update()

	if config.General.Collision {
		g.grid.Build(g.system.Content)
		g.grid.ResolveCollisions()
	}

	if g.ui.Visible {
		g.ui.Update()
	}

	return nil
}
