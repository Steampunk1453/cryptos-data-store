package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cryptos-data-store/models"
	u "github.com/cryptos-data-store/utils"

	"github.com/gorilla/mux"
)

func CreateCrypto(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	crypto := &models.Crypto{}

	err := json.NewDecoder(r.Body).Decode(crypto)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	crypto.UserId = user
	resp := crypto.Create()
	u.Respond(w, resp)
}

func FindCryptoById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.Get(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func FindAllCrypto(w http.ResponseWriter) {

	data := models.GetAll()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func FindCryptoByUserId(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetUserCrypto(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func UpdateCrypto(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	crypto := &models.Crypto{}

	err := json.NewDecoder(r.Body).Decode(crypto)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	crypto.UserId = user
	resp := crypto.Update(*crypto)
	u.Respond(w, resp)
}

func DeleteCrypto(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.Delete(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
