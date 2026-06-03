package particles

import (
	"container/list"
	"math"
	"math/rand"
	"project-particles/config"
	"project-particles/particle"
)

// Update met à jour l'état du système de particules pour une frame.
//
// Cette méthode est appelée environ 60 fois par seconde par la boucle principale.
//
// Fonctionnement :
//  1. Calcul du nombre de particules à générer selon SpawnRate
//     (partie entière + fraction utilisée comme probabilité aléatoire).
//  2. Génération de nouvelles particules :
//     - Réutilisation des particules mortes via ParticleDead si possible (Reset).
//     - Sinon création d'une nouvelle particule et ajout à la liste.
//  3. Parcours de toutes les particules vivantes (avant ParticleDead) :
//     - Application de la gravité (VelY).
//     - Mise à jour de la position (PosX, PosY) et son oriention (Rotation) dans l'intervalle ]-pi; pi].
//     - Mise à jour de l'opacité selon l'âge (fonction sigmoïde).
//     - Gestion de la mort de la particule : swap avec la dernière particule vivante
//     et mise à jour de ParticleDead pour réutilisation ultérieure.
func (s *System) Update() {

	cfg := config.General
	spawnRate := cfg.SpawnRate

	spawnInt := int(spawnRate)
	spawnFrac := spawnRate - float64(spawnInt)

	if rand.Float64() < spawnFrac {
		spawnInt++
	}

	for i := 0; i < spawnInt; i++ {
		if s.ParticleDead != nil {
			p := s.ParticleDead.Value.(*particle.Particle)
			p.Reset()
			s.ParticleDead = s.ParticleDead.Next()
		} else {
			p := particle.NewParticle()
			s.Content.PushBack(&p)
		}
	}

	e := s.Content.Front()
	for e != s.ParticleDead {
		next := e.Next()

		p, ok := e.Value.(*particle.Particle)
		if !ok {
			e = next
			continue
		}

		p.Rotation += p.AngVel
		if p.Rotation > math.Pi {
			p.Rotation -= 2 * math.Pi
		}

		p.VelY += cfg.Gravity
		p.PosY += p.VelY
		p.PosX += p.VelX

		lifetime := float64(cfg.Lifetime)
		rate := 16 / lifetime
		age := float64(p.Age)

		p.Opacity = 1.0 / (1.0 + math.Exp(rate*(age-lifetime*0.8)))
		p.Age++

		if p.IsDead() {
			var lastAlived *list.Element
			if s.ParticleDead == nil {
				lastAlived = s.Content.Back()
			} else {
				lastAlived = s.ParticleDead.Prev()
			}

			e.Value, lastAlived.Value = lastAlived.Value, e.Value
			s.ParticleDead = lastAlived

			if e == lastAlived {
				break
			}
			continue
		}

		e = next
	}
}
