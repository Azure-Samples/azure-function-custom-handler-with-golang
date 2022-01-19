package drink_test

import (
	"azure-function-custom-handler-with-golang/internal/drink"
	"azure-function-custom-handler-with-golang/pkg/storage"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	sampleDrink = storage.Drink{
		ID:    "test",
		Name:  "Water",
		Price: 1,
	}
	expected = `{"id":"test","name":"Water","price":1}`
)

type mockStorageHandler struct{}

func (m *mockStorageHandler) GetDrinks() storage.Drinks {
	return storage.Drinks{}
}
func (m *mockStorageHandler) GetDrink(id string) (storage.Drink, error) {
	return sampleDrink, nil
}
func (m *mockStorageHandler) SaveDrink(storage.Drink) storage.Drink {
	return sampleDrink
}
func (m *mockStorageHandler) DeleteDrink(id string) (err error) {
	return nil
}

func TestDrinkHandlerGET(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/drink", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("id", "test")
	runTest(t, req, expected)
}

func TestDrinkHandlerPOST(t *testing.T) {
	body := storage.Drink{
		Name:  "Water",
		Price: 1,
	}
	bodyByte, _ := json.Marshal(body)
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/drink", bytes.NewBuffer(bodyByte))
	if err != nil {
		t.Fatal(err)
	}
	runTest(t, req, expected)
}

func TestDrinkHandlerDELETE(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/api/drink", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	runTest(t, req, "")
}

func runTest(t *testing.T, req *http.Request, expected string) {
	drinksManager := drink.DrinkManager{
		&mockStorageHandler{},
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(drinksManager.DrinkHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	if rr.Body.String() != expected {
		t.Errorf("The handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
