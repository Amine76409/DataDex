package main

import (
	"fmt"

	"datadex/api"
)

func main() {
	avions, err := api.RecupererAvions()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("Nombre d'avions :", len(avions))
	if len(avions) > 0 {
		fmt.Println(avions[0])
	}
}
