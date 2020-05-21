package web

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/hoorinaz/TodoList/pkg/user"
)

func (uws *UserWebService) Register(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	ctx = context.WithValue(ctx, "name", "amir")
	u := new(user.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = uws.UserProcessor.AddUser(ctx, u); err != nil {
		log.Println("error in web level", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}
