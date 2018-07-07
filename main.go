package main

import (
	"github.com/skamenetskiy/jsonapi"
	"github.com/compo-io/user/ctrl"
	"log"
	"github.com/compo-io/db"
	"github.com/compo-io/user/user"
)

// entry point
func main() {
	// initialize database
	if err := db.Init(nil, nil); err != nil {
		log.Fatal(err)
	}

	// initialize server
	server := jsonapi.
		NewServer(user.Addr).
		Controller("/", new(ctrl.Controller))
	if err := server.Listen(); err != nil {
		log.Fatal(err)
	}
}
