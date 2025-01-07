package game

import (
	"errors"
	"time"
)

var ErrAnimalStarved = errors.New("An Animal just died because it did not have enough food. :(")

type AnimalKind string

const (
	Herbivore AnimalKind = "herbivore"
	Carnivore AnimalKind = "carnivore"
)

type Animal interface {
	Entity
	GetKind() AnimalKind
	SetKind(kind AnimalKind)
}

func NewWolf(strength float64) *Wolf {
	return &Wolf{
		Kind:     Carnivore,
		Strength: strength,
	}
}

type Wolf struct {
	ID int `json:"id"`

	Kind     AnimalKind `json:"kind"`
	Strength float64    `json:"strength"`
}

func (w *Wolf) GetID() int {
	return w.ID
}
func (w *Wolf) SetID(id int) {
	w.ID = id
}

func (w *Wolf) GetKind() AnimalKind {
	return w.Kind
}
func (w *Wolf) SetKind(kind AnimalKind) {
	w.Kind = kind
}

func (w *Wolf) Tick(g *Game, deltatime time.Duration) error {
	return nil
}
