package particles

import (
	"container/list"
	"testing"

	"project-particles/config"
	"project-particles/particle"
	"project-particles/testutil"
)

// newTestSystem initialise un système vide pour les tests
func newTestSystem() *System {
	return &System{
		Content:      list.New(),
		ParticleDead: nil,
	}
}

// TestSpawn vérifie que le nombre correct de particules est généré selon SpawnRate.
func TestSpawn(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.SpawnRate = 3
	testutil.UseConfig(cfg)

	sys := newTestSystem()
	sys.Update()

	if sys.Content.Len() != 3 {
		t.Errorf("Attendu 3 particules, obtenu %d", sys.Content.Len())
	}
}

// TestReuse vérifie que les particules mortes sont réutilisées.
func TestReuse(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.SpawnRate = 2
	testutil.UseConfig(cfg)

	sys := newTestSystem()

	p := particle.NewParticle()
	sys.Content.PushBack(&p)
	sys.ParticleDead = sys.Content.Front()

	sys.Update()

	first := sys.Content.Front()
	reused := first.Value.(*particle.Particle)
	if reused.Age != 1 {
		t.Errorf("Attendu que la particule morte soit reset, Age = 1, obtenu %v", reused.Age)
	}

	if sys.Content.Len() != 2 {
		t.Errorf("Attendu 2 particules après Update, obtenu %d", sys.Content.Len())
	}

	if sys.ParticleDead != nil {
		t.Errorf("Attendu que ParticleDead soit nil après Update")
	}
}

// TestMovement vérifie que la gravité et la position sont mises à jour
func TestMovement(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.SpawnRate = 1
	cfg.Gravity = 0.1
	cfg.MinVelY = 0
	cfg.MaxVelY = 0
	cfg.SpawnY = 0
	cfg.RandomSpawn = false
	testutil.UseConfig(cfg)

	sys := newTestSystem()
	sys.Update()

	e := sys.Content.Front()
	p := e.Value.(*particle.Particle)

	if p.VelY != cfg.Gravity {
		t.Errorf("Attendu VelY %v, obtenu %v", cfg.Gravity, p.VelY)
	}
	if p.PosY != cfg.Gravity {
		t.Errorf("Attendu PosY %v, obtenu %v", cfg.Gravity, p.PosY)
	}
}

// TestOpacity vérifie que l'opacité est calculée correctement en fonction de l'âge
func TestOpacity(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.Lifetime = 10
	cfg.SpawnRate = 1
	testutil.UseConfig(cfg)

	sys := newTestSystem()
	sys.Update()

	e := sys.Content.Front()
	p := e.Value.(*particle.Particle)

	if p.Opacity <= 0 || p.Opacity > 1 {
		t.Errorf("Attendu opacité entre (0,1], obtenu %v", p.Opacity)
	}
}
