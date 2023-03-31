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
	var users models.Users
	db.Database.Find(&models.Users{}).Scan(&users)
	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	id := GetRequestId(r)
	user := models.User{}
	db.Database.First(&models.User{}, id).Scan(&user)
	response, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(response))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	usuario := GetBodyJSON(r)
	db.Database.Create(usuario)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	id := GetRequestId(r)
	usuario := GetBodyJSON(r)
	usuario.Id = int64(id)
	fmt.Println(usuario)
	db.Database.Save(usuario)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	id := GetRequestId(r)
	db.Database.Delete(models.User{}, id)
	GetUsers(rw, r)
}
