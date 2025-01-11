package mathutil

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v Vector2D) Abs() Vector2D {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
	return v
}

func (v Vector2D) Add(val Vector2D) Vector2D {
	v.X += val.X
	v.Y += val.Y
	return v
}

func (v Vector2D) Sub(val Vector2D) Vector2D {
	v.X -= val.X
	v.Y -= val.Y
	return v
}

func (v Vector2D) DistanceTo(to Vector2D) float64 {
	return math.Sqrt(math.Pow(v.X-to.X, 2) + math.Pow(v.Y-to.Y, 2))
}

func (v *Vector2D) Scale(val float64) {
	v.X *= val
	v.Y *= val

}

func (v Vector2D) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vector2D) Normalize() {
	len := v.Length()
	v.X /= len
	v.Y /= len
}
