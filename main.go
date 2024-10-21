package main

import (
	"log"

	"Social-Network-01/api"
)

func main() {
	api, err := api.NewAPI(":3001", "api/db/db.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Server listening on port %s", api.Addr)
	err = api.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
