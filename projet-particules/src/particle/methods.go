package particle

import "project-particles/config"

// idTracker sert à donner un ID unique à chaque particule.
var idTracker int

// NewParticle initialise une nouvelle particule en créant une nouvelle.
func NewParticle() Particle {
	p := Particle{
		ID: idTracker,
	}
	idTracker++

	initParticle(&p)
	return p
}

// Reset permet de recycler une particule déjà existante.
func (p *Particle) Reset() {
	initParticle(p)
}

// IsDead permet de savoir si une particule est 'morte'.
func (p Particle) IsDead() bool {
	cfg := config.General

	m := cfg.Margin
	w := float64(cfg.WindowSizeX)
	h := float64(cfg.WindowSizeY)

	return p.Age >= cfg.Lifetime || p.PosX < -m || p.PosX > w+m || p.PosY < -m || p.PosY > h+m
}
