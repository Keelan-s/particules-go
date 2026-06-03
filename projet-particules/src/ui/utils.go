package ui

import "strconv"

// NextLine déplace le curseur à la ligne suivante dans le layout.
func (c *LayoutCursor) NextLine() {
	c.Y += c.Parent.LineHeight
	c.Parent.NumLine++
}

// SetColumn positionne le curseur dans la colonne spécifiée.
//
// Les colonnes commencent à l'indice 0 et n est l'indice de la colonne.
func (c *LayoutCursor) SetColumn(n int) {
	c.X = c.Parent.X + n*c.Parent.ColumnWidth
}

// NewToggle crée et initialise un Toggle.
//
// label : texte affiché à côté du Toggle
// c     : curseur de layout déterminant la position initiale
// v     : pointeur vers la valeur booléenne associée au Toggle
func NewToggle(label string, c LayoutCursor, v *bool) *Toggle {

	return &Toggle{
		Label: label,
		X:     c.X,
		Y:     c.Y,
		Value: v,
	}
}

// NewFloatField crée un TextField pour modifier une valeur float64.
//
// get : fonction pour lire la valeur actuelle
// set : fonction pour appliquer la nouvelle valeur
func NewFloatField(label string, c LayoutCursor,
	get func() float64,
	set func(float64)) *TextField {

	return &TextField{
		Label:  label,
		X:      c.X,
		Y:      c.Y,
		Buffer: strconv.FormatFloat(get(), 'f', 3, 64),
		Sync: func() string {
			return strconv.FormatFloat(get(), 'f', 3, 64)
		},
		Apply: func(s string) {
			v, err := strconv.ParseFloat(s, 64)
			if err == nil {
				set(v)
			}
		},
	}
}

// NewIntField crée un TextField pour modifier une valeur entière.
//
// Les paramètres suivent la même logique que ceux de NewFloatField()
func NewIntField(label string, c LayoutCursor,
	get func() int,
	set func(int)) *TextField {

	return &TextField{
		Label:  label,
		X:      c.X,
		Y:      c.Y,
		Buffer: strconv.Itoa(get()),
		Sync: func() string {
			return strconv.Itoa(get())
		},
		Apply: func(s string) {
			v, err := strconv.Atoi(s)
			if err == nil {
				set(v)
			}
		},
	}
}

// NewFloatFieldConstraint créé un TextField pour un float64 avec contrainte.
//
// constraint : fonction qui est appliquée avant d'appliquer la valeur.
// Cela permet de limiter la valeur (ex: intervalle entre min et max).
func NewFloatFieldConstraint(label string, c LayoutCursor,
	get func() float64,
	set func(float64),
	constraint func(v float64) float64) *TextField {

	return &TextField{
		Label:  label,
		X:      c.X,
		Y:      c.Y,
		Buffer: strconv.FormatFloat(get(), 'f', 2, 64),
		Sync: func() string {
			return strconv.FormatFloat(get(), 'f', 2, 64)
		},
		Apply: func(s string) {
			v, err := strconv.ParseFloat(s, 64)
			if err == nil {
				set(constraint(v))
			}
		},
	}
}

// clamp contraint une valeur entre min et max inclus.
func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// colorConstraint contraint une valeur de couleur entre 0 et 255.
func colorConstraint(v float64) float64 {
	return clamp(v, 0, 255)
}

// pointInRect vérifie si un point (px, py) est à l'intérieur d'un rectangle.
//
// Le rectangle est défini par (x, y, w, h) : position et dimensions.
func pointInRect(px, py, x, y, w, h int) bool {
	return px >= x && px <= x+w && py >= y && py <= y+h
}
