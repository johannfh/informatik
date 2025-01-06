package game

import "errors"

var ErrPlantDriedUp = errors.New("A Plant just died because it did not have enough water. :(")

type Plant interface {
	Entity
	GetHeight() float64
	SetHeight(height float64)
}
