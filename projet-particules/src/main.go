package main

import (
	"log"
	"project-particles/assets"
	"project-particles/collision"
	"project-particles/config"
	"project-particles/particles"
	"project-particles/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

// main est la fonction principale du projet.
//
// Elle effectue les étapes suivantes :
//  1. Lecture de la configuration
//  2. Chargement des assets
//  3. Initialisation de la fenêtre
//  4. Création d'un système de particule, d'une grille de collision
//     et d'une UI dans une structure game
//  5. Lancement de la boucle principale du jeu via Ebiten.
func main() {

	config.Get("config.json")
	assets.Get()

	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)

	g := game{system: particles.NewSystem(), grid: collision.Init(), ui: ui.Init()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
