package main

import (
	"log"
	"os"

	"Social-Network-01/api"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("no command")
	}

	api, err := api.NewAPI(":3001", "api/db/db.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}

	switch os.Args[1] {
	case "up":
		err = api.Storage.Up("api/db/migrations")

	case "down":
		err = api.Storage.Down("api/db/migrations")

	case "serve":
		log.Printf("Starting server on port %s", api.Addr)
		err = api.ListenAndServe()
	}

	if err != nil {
		log.Fatalln(err)
	}
}
