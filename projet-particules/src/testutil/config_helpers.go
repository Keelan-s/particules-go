package testutil

import "project-particles/config"

// NewFakeConfig retourne une fausse configuration utilisée pour les tests.
func NewFakeConfig() config.Config {
	return config.Config{
		WindowSizeX:      800,
		WindowSizeY:      600,
		Debug:            true,
		InitNumParticles: 1,
		SpawnRate:        1,
		RandomSpawn:      true,
		SpawnX:           400,
		SpawnY:           300,
		RandomSize:       false,
		SizeX:            1,
		SizeY:            1,
		MinSizeX:         1,
		MinSizeY:         1,
		MaxSizeX:         1,
		MaxSizeY:         1,
		RandomColor:      true,
		ColorR:           1,
		ColorG:           1,
		ColorB:           1,
		MinVelX:          0.1,
		MinVelY:          0.1,
		MaxVelX:          1.5,
		MaxVelY:          1.5,
		Gravity:          0,
		Collision:        true,
		CellSize:         32,
		Lifetime:         500,
		Margin:           50,
	}
}

// UseConfig remplace la configuration globale active par celle fournie.
// Utilisé en test pour injecter un config contrôlé.
func UseConfig(c config.Config) {
	config.General = c
}
