package collision

import "project-particles/particle"

// Coord représente les coordonnées d'une cellule dans la grille de collision.
type Coord struct {
	X, Y int
}

// Pair représente une paire unique d'identifiants de particules.
type Pair struct {
	A, B int
}

// CollisionGrid est la grille utilisée pour gérer efficacement les collisions.
//
// Grid     : map des coordonnées de cellules vers les particules qui s'y trouvent.
// resolved : map des paires de particules déjà traitées lors de la résolution
type CollisionGrid struct {
	Grid     map[Coord][]*particle.Particle
	resolved map[Pair]bool
}
