package main

import (
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

func changeFeedHelper(cursor *r.Cursor, changeEventName string,
	send chan<- Message, stop <-chan bool) {
	change := make(chan r.ChangeResponse)
	cursor.Listen(change)

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
			send <- Message{eventName, data}
		}
	}
}
