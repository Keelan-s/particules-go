package main

import (
	"fmt"
	"image/color"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particle"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Draw affiche l'état actuel du système de particules sur l'écran.
//
// Cette fonction est appelée automatiquement par Ebiten environ 60 fois
// par seconde. Elle gère :
//  1. Le rendu de chaque particule active avec rotation, échelle, position
//     et couleur/opacity. La rotation s'effectue au centre de la particule et non à l'origine.
//  2. L'affichage optionnel des grilles de collision et debug si activé.
//  3. Le rendu de l'interface utilisateur (UI) si elle est visible.
func (g *game) Draw(screen *ebiten.Image) {
	cfg := config.General
	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particle.Particle)
		if !ok {
			continue
		}

		if p.IsDead() {
			break
		}

		options := ebiten.DrawImageOptions{}

		options.GeoM.Scale(p.ScaleX, p.ScaleY)
		options.GeoM.Translate(-p.Width/2, -p.Height/2)
		options.GeoM.Rotate(p.Rotation)
		options.GeoM.Translate(p.PosX+p.Width/2, p.PosY+p.Height/2)

		opacity := float32(p.Opacity) // Permet de bien voir la transparence sur fond noir
		options.ColorScale.Scale(
			opacity*float32(p.ColorR),
			opacity*float32(p.ColorG),
			opacity*float32(p.ColorB),
			1,
		)

		screen.DrawImage(assets.ParticleImage, &options)
	}

	if cfg.Debug {
		if cfg.Collision {
			for coord := range g.grid.Grid {
				x := coord.X * cfg.CellSize
				y := coord.Y * cfg.CellSize

				vector.FillRect(
					screen,
					float32(x),
					float32(y),
					float32(cfg.CellSize),
					float32(cfg.CellSize),
					color.RGBA{25, 25, 25, 0},
					true,
				)
			}
		}

		tps := fmt.Sprintf("%.2f", ebiten.ActualTPS())
		ebitenutil.DebugPrint(screen, fmt.Sprint(tps, "\n", g.system.Content.Len()))
	}

	if g.ui.Visible {
		g.ui.Draw(screen)
	}

}
