package particles

import (
	"project-particles/config"
	"project-particles/particle"
	"project-particles/testutil"
	"testing"
)

// TestNewSystem vérifie que NewSystem initialise correctement le système.
func TestNewSystem(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.InitNumParticles = 5
	testutil.UseConfig(cfg)

	sys := NewSystem()

	if sys.Content.Len() != cfg.InitNumParticles {
		t.Errorf("Attendu %d particules, obtenu %d", cfg.InitNumParticles, sys.Content.Len())
	}

	if sys.ParticleDead != nil {
		t.Errorf("Attendu ParticleDead == nil, obtenu %v", sys.ParticleDead)
	}

	for e := sys.Content.Front(); e != nil; e = e.Next() {
		_, ok := e.Value.(*particle.Particle)
		if !ok {
			t.Errorf("Attend que des *Particle, obtenu %T", e.Value)
		}
	}
}
