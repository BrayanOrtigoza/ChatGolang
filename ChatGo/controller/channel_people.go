package controller

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"ChatGolang/ChatGo/models"
	"fmt"
	"github.com/labstack/echo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type groupIdChannel struct {
	Id_chanel  string  `json:"id_chanel" gorethink:"id_channel"`
}

func FindChannelPeople(c echo.Context)  (err error)  {

	updateIdUser(c)

	u := new(models.ChannelPeople)
	if err = c.Bind(u); err != nil {
		return err
	}

	resChannelPeople, err := r.Table("channel_people").
		Filter(r.Row.Field("id_people").Eq(findIdUserPeople())).
		Pluck("id_channel").
		Run(ConnectionDB.Session)

	var channels []groupIdChannel

	var id_channels []string

	err = resChannelPeople.All(&channels)


	for _, element := range channels {
		id_channels = append(id_channels , element.Id_chanel)
	}


	res , err := r.Table("channel_people").
		Filter(func(p r.Term) interface{} {
			return r.Expr(id_channels).Contains(p.Field("id_channel"))
		}).
		Filter(r.Row.Field("id_people").Eq(u.Id_people)).
		Pluck("id_channel").
		Run(ConnectionDB.Session)

	var channel groupIdChannel

	err = res.One(&channel)


	if err != nil {
		channel.Id_chanel = makeChannel(u.Id_people)
	}

	fmt.Println(channel.Id_chanel)

	return findMessageChannel(c,channel.Id_chanel)
}


func makeChannel(id_people string) string {

	var idChannel string

	var channel models.Channel
	res, err := r.Table("channel").Insert(channel).RunWrite(ConnectionDB.Session)

	if err != nil {
		fmt.Println(err)
	}

	if len(res.GeneratedKeys) > 0 {
		idChannel = res.GeneratedKeys[0]
	}


	if !(makeChannelUser(idChannel) &&  makeChannelPeople(id_people,idChannel)){
		fmt.Println("Error canales no creados")
	}

	fmt.Println(idChannel)

	return idChannel

}

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
	fmt.Println(idChannel)
	var chanel_people models.ChannelPeople

	chanel_people.Id_channel = idChannel
	chanel_people.Id_people = id_people

	_,err := r.Table("channel_people").Insert(chanel_people).Run(ConnectionDB.Session)

	if err != nil {
		fmt.Println(err)
	}
	return true

}



