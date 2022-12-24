package main

type ResourceType int

const (
	Ore ResourceType = iota
	Clay
	Obsidian
	Geode
)

type RobotType int

const (
	OreRobot RobotType = iota
	ClayRobot
	ObsidianRobot
	GeodeRobot
)

func (r RobotType) ResourceType() ResourceType {
	switch r {
	case OreRobot:
		return Ore
	case ClayRobot:
		return Clay
	case ObsidianRobot:
		return Obsidian
	default:
		return Geode
	}
}

var AllRobotTypes = []RobotType{GeodeRobot, ObsidianRobot, ClayRobot, OreRobot}

type ResourceBundle map[ResourceType]int

func (r ResourceBundle) Clone() ResourceBundle {
	res := make(ResourceBundle)
	for rt, n := range r {
		res[rt] = n
	}
	return res
}

func (r ResourceBundle) Get(resourceType ResourceType) int {
	n, ok := r[resourceType]
	if !ok {
		return 0
	}
	return n
}

type RobotPark map[RobotType]int

func (r RobotPark) Clone() RobotPark {
	res := make(RobotPark)
	for rt, n := range r {
		res[rt] = n
	}
	return res
}
