package particles

import (
	"container/list"
)

// System définit un système de particules.
//
// Content : liste contenant toutes les particules.
// ParticleDead : élément utilisé pour marquer la première particule morte.
type System struct {
	Content      *list.List
	ParticleDead *list.Element
}
