package websocket

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"net/http"
)

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Router struct {
	rules   map[string]Handler
	session *r.Session
}

func NewRouter(session *r.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}

func makeRoutesWebsocket() {
	router := NewRouter(ConnectionDB.Session)

	router.Handle("message subscribe", subscribeChannelMessage)
	fmt.Println("entro2")
}

func InitWebsocket(c echo.Context) (err error)  {

	fmt.Println("entro")
	var e *Router

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	client := NewClient(ws, e.FindHandler, e.session)

	defer client.Close()

	go client.Write()

	client.Read()


	makeRoutesWebsocket()

	return nil

	/*router := NewRouter(ConnectionDB.Session)

	router.Handle("message subscribe", subscribeChannelMessage)

	return c.String(http.StatusOK, "Websocket iniciado")*/
}