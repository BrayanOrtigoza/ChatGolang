package controller

import (
	"ChatGolang/ChatGo/ConnectionDB"
	"ChatGolang/ChatGo/Constans"
	"ChatGolang/ChatGo/models"
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"net/http"
	"time"
)

type AuthToken struct {
	IdUser string
	jwt.StandardClaims
}

var idUser string

func updateIdUser(c echo.Context) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*AuthToken)
	idUser = claims.IdUser
}


func LoginUser(c echo.Context) (err error)  {

	u := new(models.DataUser)
	if err = c.Bind(u); err != nil {
		return err
	}
	filter := make(map[string]interface{})

	filter["username"] = u.Username
	filter["password"] = u.Password


	res, err := r.Table("data_user").Filter(filter).Pluck("id").Run(ConnectionDB.Session)

	var data_user models.DataUser
	err = res.One(&data_user)

	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "Usuario y/o Contraseña",
		})
	}

	idUser = data_user.Id

	return createToken(c, data_user.Id)

}


func createToken(c echo.Context,idUser string) error{

	claims := &AuthToken{
		idUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}


	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func FindDataUser(c echo.Context) error {

	updateIdUser(c)

	filter := make(map[string]interface{})

	filter["id_user"] = idUser
	fmt.Println(idUser)
	res, err := r.Table("people").Filter(filter).Run(ConnectionDB.Session)

	changeStatusPeople("login")



	var people models.People

	err = res.One(&people)

	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "Usuario no encontrado",
		})
	}

	return c.JSON(http.StatusOK, people)

}

func changeStatusPeople(action string){


	filter := make(map[string]interface{})

	filter["id_user"] = idUser

	update := make(map[string]interface{})



	if action == "login"{
		update["status"] = Constans.Active

		_, err := r.Table("people").Filter(filter).Update(update).Run(ConnectionDB.Session)

		if err != nil {
			fmt.Println("Usuario Activo")
		}
	}
	if action == "logout"{
		update["status"] = Constans.Inactive

		_, err := r.Table("people").Filter(filter).Update(update).Run(ConnectionDB.Session)
		if err != nil {
			fmt.Println("Usuario Inactivo")
		}
	}
}

func LogOutUser(c echo.Context) error {
		updateIdUser(c)

		changeStatusPeople("logout")


	return c.JSON(http.StatusOK, echo.Map{
		"message": "Salio de sesion",
	})
}

func findNameUser() string {

	filter := make(map[string]interface{})

	filter["id_user"] = idUser

	res, err := r.Table("people").
	    Filter(filter).
		Pluck("name","last_name").
		Run(ConnectionDB.Session)

	var user models.People

	err = res.One(&user)

	if err != nil {
		 fmt.Println("Usuario no encontrado")
	}

	return ConcatNameUser(user.Name, user.LastName)

}

func findIdUserPeople() string {

	filter := make(map[string]interface{})

	filter["id_user"] = idUser

	res, err := r.Table("people").
		Filter(filter).
		Pluck("id").
		Run(ConnectionDB.Session)

	var user models.People

	err = res.One(&user)

	if err != nil {
		fmt.Println("Usuario no encontrado")
	}

	return user.Id

}


func ConcatNameUser(name string, last_name string) string {
	var buffer bytes.Buffer
	buffer.WriteString(name)
	buffer.WriteString(Constans.Space)
	buffer.WriteString(last_name)
	fullName := buffer.String()
	return fullName
}
