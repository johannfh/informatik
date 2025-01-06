package game

import "errors"

var ErrAnimalStarved = errors.New("An Animal just died because it did not have enough food. :(")

type Animal interface {
	Entity
}
