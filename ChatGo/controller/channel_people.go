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


/*
func makeArrayIdchannels(idChannelsUser *r.Cursor) []models.Channel  {

	var channels []idchannel

	var arrayIdChannels []string

	err := idChannelsUser.All(&channels)

	if err != nil {
		fmt.Println(err)
	}

	for _, element := range channels {
		arrayIdChannels = append(arrayIdChannels, element.Id_channel)
	}

	return findDatesGroupsChannels(arrayIdChannels)
}

func findDatesGroupsChannels(arrayIdChannels []string)  []models.Channel{

	res, err := r.Table("channel").
		Filter(func(p r.Term) interface{} {
			return r.Expr(arrayIdChannels).Contains(p.Field("id"))
		}).
		Filter(r.Row.Field("type").Eq(Constans.Group)).
		Run(ConnectionDB.Session)

	if err != nil {
		fmt.Println(err)
	}


	var groupsChannel []models.Channel

	err = res.All(&groupsChannel)

	return groupsChannel
}


func ListDataGroupChannel(c echo.Context)  (err error) {

	updateIdUser(c)

	idChannelsUser, err := r.Table("channel_people").
		Filter(r.Row.Field("id_people").Eq(findIdUserPeople())).
		Run(ConnectionDB.Session)

	groupsChannel :=  makeArrayIdchannels(idChannelsUser)


	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "No tiene Grupos",
		})
	}

	return c.JSON(http.StatusOK, groupsChannel)

}

*/