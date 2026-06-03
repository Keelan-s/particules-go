package particle

// Particle définit une particule.
//
//Elle possède :
//	- une position,
//	- une vitesse
//	- une taille,
//	- une orientation,
//	- une couleur,
//	- une opacité,
//	- une largeur physique,
//	- un age,
//	- un identifiant.
//
type Particle struct {
	PosX, PosY             float64
	VelX, VelY             float64
	ScaleX, ScaleY         float64
	Rotation               float64
	AngVel                 float64
	ColorR, ColorG, ColorB float64
	Opacity                float64
	Width, Height          float64
	Age                    int
	ID                     int
}
