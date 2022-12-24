package main

import "fmt"

func main() {
	blueprints := example
	assets := ResourceBundle{}
	robots := RobotPark{OreRobot: 1}
	geodes := blueprints[0].CalcGeodes(assets, robots, 16)
	fmt.Printf("geodes: %d\n", geodes)
}
