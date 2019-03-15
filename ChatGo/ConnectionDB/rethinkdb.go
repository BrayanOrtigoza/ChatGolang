package ConnectionDB

import (
	"API/constants"
	"ChatGolang/ChatGo/Constans"
	"fmt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"time"
)

var (
	Session *r.Session
)

func newconecctionRethinkDB() error {

	var err error
	Session, err = r.Connect(r.ConnectOpts{
		Address: Constans.ConcatHostPortRethinkdb(constants.DefaultHost,Constans.Port),
		Database: Constans.NameDatabases,
		Username: Constans.UserRethinkDBServer,
		Password: Constans.PasswordRethinkDBServer,
	})

	if err != nil {
		fmt.Println(err)
		reconnectRethinkDB()
	}

	return err

}

func reconnectRethinkDB() {
	err := newconecctionRethinkDB()
	for {
		if !veryfyConenctionRethinkDB() || err != nil {
			err = newconecctionRethinkDB()
		}
		waitTimeSeconds := time.Duration(Constans.SecondsReconnectRethinkDB)
		time.Sleep(waitTimeSeconds * time.Second)
	}
}

func veryfyConenctionRethinkDB() bool {
	return Session.IsConnected()
}


func Init() {
	go newconecctionRethinkDB()
}

