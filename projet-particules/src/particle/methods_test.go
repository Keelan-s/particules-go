package particle

import (
	"project-particles/config"
	"project-particles/testutil"
	"testing"
)

// TestNewParticle vérifie que NewParticle crée une particule valide avec un ID unique
func TestNewParticle(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	testutil.UseConfig(cfg)

	idBefore := idTracker
	p := NewParticle()

	if p.ID != idBefore {
		t.Errorf("Attendu ID %v, obtenu %v", idBefore, p.ID)
	}
	if p.Age != 0 {
		t.Errorf("Attendu Age 0, obtenu %v", p.Age)
	}
	if p.Opacity != 1 {
		t.Errorf("Attendu Opacity 1, obtenu %v", p.Opacity)
	}
}

// TestReset vérifie que Reset réinitialise les propriétés importantes
func TestReset(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	testutil.UseConfig(cfg)

	p := NewParticle()
	p.Age = 10
	p.Opacity = 0.5
	p.Reset()

	if p.Age != 0 {
		t.Errorf("Attendu Age 0 après Reset, obtenu %v", p.Age)
	}
	if p.Opacity != 1 {
		t.Errorf("Attendu Opacity 1 après Reset, obtenu %v", p.Opacity)
	}
}

// TestIsDead vérifie les conditions de mort de la particule
func TestIsDead(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.Lifetime = 5
	cfg.WindowSizeX = 100
	cfg.WindowSizeY = 100
	cfg.Margin = 10
	testutil.UseConfig(cfg)

	p := Particle{Age: 5, PosX: 50, PosY: 50}
	if !p.IsDead() {
		t.Errorf("Attendu particule morte par Age")
	}

	p = Particle{Age: 0, PosX: -11, PosY: 50}
	if !p.IsDead() {
		t.Errorf("Attendu particule morte par PosX")
	}

	p = Particle{Age: 0, PosX: 50, PosY: -11}
	if !p.IsDead() {
		t.Errorf("Attendu particule morte par PosY")
	}

	p = Particle{Age: 0, PosX: 50, PosY: 50}
	if p.IsDead() {
		t.Errorf("Attendu particule en vie")
	}
}
