package controller

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"ChatGolang/ChatGo/models"
	"github.com/labstack/echo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"net/http"
)

func ListDataPeople(c echo.Context) error {

	updateIdUser(c)

	res, err := r.Table("people").
		Filter(r.Not(r.Row.Field("id_user").Eq(idUser))).
		Run(ConnectionDB.Session)

	var datapeople []models.People

	err = res.All(&datapeople)

	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "No hay Usuario",
		})
	}
	return c.JSON(http.StatusOK, datapeople)

}
