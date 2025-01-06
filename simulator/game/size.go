package game

// Creates a size of s in meters
func NewSize(s float64) Size {
	return Size(s)
}

type Size float64

func (s Size) Meters() float64 {
	return float64(s)
}
func (s Size) Decimeters() float64 {
	return float64(s) / 10
}
func (s Size) Centimeters() float64 {
	return float64(s) / 100
}
func (s Size) Millimeters() float64 {
	return float64(s) / 1000
}
