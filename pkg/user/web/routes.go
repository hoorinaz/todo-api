package web

import (
	"encoding/json"
	"github.com/hoorinaz/TodoList/pkg/user"
	"log"
	"net/http"
)

func (uws *UserWebService) Register(w http.ResponseWriter, r *http.Request) {

	u := new(user.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = uws.UserProcessor.AddUser(u); err != nil {
		log.Println("error in web level", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
