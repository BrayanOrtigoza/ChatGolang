package main

import (

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"log"
	"net/http"
)

func main() {

	session, err := r.Connect(r.ConnectOpts{
		Address: "127.0.0.1:28015",
		Database: "chat3",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)


	router.Handle("message subscribe", subscribeChannelMessage)
	router.Handle("People Data", UserStatus)
	router.Handle("Channel People", ChannelPeople)
	http.Handle("/", router)


	http.ListenAndServe(":4000", nil)
}
