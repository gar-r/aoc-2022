package main

func (s *Scan) CanEscape(v *Vector, boundingCube int) bool {
	visited := NewPartition()
	return s.canEscape(v, boundingCube, visited)
}

func (s *Scan) canEscape(v *Vector, boundingCube int, visited *Partition) bool {
	if escaped(v, boundingCube) {
		return true
	}
	if s.Contains(v) {
		return false
	}
	visited.Set(v)
	for _, dir := range Directions {
		next := v.Add(dir)
		if !visited.Contains(next) &&
			s.canEscape(v.Add(dir), boundingCube, visited) {
			return true
		}
	}
	return false
}

func escaped(v *Vector, bounds int) bool {
	return v.X > bounds || v.Y > bounds || v.Z > bounds
}
