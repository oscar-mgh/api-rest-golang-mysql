package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oscar-mgh/api-rest-golang-mysql/common"
	"github.com/oscar-mgh/api-rest-golang-mysql/model"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	people := []model.Person{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&people)
	jsonData, _ := json.Marshal(people)
	common.SendResponse(w, http.StatusOK, jsonData)
}

func Get(w http.ResponseWriter, r *http.Request) {
	person := model.Person{}
	id := mux.Vars(r)["id"]
	db := common.GetConnection()
	defer db.Close()
	db.Find(&person, id)
	if person.ID > 0 {
		jsonData, _ := json.Marshal(person)
		common.SendResponse(w, http.StatusOK, jsonData)
	} else {
		common.SendError(w, http.StatusNotFound)
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
	person := model.Person{}
	db := common.GetConnection()
	defer db.Close()
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
		common.SendError(w, http.StatusBadRequest)
		return
	}
	err = db.Save(&person).Error
	if err != nil {
		log.Fatal(err)
		common.SendError(w, http.StatusInternalServerError)
		return
	}
	jsonData, _ := json.Marshal(person)
	common.SendResponse(w, http.StatusCreated, jsonData)
}

func Update(w http.ResponseWriter, r *http.Request) {
	person := model.Person{}
	id := mux.Vars(r)["id"]
	db := common.GetConnection()
	defer db.Close()
	db.Find(&person, id)
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
		common.SendError(w, http.StatusBadRequest)
		return
	}
	err = db.Save(&person).Error
	if err != nil {
		log.Fatal(err)
		common.SendError(w, http.StatusInternalServerError)
		return
	}
	if person.ID > 0 {
		jsonData, _ := json.Marshal(person)
		common.SendResponse(w, http.StatusOK, jsonData)
	} else {
		common.SendError(w, http.StatusNotFound)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	person := model.Person{}
	db := common.GetConnection()
	defer db.Close()
	id := mux.Vars(r)["id"]
	db.Find(&person, id)
	if person.ID > 0 {
		db.Delete(person)
		common.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(w, http.StatusNotFound)
	}
}
