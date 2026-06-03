package ui

import (
	"math"
	"project-particles/config"
)

// Init construit et retourne l'interface utilisateur de configuration.
//
// L'UI est directement liée à config.General : chaque widget modifie
// dynamiquement les paramètres du système via des closures.
func Init() UI {
	cfg := &config.General

	mainFrame := newMainFrame(cfg)
	cursor := newCursor(&mainFrame)

	widgets := buildWidgets(cursor, cfg)
	mainFrame.Widgets = widgets

	return mainFrame
}

// newMainFrame crée et initialise le conteneur principal de l'interface utilisateur.
func newMainFrame(cfg *config.Config) UI {
	return UI{
		X:           cfg.WindowSizeX / 10,
		Y:           cfg.WindowSizeY / 10,
		Visible:     false,
		ColumnWidth: (cfg.WindowSizeX / 30) * 8,
		NumColumn:   3,
		LineHeight:  cfg.WindowSizeY / 25,
		NumLine:     1,
		Margin:      25,
	}
}

// newCursor initialise un curseur de layout associé à un conteneur UI.
//
// Le curseur est utilisé pour positionner séquentiellement les widgets
// dans l'interface en respectant les lignes et colonnes du layout.
func newCursor(frame *UI) LayoutCursor {
	return LayoutCursor{
		X:      frame.X,
		Y:      frame.Y + frame.LineHeight,
		Parent: frame,
	}
}

// buildWidgets construit l'ensemble des widgets de l'interface utilisateur.
//
// Chaque widget est lié dynamiquement à config.General via des closures,
// permettant une modification en temps réel des paramètres du système.
// La fonction gère également le positionnement des widgets à l'aide
// du LayoutCursor fourni.
func buildWidgets(cursor LayoutCursor, cfg *config.Config) []Widget {
	var widgets []Widget

	widgets = append(widgets,
		NewToggle("Debug", cursor, &cfg.Debug),
	)
	cursor.NextLine()

	widgets = append(widgets,
		NewFloatField(
			"SpawnRate",
			cursor,
			func() float64 { return cfg.SpawnRate },
			func(v float64) { cfg.SpawnRate = v },
		),
	)
	cursor.NextLine()

	widgets = append(widgets,
		NewToggle("RandomSpawn", cursor, &cfg.RandomSpawn),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewIntField(
			"SpawnX",
			cursor,
			func() int { return cfg.SpawnX },
			func(v int) { cfg.SpawnX = v },
		),
	)

	cursor.SetColumn(2)
	widgets = append(widgets,
		NewIntField(
			"SpawnY",
			cursor,
			func() int { return cfg.SpawnY },
			func(v int) { cfg.SpawnY = v },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewToggle("RandomSize", cursor, &cfg.RandomSize),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewFloatField(
			"SizeX",
			cursor,
			func() float64 { return cfg.SizeX },
			func(v float64) { cfg.SizeX = v },
		),
	)

	cursor.SetColumn(2)
	widgets = append(widgets,
		NewFloatField(
			"SizeY",
			cursor,
			func() float64 { return cfg.SizeY },
			func(v float64) { cfg.SizeY = v },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MinSizeX",
			cursor,
			func() float64 { return cfg.MinSizeX },
			func(v float64) { cfg.MinSizeX = v },
			func(v float64) float64 { return clamp(v, 0, cfg.MaxSizeX) },
		),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MaxSizeX",
			cursor,
			func() float64 { return cfg.MaxSizeX },
			func(v float64) { cfg.MaxSizeX = v },
			func(v float64) float64 { return clamp(v, cfg.MinSizeX, math.MaxFloat64) },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MinSizeY",
			cursor,
			func() float64 { return cfg.MinSizeY },
			func(v float64) { cfg.MinSizeY = v },
			func(v float64) float64 { return clamp(v, 0, cfg.MaxSizeY) },
		),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MaxSizeY",
			cursor,
			func() float64 { return cfg.MaxSizeY },
			func(v float64) { cfg.MaxSizeY = v },
			func(v float64) float64 { return clamp(v, cfg.MinSizeY, math.MaxFloat64) },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewToggle("RandomColor", cursor, &cfg.RandomColor),
	)
	cursor.NextLine()

	widgets = append(widgets,
		NewFloatFieldConstraint(
			"ColorR",
			cursor,
			func() float64 { return cfg.ColorR },
			func(v float64) { cfg.ColorR = v },
			colorConstraint,
		),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"ColorG",
			cursor,
			func() float64 { return cfg.ColorG },
			func(v float64) { cfg.ColorG = v },
			colorConstraint,
		),
	)

	cursor.SetColumn(2)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"ColorB",
			cursor,
			func() float64 { return cfg.ColorB },
			func(v float64) { cfg.ColorB = v },
			colorConstraint,
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MinVelX",
			cursor,
			func() float64 { return cfg.MinVelX },
			func(v float64) { cfg.MinVelX = v },
			func(v float64) float64 { return clamp(v, 0, cfg.MaxVelX) },
		),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MaxVelX",
			cursor,
			func() float64 { return cfg.MaxVelX },
			func(v float64) { cfg.MaxVelX = v },
			func(v float64) float64 { return clamp(v, cfg.MinVelX, math.MaxFloat64) },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MinVelY",
			cursor,
			func() float64 { return cfg.MinVelY },
			func(v float64) { cfg.MinVelY = v },
			func(v float64) float64 { return clamp(v, 0, cfg.MaxVelY) },
		),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MaxVelY",
			cursor,
			func() float64 { return cfg.MaxVelY },
			func(v float64) { cfg.MaxVelY = v },
			func(v float64) float64 { return clamp(v, cfg.MinVelY, math.MaxFloat64) },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MinAngVel",
			cursor,
			func() float64 { return cfg.MinAngVel },
			func(v float64) { cfg.MinAngVel = v },
			func(v float64) float64 { return clamp(v, 0, cfg.MaxAngVel) },
		),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewFloatFieldConstraint(
			"MaxAngVel",
			cursor,
			func() float64 { return cfg.MaxAngVel },
			func(v float64) { cfg.MaxAngVel = v },
			func(v float64) float64 { return clamp(v, cfg.MinAngVel, math.MaxFloat64) },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewFloatField(
			"Gravity",
			cursor,
			func() float64 { return cfg.Gravity },
			func(v float64) { cfg.Gravity = v },
		),
	)
	cursor.NextLine()

	widgets = append(widgets,
		NewToggle("Collision", cursor, &cfg.Collision),
	)

	cursor.SetColumn(1)
	widgets = append(widgets,
		NewIntField(
			"CellSize",
			cursor,
			func() int { return cfg.CellSize },
			func(v int) { cfg.CellSize = v },
		),
	)
	cursor.NextLine()

	cursor.SetColumn(0)
	widgets = append(widgets,
		NewIntField(
			"Lifetime",
			cursor,
			func() int { return cfg.Lifetime },
			func(v int) { cfg.Lifetime = v },
		),
	)
	cursor.NextLine()

	widgets = append(widgets,
		NewFloatField(
			"Margin",
			cursor,
			func() float64 { return cfg.Margin },
			func(v float64) { cfg.Margin = v },
		),
	)
	cursor.NextLine() // Laisse une ligne vide

	return widgets
}
