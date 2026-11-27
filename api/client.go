package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"datadex/models"
)

const urlOpenSky = "https://opensky-network.org/api/states/all?lamin=41&lomin=-5&lamax=51.5&lomax=10"

func RecupererAvions() ([]models.Avion, error) {
	reponse, err := http.Get(urlOpenSky)
	if err != nil {
		return nil, fmt.Errorf("API injoignable")
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != 200 {
		return nil, fmt.Errorf("erreur API, code %d", reponse.StatusCode)
	}
}
