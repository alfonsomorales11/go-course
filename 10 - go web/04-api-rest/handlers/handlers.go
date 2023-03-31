package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"paquetes/db"
	"paquetes/models"
	"strconv"
	"strings"
)

func GetRequestId(r *http.Request) int {
	id := strings.Split(r.RequestURI, "/api/user/")[1]
	idInteger, _ := strconv.Atoi(id)
	return idInteger
}

func GetBodyJSON(r *http.Request) *models.User {
	body, _ := ioutil.ReadAll(r.Body)
	bodyJson := string(body)

	usuario := models.User{}
	json.Unmarshal([]byte(bodyJson), &usuario)

	return &usuario
}

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	db.Connect()
	users := models.ShowList()
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	id := GetRequestId(r)

	db.Connect()
	user := models.GetUser(id)

	response, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(response))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	usuario := GetBodyJSON(r)

	db.Connect()
	models.Insertar(usuario.UserName, usuario.Password, usuario.Email)
	db.Close()
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	id := GetRequestId(r)

	usuario := GetBodyJSON(r)

	db.Connect()
	models.UpdateUser(id, usuario.UserName, usuario.Password, usuario.Email)
	db.Close()
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	id := GetRequestId(r)

	db.Connect()
	models.DeleteUser(id)
	db.Close()
}
