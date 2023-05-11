package controller

import (
	"encoding/json"
	"final-project/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type authorizationBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type depositBody struct {
	Balance float32 `json:"balance"`
}

func (c *Controller) register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var auth model.User

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	err = c.services.Authorization.RegisterUser(auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var auth authorizationBody

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	token, refreshToken, err := c.services.Authorization.GenerateToken(auth.Username, auth.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
	})

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	header := r.Header.Get("Authorization")
	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		w.WriteHeader(401)
		return
	}

	token, refreshToken, err := c.services.Authorization.RefreshToken(headerParts[1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
	})

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) deposit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var deposit depositBody
	err = json.NewDecoder(r.Body).Decode(&deposit)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = c.services.Users.Deposit(id, deposit.Balance); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (c *Controller) getBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	balance, err := c.services.Users.GetBalance(id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"balance": balance,
	})
	w.WriteHeader(http.StatusOK)

}
