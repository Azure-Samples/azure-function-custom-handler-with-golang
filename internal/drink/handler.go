// Package drink contains the http handler for the drink API
package drink

import (
	"azure-function-custom-handler-with-golang/pkg/log"
	"azure-function-custom-handler-with-golang/pkg/storage"
	"encoding/json"
	"net/http"
)

type DrinkManager struct {
	StorageHandler storage.IstorageHandler
}

func (d *DrinkManager) DrinkHandler(w http.ResponseWriter, r *http.Request) {
	var resByte []byte
	switch r.Method {
	case http.MethodGet:
		resByte = d.getHandler(w, r)
	case http.MethodPost:
		resByte = d.postHandler(w, r)
	case http.MethodDelete:
		d.deleteHandler(w, r)
	}
	w.Header().Set("Content-Type", "application/json")
	log.LogErr(w.Write(resByte))
}

func (d *DrinkManager) getHandler(w http.ResponseWriter, r *http.Request) []byte {
	id := r.Header.Get("id")
	if res, err := d.StorageHandler.GetDrink(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil
	} else {
		resByte, _ := json.Marshal(res)
		return resByte
	}
}

func (d *DrinkManager) postHandler(w http.ResponseWriter, r *http.Request) []byte {
	var drink storage.Drink
	err := json.NewDecoder(r.Body).Decode(&drink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	res := d.StorageHandler.SaveDrink(drink)
	resByte, _ := json.Marshal(res)
	return resByte
}

func (d *DrinkManager) deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	if err := d.StorageHandler.DeleteDrink(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
