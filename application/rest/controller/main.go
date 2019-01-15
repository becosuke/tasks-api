package main

import (
	"log"

	"github.com/becosuke/tasks-api/application/rest/router"
)

func main() {
	var err error
	r := router.NewRouter()

	if err = r.Setup(); err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
