package main

type Blueprint map[RobotType]ResourceBundle

func (b Blueprint) CalcGeodes(resources ResourceBundle, robots RobotPark, timeLimit int) int {
	return b.calcGeodes(resources, robots, timeLimit)
}

func (b Blueprint) calcGeodes(resources ResourceBundle, robots RobotPark, timeLimit int) int {
	if timeLimit == 0 {
		resources.Get(Geode)
	}
	units := b.buildableUnits(resources)
	resourcesAfter := b.calcYields(resources, robots)
	res := 0
	for _, unit := range units {
		robotsAfter := robots.Clone()
		robotsAfter[unit]++
		res = max(res, b.calcGeodes(resourcesAfter, robotsAfter, timeLimit-1))
	}
	return max(res, b.calcGeodes(resourcesAfter, robots, timeLimit-1))
}

func (b Blueprint) buildableUnits(resources ResourceBundle) []RobotType {
	res := make([]RobotType, 0)
	for _, robotType := range AllRobotTypes {
		if b.canBuild(robotType, resources) {
			res = append(res, robotType)
		}
	}
	return res
}

func (b Blueprint) canBuild(robotType RobotType, resources ResourceBundle) bool {
	requiredResources := b[robotType]
	for resourceType, amount := range requiredResources {
		if resources.Get(resourceType) < amount {
			return false
		}
	}
	return true
}

func (b Blueprint) calcYields(assets ResourceBundle, robots RobotPark) ResourceBundle {
	a := assets.Clone()
	for rt, n := range robots {
		a[rt.ResourceType()] += n
	}
	return a
}

func max(i, j int) int {
	if j > i {
		return j
	}
	return i
}
