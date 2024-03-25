package topInterview150

func GasStation(gas []int, cost []int) int {
	var fuelLeft, globalFuelLeft, start int
	for i := 0; i < len(gas); i++ {
		globalFuelLeft += gas[i] - cost[i]
		fuelLeft += gas[i] - cost[i]
		if fuelLeft < 0 {
			start = i + 1
			fuelLeft = 0
		}
	}

	if globalFuelLeft < 0 {
		return -1
	}
	return start
}
