package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           int
	SpawnRate                float64
	RandomSize               bool
	SizeX, SizeY             float64
	MinSizeX, MinSizeY       float64
	MaxSizeX, MaxSizeY       float64
	RandomColor              bool
	ColorR, ColorG, ColorB   float64
	MinVelX, MinVelY         float64
	MaxVelX, MaxVelY         float64
	MinAngVel, MaxAngVel     float64
	Gravity                  float64
	Collision                bool
	CellSize                 int
	Lifetime                 int
	Margin                   float64
}

var General Config
