package game

// Creates a new water measurement (in liters)
func NewWater(water float64) Water {
	return Water(water)
}

// Water amount measured in liters
type Water float64
