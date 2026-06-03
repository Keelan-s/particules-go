package collision

import (
	"math"
	"project-particles/particle"
	"testing"
)

// makeParticle créé une fausse particule
func makeParticle(id int, x, y, w, h, r float64) *particle.Particle {
	return &particle.Particle{
		ID:       id,
		PosX:     x,
		PosY:     y,
		Width:    w,
		Height:   h,
		Rotation: r,
	}
}

func TestGetSummits(t *testing.T) {
	p := makeParticle(1, 0, 0, 2, 4, 0)
	summits := getSummits(p)

	expected := [4][2]float64{
		{-1, -2}, {1, -2}, {1, 2}, {-1, 2},
	}

	for i := 0; i < 4; i++ {
		if math.Abs(summits[i][0]-expected[i][0]) > 1e-6 || math.Abs(summits[i][1]-expected[i][1]) > 1e-6 {
			t.Errorf("Sommet %d incorrect. Obtenu %f, attendu %f", i, summits[i], expected[i])
		}
	}
}

func TestCheckCollision_NoCollision(t *testing.T) {
	p1 := makeParticle(1, 0, 0, 2, 2, 0)
	p2 := makeParticle(2, 5, 5, 2, 2, 0)

	colliding, _, _, _ := checkCollision(p1, p2)
	if colliding {
		t.Errorf("Les particules ne devraient pas être en collision")
	}
}

func TestCheckCollision_WithCollision(t *testing.T) {
	p1 := makeParticle(1, 0, 0, 2, 2, 0)
	p2 := makeParticle(2, 1, 1, 2, 2, 0)

	colliding, overlap, nx, ny := checkCollision(p1, p2)
	if !colliding {
		t.Errorf("Les particules devraient être en collision")
	}

	if overlap <= 0 {
		t.Errorf("L'overlap doit être positif")
	}

	if nx == 0 && ny == 0 {
		t.Errorf("La normale de collision ne doit pas être nulle")
	}
}

func TestCorrectOverlap(t *testing.T) {
	p1 := makeParticle(1, 0, 0, 2, 2, 0)
	p2 := makeParticle(2, 1, 0, 2, 2, 0)

	_, overlap, nx, ny := checkCollision(p1, p2)
	correctOverlap(p1, p2, nx, ny, overlap)

	colliding, _, _, _ := checkCollision(p1, p2)
	if colliding {
		t.Errorf("Les particules ne devraient plus se chevaucher après correction")
	}
}

func TestApplyImpulse(t *testing.T) {
	p1 := makeParticle(1, 0, 0, 2, 2, 0)
	p2 := makeParticle(2, 1, 0, 2, 2, 0)

	p1.VelX = 1
	p1.VelY = 0
	p2.VelX = -1
	p2.VelY = 0

	applyImpulse(p1, p2, 1, 0)

	if p1.VelX <= 0 {
		t.Errorf("La vitesse de p1 devrait changer positivement après collision")
	}

	if p2.VelX >= 0 {
		t.Errorf("La vitesse de p2 devrait changer négativement après collision")
	}
}
