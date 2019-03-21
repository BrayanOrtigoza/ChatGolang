package controller

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"ChatGolang/ChatGo/Constans"
	"ChatGolang/ChatGo/models"
	"fmt"
	"github.com/labstack/echo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"net/http"
)

type idchannel struct {
	Id_channel  string  `json:"id_channel" gorethink:"id_channel"`
}

 func makeArrayIdchannels(idChannelsUser *r.Cursor) []string {
		 var channels []idchannel

		 var arrayIdChannels []string

		 err := idChannelsUser.All(&channels)

		 if err != nil {
			 fmt.Println(err)
		 }

		 for _, element := range channels {
			 arrayIdChannels = append(arrayIdChannels, element.Id_channel)
		 }

		 return arrayIdChannels
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


func findDatesPrivateChannels(arrayIdChannels []string)  []string{

	res , err := r.Table("channel").
		Filter(func(p r.Term) interface{} {
			return r.Expr(arrayIdChannels).Contains(p.Field("id"))
		}).
		Filter(r.Row.Field("type").Eq(Constans.Private)).
		Pluck("id").
		Run(ConnectionDB.Session)

	var channels []models.Channel

	var arrayPrivatesChannel []string

	err = res.All(&channels)

	if err != nil {
		fmt.Println(err)
	}

	for _, element := range channels {
		arrayPrivatesChannel = append(arrayPrivatesChannel , element.Id)
	}

	return arrayPrivatesChannel
}

func makeChannel(id_people string) string {

	var idChannel string

	var channel models.Channel
	channel.Type =  Constans.Private
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


func MakeGroupChannel(c echo.Context)  (err error)  {

    var idChannel string
	var channel models.Channel
	channel.Type = Constans.Group
	channel.Name = "TI"
	res, err := r.Table("channel").Insert(channel).RunWrite(ConnectionDB.Session)

	if err != nil {
		fmt.Println(err)
	}

	if len(res.GeneratedKeys) > 0 {
		idChannel = res.GeneratedKeys[0]
	}

	var channelPeople models.ChannelPeople
	channelPeople.Id_people = "ff1ca430-42d3-4540-82aa-c2df9d570f37"
	channelPeople.Id_channel = idChannel
	_, err = r.Table("channel_people").Insert(channelPeople).RunWrite(ConnectionDB.Session)

	if err != nil {
		fmt.Println(err)
	}


	return c.String(http.StatusOK, "OK")

}



func FindChannelGroup (c echo.Context)  (err error) {
	updateIdUser(c)

	u := new(models.Channel)
	if err = c.Bind(u); err != nil {
		return err
	}
	return findMessageChannel(c,u.Id)
}


func ListDataGroupChannel(c echo.Context)  (err error) {

	updateIdUser(c)

	idChannelsUser, err := r.Table("channel_people").
		Filter(r.Row.Field("id_people").Eq(findIdUserPeople())).
		Run(ConnectionDB.Session)


	arrayIdChannels := makeArrayIdchannels(idChannelsUser)

	groupsChannel := findDatesGroupsChannels(arrayIdChannels)


	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "No tiene Grupos",
		})
	}

	return c.JSON(http.StatusOK, groupsChannel)
}


func FindChannelPeople(c echo.Context)  (err error)  {

	updateIdUser(c)

	u := new(models.ChannelPeople)
	if err = c.Bind(u); err != nil {
		return err
	}

	idChannelPeople, err := r.Table("channel_people").
		Filter(r.Row.Field("id_people").Eq(findIdUserPeople())).
		Pluck("id_channel").
		Run(ConnectionDB.Session)

	arrayIdChannels := makeArrayIdchannels(idChannelPeople)

	arrayPrivatesChannel := findDatesPrivateChannels(arrayIdChannels)


	res , err := r.Table("channel_people").
		Filter(func(p r.Term) interface{} {
			return r.Expr(arrayPrivatesChannel).Contains(p.Field("id_channel"))
		}).
		Filter(r.Row.Field("id_people").Eq(u.Id_people)).
		Pluck("id_channel").
		Run(ConnectionDB.Session)

	var channel idchannel

	err = res.One(&channel)


	if err != nil {
		channel.Id_channel = makeChannel(u.Id_people)
	}


	return findMessageChannel(c,channel.Id_channel)
}