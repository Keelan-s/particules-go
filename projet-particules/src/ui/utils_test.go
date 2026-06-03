package ui

import "testing"

// TestClamp vérifie le comportement de clamp sur différentes valeurs
func TestClamp(t *testing.T) {
	tests := []struct {
		v, min, max, want float64
	}{
		{5, 0, 10, 5},   // normal
		{-1, 0, 10, 0},  // au dessous
		{15, 0, 10, 10}, // au dessus
		{0, 0, 10, 0},   // égal à min
		{10, 0, 10, 10}, // égal à max
	}

	for _, test := range tests {
		got := clamp(test.v, test.min, test.max)
		if got != test.want {
			t.Errorf("Atestendu %f, obtenu %f", test.want, got)
		}
	}
}

// TestColorConstraint vérifie que colorConstraint reste entre 0 et 255.
func TestColorConstraint(t *testing.T) {
	tests := []struct {
		v, want float64
	}{
		{-10, 0},
		{0, 0},
		{128, 128},
		{255, 255},
		{300, 255},
	}

	for _, test := range tests {
		got := colorConstraint(test.v)
		if got != test.want {
			t.Errorf("Attendu %f, obtenu %f", test.want, got)
		}
	}
}

// TestPointInRect vérifie les positions à l'intérieur et à l'extérieur d'un rectangle.
func TestPointInRect(t *testing.T) {
	tests := []struct {
		px, py, x, y, w, h int
		want               bool
	}{
		{5, 5, 0, 0, 10, 10, true},   // intérieur
		{0, 0, 0, 0, 10, 10, true},   // coin à gauche
		{10, 10, 0, 0, 10, 10, true}, // coin à droite
		{-1, 5, 0, 0, 10, 10, false}, // à gauche
		{5, -1, 0, 0, 10, 10, false}, // en haut
		{11, 5, 0, 0, 10, 10, false}, // à droite
		{5, 11, 0, 0, 10, 10, false}, // en bas
	}

	for _, test := range tests {
		got := pointInRect(test.px, test.py, test.x, test.y, test.w, test.h)
		if got != test.want {
			t.Errorf("Attendu %t, obtenu %t", test.want, got)
		}
	}
}
