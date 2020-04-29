package service

import (
	"fmt"
	"log"
	"net/http"
	_url "net/url"
	"sync"

	"github.com/marcodenisi/eshop-tracker/model"
)

const schema = "https"
const host = "search.nintendo-europe.com"
const baseOpaqueURL = "//search.nintendo-europe.com/it/select?fq=type:GAME AND system_type:nintendoswitch* AND product_code_txt:*&q=*&sort=sorting_title asc&wt=json"
const queryParams = "&rows=%v&start=%v"
const step = 100

// RetrieveEuGames retrieves all EU games
func RetrieveEuGames(client model.HTTPClient) ([]model.EuGame, error) {
	totalGamesNumber, err := getTotalGamesNumber(client)
	if err != nil {
		return nil, err
	}

	var (
		mu      = &sync.Mutex{}
		euGames = []model.EuGame{}
	)

	var wg sync.WaitGroup
	for i := 0; i*step < totalGamesNumber; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			opaqueURL := baseOpaqueURL + fmt.Sprintf(queryParams, step, idx*step)
			dto, err := performGetGames(opaqueURL, client)
			if err != nil {
				log.Fatal("Error", err)
				return
			}
			mu.Lock()
			euGames = append(euGames, dto.Response.Games...)
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	return euGames, nil
}

func getTotalGamesNumber(client model.HTTPClient) (int, error) {
	opaqueURL := baseOpaqueURL + fmt.Sprintf(queryParams, 0, 0)
	dto, err := performGetGames(opaqueURL, client)
	if err != nil {
		return 0, err
	}
	return dto.Response.NumFound, nil
}

func performGetGames(url string, client model.HTTPClient) (*model.EuGamesResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.URL = &_url.URL{
		Scheme: schema,
		Host:   host,
		Opaque: url,
	}

	if err != nil {
		log.Fatal("Error while creating request ", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error while performing http call", err)
		return nil, err
	}

	dto := model.DecodeEuGamesResponse(resp)
	return &dto, nil
}
