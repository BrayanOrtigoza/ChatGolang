package Router

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"ChatGolang/ChatGo/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)



func startDataBase()  {
	ConnectionDB.Init()
}

func configurationCors(e *echo.Echo)  {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}




func InitRoutes()  {

	go startDataBase()

	e := echo.New()

	configurationCors(e)

	publicRouters(e)

	privateRouters(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func publicRouters(e *echo.Echo)  {

	e.POST("/login", controller.LoginUser)


}

func privateRouters(e *echo.Echo)  {

	configurationCors(e)

	// Restricted group
	r := e.Group("/Auth")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &controller.AuthToken{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))



	r.GET("/findUser", controller.FindDataUser)

	r.GET("/ListDataPeople", controller.ListDataPeople)

	r.POST("/FindInitialMessage", controller.FindChannelPeople)

	r.POST("/MakeMessage", controller.InsertNewMessage)



}


