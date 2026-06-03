package collision

import (
	"container/list"
	"math"
	"project-particles/config"
	"project-particles/particle"
)

// Init initialise une nouvelle grille de collision vide.
func Init() CollisionGrid {
	return CollisionGrid{
		Grid:     make(map[Coord][]*particle.Particle),
		resolved: make(map[Pair]bool),
	}
}

// getSummits retourne les positions des sommets de la particule p
// en prenant en compte la rotation.
func getSummits(p *particle.Particle) [4][2]float64 {
	halfW := p.Width / 2
	halfH := p.Height / 2

	cx := p.PosX + halfW
	cy := p.PosY + halfH

	relToCenters := [4][2]float64{
		{-halfW, -halfH},
		{halfW, -halfH},
		{halfW, halfH},
		{-halfW, halfH},
	}

	var summitsPos [4][2]float64
	cos := math.Cos(p.Rotation)
	sin := math.Sin(p.Rotation)

	for i, summit := range relToCenters {
		summitsPos[i][0] = cx + summit[0]*cos - summit[1]*sin
		summitsPos[i][1] = cy + summit[1]*cos + summit[0]*sin
	}

	return summitsPos
}

// getNormalAxes retourne les normales des côtés d'un rectangle défini par les sommets donnés.
//
// Les côtés des rectangles étant parallèle, ils possèdent la même normale et ne retourne que deux normales.
func getNormalAxes(summits [4][2]float64) [2][2]float64 {

	n1x := -(summits[0][1] - summits[1][1])
	n1y := summits[0][0] - summits[1][0]
	n1len := math.Sqrt(n1x*n1x + n1y*n1y)

	n1x /= n1len
	n1y /= n1len

	n2x := -(summits[1][1] - summits[2][1])
	n2y := summits[1][0] - summits[2][0]
	n2len := math.Sqrt(n2x*n2x + n2y*n2y)

	n2x /= n2len
	n2y /= n2len

	return [2][2]float64{
		{n1x, n1y},
		{n2x, n2y},
	}
}

// projectSummits calcule la projection d'un rectangle sur un axe donné.
func projectSummits(summits [4][2]float64, axe [2]float64) (float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64

	for _, summit := range summits {
		dot := summit[0]*axe[0] + summit[1]*axe[1]
		if dot < min {
			min = dot
		}
		if dot > max {
			max = dot
		}
	}

	return min, max
}

// checkCollision vérifie si deux particules se chevauchent selon la méthode SAT.
//
// S'il y a une collision, la normale de celle-ci et son chevauchement est renvoyé.
func checkCollision(p1, p2 *particle.Particle) (isColliding bool, minOverlap, nx, ny float64) {
	sum1 := getSummits(p1)
	sum2 := getSummits(p2)

	axes1 := getNormalAxes(sum1)
	axes2 := getNormalAxes(sum2)

	minOverlap = math.MaxFloat64

	for _, axes := range [2][2][2]float64{axes1, axes2} {
		for _, axe := range axes {
			min1, max1 := projectSummits(sum1, axe)
			min2, max2 := projectSummits(sum2, axe)

			if max1 < min2 || max2 < min1 {
				return false, 0, 0, 0
			}

			overlap := math.Min(max1, max2) - math.Max(min1, min2)
			if overlap < minOverlap {
				minOverlap = overlap
				nx = axe[0]
				ny = axe[1]
			}
		}
	}

	dx := (p2.PosX + p2.Width/2) - (p1.PosX + p1.Width/2)
	dy := (p2.PosY + p2.Height/2) - (p1.PosY + p1.Height/2)

	if dx*nx+dy*ny < 0 {
		nx = -nx
		ny = -ny
	}
	return true, minOverlap, nx, ny
}

// correctOverlap corrige l'overlap en fonction de la masse.
func correctOverlap(p1, p2 *particle.Particle, nx, ny, overlap float64) {
	m1 := p1.Width * p1.Height
	m2 := p2.Width * p2.Height
	totalMass := m1 + m2

	p1.PosX -= nx * overlap * (m2 / totalMass)
	p1.PosY -= ny * overlap * (m2 / totalMass)

	p2.PosX += nx * overlap * (m1 / totalMass)
	p2.PosY += ny * overlap * (m1 / totalMass)
}

// pointSegmentDistance calcule la distance qui sépare le point (px, py) du segment [(ax, ay);(bx, by)].
func pointSegmentDistance(px, py, ax, ay, bx, by float64) float64 {
	abx := bx - ax
	aby := by - ay
	apx := px - ax
	apy := py - ay

	abLenSq := abx*abx + aby*aby
	if abLenSq == 0 {
		return math.Sqrt(apx*apx + apy*apy)
	}

	t := (apx*abx + apy*aby) / abLenSq
	t = math.Max(0, math.Min(1, t))

	cx := ax + t*abx
	cy := ay + t*aby

	dx := px - cx
	dy := py - cy
	return math.Sqrt(dx*dx + dy*dy)
}

// closestPoint cherche le(s) point(s) de contact les plus proches entre deux rectangles.
func closestPoint(points, rect [4][2]float64, minDist *float64, contacts *[2][2]float64, contactCount *int) {
	verySmallNumber := 0.0005

	for _, p := range points {
		for i := 0; i < 4; i++ {
			a := rect[i]
			b := rect[(i+1)%4]

			dist := pointSegmentDistance(
				p[0], p[1],
				a[0], a[1],
				b[0], b[1],
			)

			if math.Abs(dist-*minDist) < verySmallNumber {
				dx := p[0] - contacts[0][0]
				dy := p[1] - contacts[0][1]
				if *contactCount == 1 && dx*dx+dy*dy > verySmallNumber*verySmallNumber {
					contacts[1] = p
					*contactCount = 2
				}
			} else if dist < *minDist {
				*minDist = dist
				contacts[0] = p
				*contactCount = 1
			}
		}
	}
}

// getContactPoint retourne les positions des points de contact et le nombre de contact.
func getContactPoint(p1, p2 *particle.Particle) (contacts [2][2]float64, contactCount int) {

	minDist := math.MaxFloat64

	s1 := getSummits(p1)
	s2 := getSummits(p2)

	closestPoint(s1, s2, &minDist, &contacts, &contactCount)
	closestPoint(s2, s1, &minDist, &contacts, &contactCount)

	return contacts, contactCount
}

// applyImpulse applique l'impulsion pour résoudre la collision entre deux particules.
//
// Cette méthode ajuste les vitesses des particules selon la restitution élastique
// et la masse (aire des particules). Si les particules s'éloignent,
// aucune impulsion n'est appliquée.
func applyImpulse(p1, p2 *particle.Particle, nx, ny float64) {
	contacts, count := getContactPoint(p1, p2)

	for i := 0; i < count; i++ {
		r1x := contacts[i][0] - (p1.PosX + p1.Width*0.5)
		r1y := contacts[i][1] - (p1.PosY + p1.Height*0.5)
		r2x := contacts[i][0] - (p2.PosX + p2.Width*0.5)
		r2y := contacts[i][1] - (p2.PosY + p2.Height*0.5)

		v1x := p1.VelX - p1.AngVel*r1y
		v1y := p1.VelY + p1.AngVel*r1x
		v2x := p2.VelX - p2.AngVel*r2y
		v2y := p2.VelY + p2.AngVel*r2x

		vRelx := v2x - v1x
		vRely := v2y - v1y
		vn := vRelx*nx + vRely*ny

		if vn >= 0 {
			continue
		}

		m1 := p1.Width * p1.Height
		m2 := p2.Width * p2.Height

		I1 := (m1 / 12.0) * (p1.Width*p1.Width + p1.Height*p1.Height)
		I2 := (m2 / 12.0) * (p2.Width*p2.Width + p2.Height*p2.Height)

		r1n := r1x*ny - r1y*nx
		r2n := r2x*ny - r2y*nx

		e := 1.0 // collision élastique
		j := -(1 + e) * vn / ((1 / m1) + (1 / m2) + ((r1n * r1n) / I1) + ((r2n * r2n) / I2))
		j /= float64(count)

		impX := j * nx
		impY := j * ny

		p1.VelX -= impX / m1
		p1.VelY -= impY / m1
		p2.VelX += impX / m2
		p2.VelY += impY / m2

		p1.AngVel -= (j * r1n) / I1
		p2.AngVel += (j * r2n) / I2
	}

}

// Build construit la grille de collision pour toutes les particules.
//
// Les cellules vides sont nettoyées.
func (cg *CollisionGrid) Build(particles *list.List) {

	cfg := config.General
	cellSize := cfg.CellSize

	for k, cell := range cg.Grid {
		if len(cell) == 0 {
			delete(cg.Grid, k)
		} else {
			cg.Grid[k] = cell[:0]
		}
	}

	e := particles.Front()
	for e != nil {
		next := e.Next()

		p, ok := e.Value.(*particle.Particle)
		if !ok {
			e = next
			continue
		}

		if p.IsDead() {
			break
		}

		summits := getSummits(p)

		minX, maxX := summits[0][0], summits[0][0]
		minY, maxY := summits[0][1], summits[0][1]

		for _, s := range summits {
			if s[0] < minX {
				minX = s[0]
			}
			if s[0] > maxX {
				maxX = s[0]
			}
			if s[1] < minY {
				minY = s[1]
			}
			if s[1] > maxY {
				maxY = s[1]
			}
		}

		minI := int(math.Floor(minX / float64(cellSize)))
		maxI := int(math.Floor(maxX / float64(cellSize)))
		minJ := int(math.Floor(minY / float64(cellSize)))
		maxJ := int(math.Floor(maxY / float64(cellSize)))

		for i := minI; i <= maxI; i++ {
			for j := minJ; j <= maxJ; j++ {
				cg.Grid[Coord{i, j}] = append(cg.Grid[Coord{i, j}], p)
			}
		}

		e = next

	}
}

// isResolved vérifie si une paire de particules a déjà été traitée.
func (cg *CollisionGrid) isResolved(p1, p2 *particle.Particle) bool {
	idA, idB := p1.ID, p2.ID
	if idA > idB {
		idA, idB = idB, idA
	}

	pair := Pair{idA, idB}
	_, ok := cg.resolved[pair]
	if !ok {
		cg.resolved[pair] = true
	}

	return ok
}

// ResolveCollisions parcourt la grille et applique la résolution des collisions.
func (cg *CollisionGrid) ResolveCollisions() {
	cg.resolved = make(map[Pair]bool)

	for _, particles := range cg.Grid {
		if len(particles) < 2 {
			continue
		}

		for i, p1 := range particles {
			for j := i + 1; j < len(particles); j++ {
				p2 := particles[j]

				if cg.isResolved(p1, p2) {
					continue
				}

				isColliding, overlap, nx, ny := checkCollision(p1, p2)
				if !isColliding {
					continue
				}

				correctOverlap(p1, p2, nx, ny, overlap)
				applyImpulse(p1, p2, nx, ny)
			}
		}
	}
}
