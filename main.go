package main

import (
	"Social-Network-01/api"
	"log"
)

func main() {
	api, err := api.NewAPI(":3000")
	if err != nil {
		log.Fatalln(err)
	}

	err = api.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
