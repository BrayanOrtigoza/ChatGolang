package controller

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"ChatGolang/ChatGo/models"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"log"
	"net/http"
	"time"
)

func findMessageChannel(c echo.Context, idChannel string) (err error) {

	filter := make(map[string]interface{})

	filter["id_channel"] = idChannel


	res, err := r.Table("message").
		OrderBy(r.OrderByOpts{Index: r.Asc("createdAt")}).
		Filter(filter).
		Run(ConnectionDB.Session)

	var dataMessage []models.Message

	err = res.All(&dataMessage)



	return c.JSON(http.StatusOK, echo.Map{
		"dataMessages":dataMessage,
		"idChannel":idChannel,
		"iduser": findIdUserPeople(),
	})
}

func Message(c echo.Context) (err error) {
	filter := make(map[string]interface{})

	filter["id_channel"] = "aa9c376a-6bda-4829-99c8-d6198f099521"

	res, err := r.Table("message").
		OrderBy(r.OrderByOpts{Index: r.Asc("createdAt")}).
		Filter(filter).
		Run(ConnectionDB.Session)

	var dataMessage []models.Message

	err = res.All(&dataMessage)

	/*if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "No hay Usuario",
		})
	}*/
	return c.JSON(http.StatusOK, dataMessage)
}


func InsertNewMessage(c echo.Context) (err error)  {

	updateIdUser(c)

	u := new(models.Message)
	if err = c.Bind(u); err != nil {
		return err
	}
	var message models.Message
	var id string


	message.Message = u.Message
	message.Author = findNameUser()
	message.CreatedAt = time.Now()
	message.Id_channel = u.Id_channel
	message.Id_people = findIdUserPeople()


	res, err := r.Table("message").Insert(message).RunWrite(ConnectionDB.Session)

	if err != nil {
		log.Println(err.Error())
	}

	if len(res.GeneratedKeys) > 0 {
		id = res.GeneratedKeys[0]
	}

	return c.JSON(http.StatusOK, id)
}


func WriteChangesMessage(ws *websocket.Conn) {
	go func() {

		res, err := r.Table("message").
			Filter(r.Row.Field("id_channel").Eq("aa9c376a-6bda-4829-99c8-d6198f099521")).
			Changes(r.ChangesOpts{IncludeInitial: false}).
			Run(ConnectionDB.Session)

		if err != nil {
			log.Println(err.Error())
		}

		var message models.Message

		err = res.One(&message)

		changeFeedHelper(res, "message")
	}()
}

type Messages struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send         chan Messages
	socket       *websocket.Conn
	session      *r.Session
	stopChannels map[int]chan bool
}

func changeFeedHelper(cursor *r.Cursor, changeEventName string) {
	change := make(chan r.ChangeResponse)
	cursor.Listen(change)
    fmt.Println("Hello")
	for {
		eventName := ""
		var data interface{}

		select {
		case val := <-change:
			fmt.Println(val)
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
			fmt.Println(data, eventName)
		}
	}
}