package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcodenisi/eshop-tracker/service"
)

func main() {
	client := http.Client{}
	games, err := service.RetrieveEuGames(&client)
	if err != nil {
		log.Fatal("Error while retrieving eu games", err)
	}
	fmt.Printf("Retrieved %v games", len(games))
}
