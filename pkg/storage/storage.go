// Package storage is the persistence layer for drinks
package storage

import (
	"fmt"

	"github.com/google/uuid"
)

type Drinks struct {
	Drinks map[string]Drink
}

type Drink struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type IstorageHandler interface {
	GetDrinks() Drinks
	GetDrink(id string) (drink Drink, err error)
	SaveDrink(drink Drink) Drink
	DeleteDrink(id string) (err error)
}

func GetInstance() IstorageHandler {
	return &Drinks{Drinks: make(map[string]Drink)}
}

func (d *Drinks) GetDrinks() Drinks {
	return *d
}

func (d *Drinks) GetDrink(id string) (drink Drink, err error) {
	if d.checkKey(id) {
		drink = d.Drinks[id]
	} else {
		err = fmt.Errorf("could not find drink with id = %s", id)
	}
	return
}

func (d *Drinks) SaveDrink(drink Drink) Drink {
	drink.ID = uuid.New().String() // new id
	d.Drinks[drink.ID] = drink
	return drink
}

func (d *Drinks) DeleteDrink(id string) (err error) {
	if d.checkKey(id) {
		delete(d.Drinks, id)
	} else {
		err = fmt.Errorf("could not find drink with id = %s", id)
	}
	return
}

func (d *Drinks) checkKey(id string) bool {
	_, ok := d.Drinks[id]
	return ok
}
