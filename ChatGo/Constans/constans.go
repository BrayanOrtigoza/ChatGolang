package Constans

import (
	"bytes"
)

const(

		colon = ":"
		Empty = ""
		Ping  = "PING"
		Pong  = "PONG"
		Space = " "


		//rethinkdb
		DefaultHost   =         "127.0.0.1"
		Port          =         "28015"
		NameDatabases =         "chat3"
	    SecondsReconnectRethinkDB =  30
	    UserRethinkDBServer     =  ""
	    PasswordRethinkDBServer =  ""

        //status to login
        Enabled   =  1
		disabled  =  0

		//status to user
		active    =  1
		inactive   = 0
)

func ConcatHostPortRethinkdb(host string, port string) string {
	var buffer bytes.Buffer
	buffer.WriteString(host)
	buffer.WriteString(colon)
	buffer.WriteString(port)
	addressRethinkdb := buffer.String()
	return addressRethinkdb
}
