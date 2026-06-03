package particle

import (
	"log"
	"math/rand"
	"project-particles/assets"
	"project-particles/config"
)

// randRange retourne un nombre aléatoire compris entre min et max.
func randRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// randSignedRange retourne un nombre aléatoire compris entre min et max,
// avec une chance sur deux qu'il soit négatif.
func randSignedRange(min, max float64) float64 {
	if rand.Intn(2) == 0 {
		return randRange(min, max)
	}
	return -randRange(min, max)
}

// initSpawn initialise la position de départ de la particule selon la config.
func initSpawn(p *Particle, cfg config.Config) {
	if cfg.RandomSpawn {
		p.PosX = randRange(0, float64(cfg.WindowSizeX))
		p.PosY = randRange(0, float64(cfg.WindowSizeY))
		return
	}

	p.PosX = float64(cfg.SpawnX)
	p.PosY = float64(cfg.SpawnY)
}

// initSize initialise la taille de la particule selon la config.
func initSize(p *Particle, cfg config.Config) {
	if cfg.RandomSize {
		p.ScaleX = randRange(cfg.MinSizeX, cfg.MaxSizeX)
		p.ScaleY = randRange(cfg.MinSizeY, cfg.MaxSizeY)
		return
	}

	p.ScaleX = cfg.SizeX
	p.ScaleY = cfg.SizeY
}

// initVelocity initialise la vitesse de la particule selon la config.
func initVelocity(p *Particle, cfg config.Config) {
	p.VelX = randSignedRange(cfg.MinVelX, cfg.MaxVelX)
	p.VelY = randSignedRange(cfg.MinVelY, cfg.MaxVelY)
	p.AngVel = randSignedRange(cfg.MinAngVel, cfg.MaxAngVel)
}

// initColor initialise la couleur de la particule selon la config.
func initColor(p *Particle, cfg config.Config) {
	if cfg.RandomColor {
		p.ColorR = rand.Float64()
		p.ColorG = rand.Float64()
		p.ColorB = rand.Float64()
		return
	}

	p.ColorR = cfg.ColorR / 255
	p.ColorG = cfg.ColorG / 255
	p.ColorB = cfg.ColorB / 255
}

// initParticle initialise complètement une particule : position, taille,
// vélocité, couleur, rotation, opacité, âge et dimensions basées sur l'image.
func initParticle(p *Particle) {
	cfg := config.General

	initSpawn(p, cfg)
	initSize(p, cfg)
	initVelocity(p, cfg)
	initColor(p, cfg)

	p.Rotation = 0
	p.Opacity = 1
	p.Age = 0

	if assets.ParticleImage == nil {
		log.Println(" Attention : l'image des particules n'a pas pu être chargée ! ")
		p.Width = 0
		p.Height = 0
		return
	}

	img := assets.ParticleImage.Bounds()
	p.Width = p.ScaleX * float64(img.Dx())
	p.Height = p.ScaleY * float64(img.Dy())
}
