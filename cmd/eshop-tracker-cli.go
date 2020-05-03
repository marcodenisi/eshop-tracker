package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
	"github.com/marcodenisi/eshop-tracker/db"
	"github.com/marcodenisi/eshop-tracker/model"
	"github.com/marcodenisi/eshop-tracker/service"
)

func main() {
	ch := make(chan bool)
	go showPrompt(ch)
	<-ch
	fmt.Println("Thanks for using Eshop-Tracker-Cli")
}

func showPrompt(ch chan bool) {
	prompt := promptui.Select{
		Label: "Select Operation",
		Items: []string{"Get All Games", "Get Games by Name", "Refresh Game Prices", "Quit"},
	}

	for {
		result, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case 0:
			getAll()
		case 1:
			getByName()
		case 2:
			refresh()
		case 3:
			ch <- true
		}
	}

}

func refresh() {
	client := http.Client{}
	games, err := service.RetrieveEuGames(&client)
	if err != nil {
		log.Fatal("Error while retrieving eu games", err)
	}
	db.SaveGames(games)
	fmt.Printf("Updated %v games\n", len(games))
}

func getAll() {
	g, err := db.GetGames()
	if err != nil {
		log.Fatal("Error while retrieving games.", err)
		return
	}
	printGames(g)
}

func getByName() {
	namePrompt := promptui.Prompt{
		Label: "Enter the game you're interested in",
		Validate: func(input string) error {
			if len(input) < 3 {
				return errors.New("Search term must have at least 3 characters")
			}
			return nil
		},
	}
	name, err := namePrompt.Run()
	if err != nil {
		fmt.Printf("Game Prompt failed %v\n", err)
	}

	g, err := db.GetGamesFromName(name)
	if err != nil {
		log.Fatal("Error while retrieving games.", err)
		return
	}
	printGames(g)
}

func printGames(games []model.EuGame) {
	templates := promptui.SelectTemplates{
		Active:   `🍄 {{ .Title | cyan | bold }}`,
		Inactive: `   {{ .Title | cyan }}`,
		Selected: `{{ "✔" | green | bold }} {{ "Game" | bold }}: {{ .Title | cyan }} currently available at {{ .PriceRegularF }}€`,
		Details:  `Categories: {{ .PrettyGameCategoriesTxt }}`,
	}

	list := promptui.Select{
		Label:     "Games",
		Items:     games,
		Templates: &templates,
	}

	_, _, err := list.Run()
	if err != nil {
		log.Fatal("Error while selecting game", err)
		return
	}
}
