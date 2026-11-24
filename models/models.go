package models

type Avion struct {
	Icao24    string  `json:"icao24"`
	Indicatif string  `json:"indicatif"`
	Pays      string  `json:"pays"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Altitude  float64 `json:"altitude"`
	Vitesse   float64 `json:"vitesse"`
	AuSol     bool    `json:"au_sol"`
}

type ReponseOpenSky struct {
	Time   int64           `json:"time"`
	States [][]interface{} `json:"states"`
}
