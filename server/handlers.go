package main

import (
	"ChatGolang/ChatGo/Constans"
	"ChatGolang/ChatGo/models"
	"fmt"
	"time"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

const (
	ChannelStop = iota
	UserStop
	MessageStop
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

type idchannel struct {
	Id_channel  string  `json:"id_channel" gorethink:"id_channel"`
}

type ChannelMessage struct {
	Id          string    `json:"id" gorethink:"id,omitempty"`
	Message      string    `json:"message" gorethink:"message"`
	Id_channel  string    `json:"id_channel"  gorethink:"id_channel"`
	Id_people   string    `json:"id_people_message"  gorethink:"id_people_message"`
	Author      string    `json:"author" gorethink:"author"`
	CreatedAt   time.Time `json:"createdAt" gorethink:"createdAt"`
}

func subscribeChannelMessage(client *Client, data interface{}) {

	go func() {
		eventData := data.(map[string]interface{})

		val, ok := eventData["channelId"]
		if !ok {
			return
		}

		channelId, ok := val.(string)
		if !ok {
			return
		}
		stop := client.NewStopChannel(MessageStop)
		cursor, err := r.Table("message").
			OrderBy(r.OrderByOpts{Index: r.Desc("createdAt")}).
			Filter(r.Row.Field("id_channel").Eq(channelId)).
			Changes(r.ChangesOpts{IncludeInitial: false}).
			Run(client.session)

		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}

		changeFeedHelper(cursor, "message", client.send, stop)
	}()
}

func UserStatus(client *Client, data interface{}) {

	go func() {
		eventData := data.(map[string]interface{})

		val, ok := eventData["idUser"]
		if !ok {
			return
		}

		idUser, ok := val.(string)
		if !ok {
			return
		}
		stop := client.NewStopChannel(MessageStop)
		cursor, err := r.Table("people").Filter(r.Not(r.Row.Field("id_user").Eq(idUser))).Changes().Run(client.session)

		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}

		changeFeedHelper(cursor, "People", client.send, stop)
	}()
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


func findDatesPrivateChannels(arrayIdChannels []string,client *Client)  []string{

	res , err := r.Table("channel").
		Filter(func(p r.Term) interface{} {
			return r.Expr(arrayIdChannels).Contains(p.Field("id"))
		}).
		Filter(r.Row.Field("type").Eq(Constans.Private)).
		Pluck("id").
		Run(client.session)

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


func ChannelPeople(client *Client, data interface{})  {
	go func() {
		eventData := data.(map[string]interface{})

		val, ok := eventData["idPeopleUser"]
		if !ok {
			return
		}

		idPeopleUser, ok := val.(string)
		if !ok {
			return
		}
		fmt.Println(idPeopleUser)
		stop := client.NewStopChannel(MessageStop)
		idChannelPeople, err := r.Table("channel_people").
			Filter(r.Row.Field("id_people").Eq(idPeopleUser)).
			Pluck("id_channel").
			Run(client.session)

		arrayIdChannels := makeArrayIdchannels(idChannelPeople)

		arrayPrivatesChannel := findDatesPrivateChannels(arrayIdChannels, client)

fmt.Println(arrayPrivatesChannel)

		cursor, err := r.Table("message").Filter(func(p r.Term) interface{} {
			return r.Expr(arrayPrivatesChannel).Contains(p.Field("id_channel"))
		}).Changes().Run(client.session)

		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}

		changeFeedHelper(cursor, "MessageCount", client.send, stop)
	}()
}


func changeFeedHelper(cursor *r.Cursor, changeEventName string,
	send chan<- Message, stop <-chan bool) {
	change := make(chan r.ChangeResponse)
	cursor.Listen(change)
	fmt.Println(changeEventName)

	for {
		eventName := ""
		var data interface{}

		select {
		case <-stop:
			cursor.Close()
			return
		case val := <-change:
			if val.NewValue != nil && val.OldValue == nil {
				eventName = changeEventName + " add"
				data = val.NewValue
			} else if val.NewValue == nil && val.OldValue != nil {
				eventName = changeEventName + " remove"
				data = val.OldValue
			} else if val.NewValue != nil && val.OldValue != nil {
				eventName = changeEventName + " edit"
				data = val.NewValue
			}
			fmt.Println(eventName)
			fmt.Println(data)
			send <- Message{eventName, data}
		}
	}
}
