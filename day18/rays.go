package main

type Tracer struct {
	Flood   bool
	scan    *Scan
	raySet  map[*Vector][]*Vector
	cubeLen int
}

func NewTracer(scan *Scan, cubeLen int) *Tracer {
	return &Tracer{
		scan:    scan,
		raySet:  createRaySet(cubeLen),
		cubeLen: cubeLen,
	}
}

func createRaySet(cubeLen int) map[*Vector][]*Vector {
	rays := make(map[*Vector][]*Vector)
	for i := -cubeLen; i < cubeLen; i++ {
		for j := -cubeLen; j < cubeLen; j++ {
			rays[X] = append(rays[X], &Vector{-cubeLen, i, j})
			rays[Y] = append(rays[Y], &Vector{i, -cubeLen, j})
			rays[Z] = append(rays[Z], &Vector{i, j, -cubeLen})
		}
	}
	return rays
}

func (t *Tracer) TraceSurfaces() int {
	surfaces := 0
	for dir, rays := range t.raySet {
		surfaces += t.traceRay(dir, rays)
	}
	return surfaces
}

func (t *Tracer) traceRay(dir *Vector, rays []*Vector) int {
	surfaces := 0
	inside := false
	for _, ray := range rays {
		for i := 0; i <= 2*t.cubeLen; i++ {
			offset := dir.Scale(i)
			v := ray.Add(offset)
			if !inside && t.scan.Contains(v) {
				surfaces++
				inside = true
			} else if inside && !t.scan.Contains(v) {
				if t.Flood && !t.scan.CanEscape(v, t.cubeLen) {
					continue
				}
				surfaces++
				inside = false
			}
		}
	}
	return surfaces
}
