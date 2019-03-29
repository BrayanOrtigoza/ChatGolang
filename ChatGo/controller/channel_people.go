package controller

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"ChatGolang/ChatGo/models"
	"fmt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)



func makeChannelUser(idChannel string) bool  {

	var chanel_people models.ChannelPeople

	chanel_people.Id_channel = idChannel
	chanel_people.Id_people = findIdUserPeople()

	_, err := r.Table("channel_people").Insert(chanel_people).Run(ConnectionDB.Session)

	if err != nil {
		fmt.Println(err)
	}
	return true
}

func makeChannelPeople(id_people string,idChannel string) bool {

	var chanel_people models.ChannelPeople

	chanel_people.Id_channel = idChannel
	chanel_people.Id_people = id_people

	_,err := r.Table("channel_people").Insert(chanel_people).Run(ConnectionDB.Session)

	if err != nil {
		fmt.Println(err)
	}
	return true

}