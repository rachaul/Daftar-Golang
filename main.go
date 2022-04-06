package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Tinggal   string `json:"tinggal"`
	Ktp       int    `json:"ktp"`
	Handphone int    `json:"handphone"`
	Lahir     int    `json:"lahir"`
}

var People = []User{
	{ID: "1", Username: "Rachaul", Email: "rachaul@gmail.com", Tinggal: "Jakarta", Ktp: 12345, Handphone: 0213452, Lahir: 05 - 03 - 03},
	{ID: "2", Username: "Eren", Email: "eren@gmail.com", Tinggal: "Eldia", Ktp: 07070707, Handphone: 02154363, Lahir: 01 - 02 - 98},
	{ID: "3", Username: "Mikasa", Email: "mikasa@gmail.com", Tinggal: "Eldia", Ktp: 12345678, Handphone: 0212456, Lahir: 02 - 67 - 05},
	{ID: "4", Username: "Jean", Email: "jean@gmail.com", Tinggal: "Eldia", Ktp: 66666666, Handphone: 021362345, Lahir: 04 - 06 - 01},
	{ID: "5", Username: "Sugiono", Email: "sugiono@gmail.com", Tinggal: "Jepang", Ktp: 77777777, Handphone: 021365222, Lahir: 01 - 02 - 06},
}

func getDaftar(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, People)
}

func addDaftar(context *gin.Context) {
	var newDaftar User

	if err := context.BindJSON(&newDaftar); err != nil {
		return
	}

	People = append(People, newDaftar)

	context.IndentedJSON(http.StatusCreated, newDaftar)
}

func getDaftarUser(context *gin.Context) {
	id := context.Param("id")
	user, err := getDaftarById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Usernot Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getDaftarById(id string) (*User, error) {
	for i, t := range People {
		if t.ID == id {
			return &People[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func main() {
	router := gin.Default()
	router.GET("/daftar", getDaftar)
	router.GET("/daftar:id", getDaftarUser)
	router.POST("/daftar", addDaftar)
	router.Run("localhost:9090")
}
