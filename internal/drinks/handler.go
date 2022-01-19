// Package drinks contains the http handler for the drinks API
package drinks

import (
	"azure-function-custom-handler-with-golang/pkg/log"
	"azure-function-custom-handler-with-golang/pkg/storage"
	"encoding/json"
	"net/http"
)

type DrinksManager struct {
	StorageHandler storage.IstorageHandler
}

func (d *DrinksManager) DrinksHandler(w http.ResponseWriter, r *http.Request) {
	res := d.StorageHandler.GetDrinks()
	resByte, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	log.LogErr(w.Write(resByte))
}
