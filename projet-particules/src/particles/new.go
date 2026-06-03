package particles

import (
	"container/list"
	"project-particles/config"
	"project-particles/particle"
)

// NewSystem initialise un système de particules et retourne une structure System.
func NewSystem() System {
	l := list.New()

	for i := 0; i < config.General.InitNumParticles; i++ {
		p := particle.NewParticle()
		l.PushBack(&p)
	}

	return System{Content: l, ParticleDead: nil}
}
