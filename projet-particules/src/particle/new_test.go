package particle

import (
	"project-particles/config"
	"project-particles/testutil"
	"testing"
)

// TestInitSpawn vérifie que la position est correctement initialisée
func TestInitSpawn(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.RandomSpawn = false
	cfg.SpawnX = 100
	cfg.SpawnY = 200
	testutil.UseConfig(cfg)

	p := &Particle{}
	initSpawn(p, cfg)

	if p.PosX != 100 || p.PosY != 200 {
		t.Errorf("Attendu pos (100,200), obtenu (%v,%v)", p.PosX, p.PosY)
	}
}

// TestInitSize vérifie que la taille est correctement initialisée
func TestInitSize(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.RandomSize = false
	cfg.SizeX = 3
	cfg.SizeY = 4
	testutil.UseConfig(cfg)

	p := &Particle{}
	initSize(p, cfg)

	if p.ScaleX != 3 || p.ScaleY != 4 {
		t.Errorf("Attendu scale (3,4), obtenu (%v,%v)", p.ScaleX, p.ScaleY)
	}
}

// TestInitVelocity vérifie que la vélocité est dans les bornes attendues
func TestInitVelocity(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.MinVelX = 1
	cfg.MaxVelX = 2
	cfg.MinVelY = 3
	cfg.MaxVelY = 4
	testutil.UseConfig(cfg)

	p := &Particle{}
	initVelocity(p, cfg)

	if p.VelX < -2 || p.VelX > 2 {
		t.Errorf("VelX out of range: %v", p.VelX)
	}
	if p.VelY < -4 || p.VelY > 4 {
		t.Errorf("VelY out of range: %v", p.VelY)
	}
}

// TestInitColor vérifie la couleur selon RandomColor
func TestInitColor(t *testing.T) {
	orig := config.General
	defer testutil.UseConfig(orig)

	cfg := testutil.NewFakeConfig()
	cfg.RandomColor = false
	cfg.ColorR = 128
	cfg.ColorG = 64
	cfg.ColorB = 32
	testutil.UseConfig(cfg)

	p := &Particle{}
	initColor(p, cfg)

	if p.ColorR != 128.0/255 || p.ColorG != 64.0/255 || p.ColorB != 32.0/255 {
		t.Errorf("unAttendu color (%v,%v,%v)", p.ColorR, p.ColorG, p.ColorB)
	}
}
