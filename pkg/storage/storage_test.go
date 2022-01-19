package storage_test

import (
	"azure-function-custom-handler-with-golang/pkg/storage"
	"fmt"
	"testing"
)

var (
	id    = "test"
	drink = storage.Drink{
		Name:  "Water",
		Price: 1,
	}
)

func TestGetDrinks(t *testing.T) {
	d := storage.GetInstance()
	s := d.SaveDrink(drink) // Insert one record
	res := d.GetDrinks()
	if len(res.Drinks) != 1 {
		t.Error("The storage should have one record")
	}
	if res.Drinks[s.ID].Name != drink.Name ||
		res.Drinks[s.ID].Price != drink.Price {
		t.Errorf("Name - got %s want %s\nPrice -got %v want %v",
			res.Drinks[s.ID].Name, drink.Name, res.Drinks[s.ID].Price, drink.Price)
	}
}

func TestGetDrinksEmpty(t *testing.T) {
	d := storage.GetInstance()
	res := d.GetDrinks()
	if len(res.Drinks) != 0 {
		t.Error("The storage should be empty")
	}
}

func TestGetDrink(t *testing.T) {
	d := storage.GetInstance()
	s := d.SaveDrink(drink) // Insert one record
	res, err := d.GetDrink(s.ID)
	if res != s {
		t.Errorf("The storage returned unexpected body: got %v but expected %v", res, s)
	}
	if err != nil {
		t.Error("The error should be empty")
	}
}

func TestGetDrinkNotFound(t *testing.T) {
	d := storage.GetInstance()

	_, err := d.GetDrink(id)
	if err.Error() != fmt.Sprintf("could not find drink with id = %s", id) {
		t.Error("The storage should return an error")
	}
}

func TestSaveDrink(t *testing.T) {
	d := storage.GetInstance()
	res := d.SaveDrink(drink)
	if res.ID == "" ||
		res.Name != drink.Name ||
		res.Price != drink.Price {
		t.Errorf("ID shouldn't be empty\nName - got %s but expected %s\nPrice -got %v want %v",
			res.Name, drink.Name, res.Price, drink.Price)
	}
}

func TestDeleteDrink(t *testing.T) {
	d := storage.GetInstance()
	s := d.SaveDrink(drink) // Insert one record
	err := d.DeleteDrink(s.ID)
	if err != nil {
		t.Error("The error should be empty")
	}
	_, err = d.GetDrink(s.ID)
	if err == nil {
		t.Error("The record shouldn't exist")
	}
}

func TestDeleteDrinkNotFound(t *testing.T) {
	d := storage.GetInstance()
	err := d.DeleteDrink(id)
	if err.Error() != fmt.Sprintf("could not find drink with id = %s", id) {
		t.Error("The storage should return error")
	}
}
