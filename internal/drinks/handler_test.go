package drinks_test

import (
	"azure-function-custom-handler-with-golang/internal/drinks"
	"azure-function-custom-handler-with-golang/pkg/storage"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockStorageHandler struct{}

func (m *mockStorageHandler) GetDrinks() storage.Drinks {
	id := "test"
	drinkStorage := storage.Drinks{Drinks: make(map[string]storage.Drink)}
	drinkStorage.Drinks[id] = storage.Drink{
		ID:    id,
		Name:  "Water",
		Price: 1,
	}
	return drinkStorage
}

func (m *mockStorageHandler) GetDrink(id string) (drink storage.Drink, err error) {
	return storage.Drink{}, nil
}
func (m *mockStorageHandler) SaveDrink(drink storage.Drink) storage.Drink {
	return storage.Drink{}
}
func (m *mockStorageHandler) DeleteDrink(id string) (err error) {
	return nil
}

func TestDrinksHandler(t *testing.T) {
	drinksManager := drinks.DrinksManager{
		&mockStorageHandler{},
	}

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/drinks", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(drinksManager.DrinksHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Drinks":{"test":{"id":"test","name":"Water","price":1}}}`
	if rr.Body.String() != expected {
		t.Errorf("The handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
