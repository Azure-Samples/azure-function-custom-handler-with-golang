package main

import (
	"azure-function-custom-handler-with-golang/internal/drink"
	"azure-function-custom-handler-with-golang/internal/drinks"
	"azure-function-custom-handler-with-golang/pkg/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	storageHandler := storage.GetInstance()
	drinkManager := drink.DrinkManager{
		StorageHandler: storageHandler,
	}
	drinksManager := drinks.DrinksManager{
		StorageHandler: storageHandler,
	}
	http.HandleFunc("/api/drink", drinkManager.DrinkHandler)
	http.HandleFunc("/api/drinks", drinksManager.DrinksHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
